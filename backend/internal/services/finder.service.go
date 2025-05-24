package services

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/jackc/pgx/v5"
	repository "github.com/miloszbo/meals-finder/internal/repositories"
)

type FinderService interface {
	FindRecipe(ctx context.Context, queries url.Values) ([]repository.FilterRecipesByTagNamesAndParamsRow, error)
}

type BaseFinderService struct {
	DbConn *pgx.Conn
	Repo   *repository.Queries
}

func (b *BaseFinderService) FindRecipe(ctx context.Context, queries url.Values) ([]repository.FilterRecipesByTagNamesAndParamsRow, error) {
	minTime, _ := strconv.Atoi(queries["minTime"][0])
	maxTime, _ := strconv.Atoi(queries["maxTime"][0])
	minDifficulty, _ := strconv.Atoi(queries["minDifficulty"][0])
	maxDifficulty, _ := strconv.Atoi(queries["maxDifficulty"][0])

	recipes, _ := b.Repo.FilterRecipesByTagNamesAndParams(ctx, repository.FilterRecipesByTagNamesAndParamsParams{
		Diet:          queries["dieta"],
		Region:        queries["region"],
		RecipeType:    queries["rodzaj"],
		Allergies:     queries["alergeny"],
		Nutrients:     queries["skladnikiOdzywcze"],
		Others:        queries["inne"],
		MinTime:       int32(minTime),
		MaxTime:       int32(maxTime),
		MinDifficulty: int32(minDifficulty),
		MaxDifficulty: int32(maxDifficulty),
	})

	return recipes, nil
}

type MockFinderService struct{}

func (m *MockFinderService) FindRecipe(ctx context.Context, queries url.Values) error {
	for key, val := range queries {
		fmt.Println(key, " ", val)
	}
	return nil
}
