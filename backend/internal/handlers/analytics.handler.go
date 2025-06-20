package handlers

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/miloszbo/meals-finder/internal/models"
    "github.com/miloszbo/meals-finder/internal/services"
)

type AnalyticsHandler struct {
    AnalyticsService services.AnalyticsService
}

func (ah *AnalyticsHandler) ReceiveEvent(w http.ResponseWriter, r *http.Request) {
    var event models.AnalyticsEventRequest

    if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
        log.Println(err.Error())
        http.Error(w, ErrBadRequest.Error(), http.StatusBadRequest)
        return
    }

    if err := ah.AnalyticsService.SaveEvent(r.Context(), &event, r); err != nil {
        log.Println(err.Error())
        http.Error(w, err.Error(), StatusFromError(err))
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message":"event received"}`))
}

func (ah *AnalyticsHandler) SendToBigQuery(w http.ResponseWriter, r *http.Request) {
    if err := ah.AnalyticsService.SendToBigQuery(r.Context()); err != nil {
        log.Println(err.Error())
        http.Error(w, err.Error(), StatusFromError(err))
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message":"events sent to BigQuery"}`))
}
