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
	GetRecipe(ctx context.Context, id int32) (repository.Recipe, error)
	GetTags(ctx context.Context) ([]repository.GetAllTagsRow, error)
	CreateRecipe(ctx context.Context, recipe *models.RecipeAdd) error
}

type BaseFinderService struct {
	DbConn *pgx.Conn
	Repo   *repository.Queries
}

func NewBaseFinderService(conn *pgx.Conn) BaseFinderService {
	return BaseFinderService{
		DbConn: conn,
		Repo:   repository.New(conn),
	}
}

func (b *BaseFinderService) CreateRecipe(ctx context.Context, recipe *models.RecipeAdd) error {
	id, err := b.Repo.CreateRecipe(ctx, repository.CreateRecipeParams{
		Name:        recipe.Name,
		Recipe:      recipe.Recipe,
		Ingredients: recipe.Ingredients,
		Time:        recipe.Time,
		Difficulty:  recipe.Difficulty,
	})

	if err != nil {
		log.Println(err.Error())
		return err
	}

	for _, tag := range recipe.Tags {
		tagId, err := b.Repo.GetTagId(ctx, repository.GetTagIdParams{
			Key:   tag.TagType,
			Value: tag.Name,
		})
		if err != nil {
			log.Println(err.Error())
			return err
		}
		err = b.Repo.AddTagsForRecipe(ctx, repository.AddTagsForRecipeParams{
			TagID:    tagId,
			RecipeID: id,
		})

		if err != nil {
			log.Println(err.Error())
			return err
		}
	}

	return nil
}

func (b *BaseFinderService) GetTags(ctx context.Context) ([]repository.GetAllTagsRow, error) {
	tags, err := b.Repo.GetAllTags(ctx)
	return tags, err
}

func (b *BaseFinderService) GetRecipe(ctx context.Context, id int32) (repository.Recipe, error) {
	recipe, err := b.Repo.GetRecipeWithId(ctx, id)
	if err != nil {

	}

	return recipe, nil
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
