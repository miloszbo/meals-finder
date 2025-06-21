package services

import (
	"bytes"
	cloudctx "context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"time"
    "cloud.google.com/go/storage"
    "google.golang.org/api/iterator"
    "google.golang.org/api/option" //some inferfaces google uses for paging etc.
	"github.com/miloszbo/meals-finder/internal/models"
)

// ---------------- RUNTESTING BEGIN ----------------
// jeśli w środowisku istnieje plik klucza JSON => twórz klienta z credentials file
func newStorageClientWithOptionalKey(ctx cloudctx.Context) (*storage.Client, error) {
    keyPath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
    if keyPath != "" {
        return storage.NewClient(ctx, option.WithCredentialsFile(keyPath))
    }
    return storage.NewClient(ctx) // Cloud Run / Workload Identity
}
// ---------------- RUNTESTING END ----------------

const (
	// one-shot buffer size before flushing (≈64 KiB)
	bufferFlushThreshold = 64 * 1024
	// how long we keep writer open before flushing (safety)
	bufferFlushInterval = 30 * time.Second
)

// AnalyticsService – contract exposed to handler
type AnalyticsService interface {
	SaveEvent(ctx cloudctx.Context, req *models.AnalyticsEventRequest, r *http.Request) error
	SendToBigQuery(ctx cloudctx.Context) error
}

// BaseAnalyticsService – concrete implementation
type BaseAnalyticsService struct {
	storageClient *storage.Client
	bucketName    string
	// simple in-memory buffer for demo; in prod → Cloud Pub/Sub or tmp files
	buffers map[string]*bucketBuffer
}

type bucketBuffer struct {
	buf      *bytes.Buffer
	lastFlush time.Time
}

// NewBaseAnalyticsService initialises the service
func NewBaseAnalyticsService(connPlaceholder any) BaseAnalyticsService { // connPlaceholder kept to stay symmetric with other services
	ctx := cloudctx.Background()

	//client, err := storage.NewClient(ctx)
	// ---------------- RUNTESTING BEGIN ----------------
    client, err := newStorageClientWithOptionalKey(ctx)
    // ---------------- RUNTESTING END ------------------
	if err != nil {
		// On Cloud Run this should never fail; if it does – panic is acceptable during cold-start
		log.Fatalf("failed to create storage client: %v", err)
	}

	bucket := os.Getenv("ANALYTICS_BUCKET")
	if bucket == "" {
		bucket = "PLACEHOLDER_ANALYTICS_BUCKET"
	}

	return BaseAnalyticsService{
		storageClient: client,
		bucketName:    bucket,
		buffers:       make(map[string]*bucketBuffer),
	}
}

/* ------------------------------------------------------------------ */
/* SaveEvent ­– called per HTTP POST /analytics/events                */
/* ------------------------------------------------------------------ */
func (s *BaseAnalyticsService) SaveEvent(ctx cloudctx.Context, req *models.AnalyticsEventRequest, r *http.Request) error {
	// enrich IP if missing (only for login event)
	if req.Event == "user_logged_in" && req.IPAddress == "" {
		req.IPAddress = clientIP(r)
	}

	// enrich timestamp (to RFC3339 string) & convert to record
	record, objectPath, err := s.mapToRecord(req)
	if err != nil {
		log.Println("mapToRecord error:", err)
		return ErrInternalFailure
	}

	payload, err := json.Marshal(record)
	if err != nil {
		log.Println("json marshal error:", err)
		return ErrInternalFailure
	}

	// append newline (JSON Lines)
	payload = append(payload, '\n')

	// write to buffer (per objectPath)
	buf := s.getBuffer(objectPath)
	if _, err := buf.buf.Write(payload); err != nil {
		log.Println("buffer write error:", err)
		return ErrInternalFailure
	}

	if buf.buf.Len() >= bufferFlushThreshold || time.Since(buf.lastFlush) >= bufferFlushInterval {
		if err := s.flushBuffer(ctx, objectPath, buf); err != nil {
			return ErrInternalFailure
		}
	}

	return nil
}

/* ------------------------------------------------------------------ */
/* SendToBigQuery ­– called by Cloud Scheduler /analytics/send        */
/* ------------------------------------------------------------------ */
func (s *BaseAnalyticsService) SendToBigQuery(ctx cloudctx.Context) error {
	// flush any in-mem buffers first
	for objectPath, buf := range s.buffers {
		if buf.buf.Len() > 0 {
			if err := s.flushBuffer(ctx, objectPath, buf); err != nil {
				return ErrInternalFailure
			}
		}
	}

	bkt := s.storageClient.Bucket(s.bucketName)
	it := bkt.Objects(ctx, &storage.Query{
		Prefix: "",
	})

	for {
		objAttrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println("GCS list error:", err)
			return ErrInternalFailure
		}

		// trigger BigQuery load job (placeholder)
		if err := s.loadIntoBigQuery(ctx, objAttrs.Name); err != nil {
			log.Println("BQ load error:", err)
			return ErrInternalFailure
		}

		// delete file after successful load
		if err := bkt.Object(objAttrs.Name).Delete(ctx); err != nil {
			log.Println("GCS delete error:", err)
			return ErrInternalFailure
		}
	}

	return nil
}

/* ------------------------------------------------------------------ */
/* Helper functions                                                   */
/* ------------------------------------------------------------------ */

// mapToRecord converts inbound event → proper record struct + GCS object path
func (s *BaseAnalyticsService) mapToRecord(req *models.AnalyticsEventRequest) (interface{}, string, error) {
	now := time.Now().UTC().Format(time.RFC3339)

	switch req.Event {
	case "page_visited":
		rec := models.PageVisitedRecord{
			SessionID: req.SessionID,
			Username:  req.Username,
			Path:      req.Path,
			VisitTime: now,
			Referrer:  req.Referrer,
			Device:    req.Device,
		}
		return rec, s.buildObjectPath("page_visited"), nil

	case "user_registered":
		rec := models.UserRegisteredRecord{
			Username:         req.Username,
			RegistrationTime: now,
			Device:           req.Device,
		}
		return rec, s.buildObjectPath("user_registered"), nil

	case "user_logged_in":
		rec := models.UserLoggedInRecord{
			Username:  req.Username,
			LoginTime: now,
			Device:    req.Device,
			IPAddress: req.IPAddress,
		}
		return rec, s.buildObjectPath("user_logged_in"), nil

	case "recipe_opened":
		rec := models.RecipeOpenedRecord{
			Username: req.Username,
			RecipeID: req.RecipeID,
			OpenTime: now,
			Device:   req.Device,
		}
		return rec, s.buildObjectPath("recipe_opened"), nil
	default:
		return nil, "", fmt.Errorf("unknown event type")
	}
}

func (s *BaseAnalyticsService) buildObjectPath(event string) string {
	datePart := time.Now().UTC().Format("2006-01-02")
	filename := fmt.Sprintf("%s_%s.jsonl", event, datePart)
	return path.Join(event, filename) // folder per event
}

func (s *BaseAnalyticsService) getBuffer(objectPath string) *bucketBuffer {
	if buf, ok := s.buffers[objectPath]; ok {
		return buf
	}

	s.buffers[objectPath] = &bucketBuffer{
		buf:      &bytes.Buffer{},
		lastFlush: time.Now(),
	}
	return s.buffers[objectPath]
}

func (s *BaseAnalyticsService) flushBuffer(ctx cloudctx.Context, objectPath string, buf *bucketBuffer) error {
	bkt := s.storageClient.Bucket(s.bucketName)
	obj := bkt.Object(objectPath).If(storage.Conditions{DoesNotExist: false})
	w := obj.NewWriter(ctx)
	// open existing object for append (compose workaround):
	if err := copyExisting(ctx, s.storageClient, s.bucketName, objectPath, w); err != nil {
		return err
	}

	if _, err := io.Copy(w, buf.buf); err != nil {
		_ = w.Close()
		return err
	}
	if err := w.Close(); err != nil {
		return err
	}

	buf.buf.Reset()
	buf.lastFlush = time.Now()
	return nil
}

// copyExisting appends existing object (GCS has no append API, we re-write)
func copyExisting(ctx cloudctx.Context, client *storage.Client, bucket, object string, dst *storage.Writer) error {
	src := client.Bucket(bucket).Object(object)
	r, err := src.NewReader(ctx)
	if err != nil {
		// if not exist, that's fine
		if err == storage.ErrObjectNotExist {
			return nil
		}
		return err
	}
	defer r.Close()

	_, err = io.Copy(dst, r)
	return err
}

// loadIntoBigQuery triggers placeholder load job
func (s *BaseAnalyticsService) loadIntoBigQuery(ctx cloudctx.Context, objectName string) error {
	// TODO: implement BigQuery load via bigquery.Client
	log.Printf("PLACEHOLDER: loading gs://%s/%s into BigQuery\n", s.bucketName, objectName)
	return nil
}

// clientIP extracts IP from http.Request
func clientIP(r *http.Request) string {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}
