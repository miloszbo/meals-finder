package services

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/miloszbo/meals-finder/internal/models"
	repository "github.com/miloszbo/meals-finder/internal/repositories"
)

type FinderService interface {
	FindRecipe(ctx context.Context, recipeParams models.RecipesFinderParams) ([]repository.FilterRecipesByTagNamesAndParamsRow, error)
}

type BaseFinderService struct {
	DbConn *pgx.Conn
	Repo   *repository.Queries
}

func (b *BaseFinderService) FindRecipe(ctx context.Context, recipeParams models.RecipesFinderParams) ([]repository.FilterRecipesByTagNamesAndParamsRow, error) {
	recipes, _ := b.Repo.FilterRecipesByTagNamesAndParams(ctx, repository.FilterRecipesByTagNamesAndParamsParams{
		Diet:          recipeParams.Diet,
		Region:        recipeParams.Region,
		RecipeType:    recipeParams.RecipeType,
		Allergies:     recipeParams.Allergies,
		Nutrients:     recipeParams.Nutrients,
		Others:        recipeParams.Others,
		MinTime:       recipeParams.MinTime,
		MaxTime:       recipeParams.MaxTime,
		MinDifficulty: recipeParams.MinDifficulty,
		MaxDifficulty: recipeParams.MaxDifficulty,
		RecipesOffset: recipeParams.Offset,
		RecipesLimit:  recipeParams.Limit,
	})

	return recipes, nil
}

type MockFinderService struct{}

func (m *MockFinderService) FindRecipe(ctx context.Context, recipeParams models.RecipesFinderParams) ([]repository.FilterRecipesByTagNamesAndParamsRow, error) {
	log.Println(recipeParams)
	return nil, nil
}
