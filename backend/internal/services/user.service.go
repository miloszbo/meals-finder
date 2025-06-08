package services

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	_ "github.com/joho/godotenv/autoload"
	"github.com/miloszbo/meals-finder/internal/models"
	repository "github.com/miloszbo/meals-finder/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

var key []byte = []byte(os.Getenv("APP_JWT_KEY"))

type UserService interface {
	LoginUser(ctx context.Context, loginData *models.LoginUserRequest) (string, error)
	CreateUser(ctx context.Context, req *models.CreateUserRequest) error
	GetUser(ctx context.Context, username string) (repository.GetUserDataRow, error)
	UpdateUserSettings(ctx context.Context, req *models.UpdateUserSettingsRequest) error
}

type BaseUserService struct {
	DbConn *pgx.Conn
	Repo   *repository.Queries
}

func NewBaseUserService(conn *pgx.Conn) BaseUserService {
	return BaseUserService{
		DbConn: conn,
		Repo:   repository.New(conn),
	}
}

func (s *BaseUserService) LoginUser(ctx context.Context, loginData *models.LoginUserRequest) (string, error) {
	user, err := s.Repo.LoginUserWithUsername(ctx, loginData.Login)
	if err != nil {
		return "", ErrUnauthorizedUser
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Passwdhash), []byte(loginData.Password)); err != nil {
		return "", ErrUnauthorizedUser
	}

	token, err := s.generateJWT(user.Username)
	if err != nil {
		log.Println(err.Error())
		return "", ErrInternalFailure
	}

	return token, nil
}

func (s *BaseUserService) CreateUser(ctx context.Context, req *models.CreateUserRequest) error {
	if err := req.Validate(); err != nil {
		return ErrInternalFailure
	}

	hashedPasswd, err := bcrypt.GenerateFromPassword([]byte(req.Passwdhash), bcrypt.DefaultCost)
	if err != nil {
		log.Println("password hashing failed:", err)
		return ErrInternalFailure
	}

	err = s.Repo.CreateUser(ctx, repository.CreateUserParams{
		Username:    req.Username,
		Passwdhash:  string(hashedPasswd),
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Age:         req.Age,
		Sex:         req.Sex,
	})

	if err != nil {
		log.Println("create user failed:", err)
		return ErrInternalFailure
	}

	return nil
}

func (s *BaseUserService) GetUser(ctx context.Context, username string) (repository.GetUserDataRow, error) {
	data, err := s.Repo.GetUserData(ctx, username)

	return data, err
}

func (s *BaseUserService) generateJWT(username string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": username,
			"exp": time.Now().Add(24 * time.Hour).Unix(),
			"iat": time.Now().Unix(),
		})
	return t.SignedString(key)
}

func (s *BaseUserService) UpdateUserSettings(ctx context.Context, req *models.UpdateUserSettingsRequest) error {
	// Validate required input
	if err := req.Validate(); err != nil {
		return ErrInternalFailure
	}

	// Update user settings
	err := s.Repo.UpdateUserSettings(ctx, repository.UpdateUserSettingsParams{
		Username:    req.Username,
		Email:       req.Email,
		Name:        req.Name,
		Surname:     req.Surname,
		PhoneNumber: req.PhoneNumber,
		Age:         req.Age,
		Sex:         req.Sex,
		Weight:      req.Weight,
		Height:      req.Height,
		Bmi:         req.Bmi,
	})
	if err != nil {
		log.Println("update user settings failed:", err)
		return ErrInternalFailure
	}

	// if user's tag update list isn't empty update tags
	if req.TagIDs != nil {
		// Get existing tags
		existingTags, err := s.Repo.GetUserTags(ctx, req.Username)
		if err != nil {
			log.Println("get user tags failed:", err)
			return ErrInternalFailure
		}

		// Convert to set
		existingTagSet := make(map[int32]struct{}, len(existingTags))
		for _, tagID := range existingTags {
			existingTagSet[tagID] = struct{}{}
		}

		incomingTagSet := make(map[int32]struct{}, len(req.TagIDs))
		for _, tagID := range req.TagIDs {
			incomingTagSet[tagID] = struct{}{}
		}

		// Insert new tags
		for tagID := range incomingTagSet {
			if _, exists := existingTagSet[tagID]; !exists {
				err = s.Repo.InsertUserTag(ctx, repository.InsertUserTagParams{
					Username: req.Username,
					TagID:    tagID,
				})
				if err != nil {
					log.Println("insert user tag failed:", err)
					return ErrInternalFailure
				}
			}
		}

		// Delete removed tags
		for tagID := range existingTagSet {
			if _, exists := incomingTagSet[tagID]; !exists {
				err = s.Repo.DeleteUserTag(ctx, repository.DeleteUserTagParams{
					Username: req.Username,
					TagID:    tagID,
				})
				if err != nil {
					log.Println("delete user tag failed:", err)
					return ErrInternalFailure
				}
			}
		}
	}

	return nil
}


// For testing
type MockUserService struct{}

func (s *MockUserService) LoginUser(ctx context.Context, loginData *models.LoginUserRequest) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": "testUser",
			"exp": time.Now().Add(24 * time.Hour).Unix(),
			"iat": time.Now().Unix(),
		})
	return t.SignedString(key)
}

func (s *MockUserService) CreateUser(ctx context.Context, req *models.CreateUserRequest) error {
	return nil
}
