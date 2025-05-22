package services

import (
	"context"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5"
	repository "github.com/miloszbo/meals-finder/internal/repositories"
)

type FinderService interface {
	FindRecipe(ctx context.Context, queries *url.Values) error
}

type BaseFinderService struct {
	DbConn *pgx.Conn
	Repo   *repository.Queries
}

func (b *BaseFinderService) FindRecipe(ctx context.Context, queries *url.Values) error {
	return nil
}

type MockFinderService struct{}

func (m *MockFinderService) FindRecipe(ctx context.Context, queries *url.Values) error {
	for key, val := range *queries {
		fmt.Println(key, " ", val)
	}
	return nil
}
