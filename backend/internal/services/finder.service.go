package services

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/miloszbo/meals-finder/internal/models"
	repository "github.com/miloszbo/meals-finder/internal/repositories"
)

type FinderService interface {
	FindRecipe(ctx context.Context, recipeParams models.RecipesFinderParams) ([]repository.FilterRecipesByTagNamesAndParamsRow, error)
	GetRecipe(ctx context.Context, id int32) (repository.Recipe, error)
	GetReview(ctx context.Context, id int32, username string) (repository.Review, error)
	GetTags(ctx context.Context) ([]repository.GetAllTagsRow, error)
	CreateRecipe(ctx context.Context, recipe *models.RecipeAdd) error
	AddReview(ctx context.Context, review *models.Review, username string) error
	GetRatings(ctx context.Context, id int32) ([]byte, error)
}

type BaseFinderService struct {
	Repo *repository.Queries
}

func NewBaseFinderService(conn *pgxpool.Pool) BaseFinderService {
	return BaseFinderService{
		Repo: repository.New(conn),
	}
}

func (b *BaseFinderService) GetRatings(ctx context.Context, id int32) ([]byte, error) {
	ratings, err := b.Repo.GetRatings(ctx, id)
	if err != nil {
		return nil, err
	}

	return ratings, nil
}

func (b *BaseFinderService) AddReview(ctx context.Context, review *models.Review, username string) error {
	_, err := b.Repo.GetReview(ctx, repository.GetReviewParams{
		Username: username,
		RecipeID: review.RecipeId,
	})

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = b.Repo.InsertReview(ctx, repository.InsertReviewParams{
				RecipeID:    review.RecipeId,
				Username:    username,
				ReviewScore: review.Stars,
			})

			if err != nil {
				return err
			}
			return nil
		}
		return err
	}

	err = b.Repo.UpdateReview(ctx, repository.UpdateReviewParams{
		RecipeID:       review.RecipeId,
		Username:       username,
		NewReviewScore: review.Stars,
	})

	if err != nil {
		return err
	}

	return nil
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
		tagId, err := b.Repo.GetTagId(ctx, tag.Name)
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

func (b *BaseFinderService) GetReview(ctx context.Context, id int32, username string) (repository.Review, error) {
	review, err := b.Repo.GetReview(ctx, repository.GetReviewParams{
		RecipeID: id,
		Username: username,
	})
	if err != nil {

	}

	return review, nil
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
		Username:      recipeParams.Username,
	})

	return recipes, nil
}

type MockFinderService struct{}

func (m *MockFinderService) FindRecipe(ctx context.Context, recipeParams models.RecipesFinderParams) ([]repository.FilterRecipesByTagNamesAndParamsRow, error) {
	log.Println(recipeParams)
	return nil, nil
}
