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
		Name:        req.Name,
		Surname:     req.Surname,
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



func (s *BaseUserService) generateJWT(username string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": username,
			"exp": time.Now().Add(24 * time.Hour).Unix(),
			"iat": time.Now().Unix(),
		})
	return t.SignedString(key)
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
