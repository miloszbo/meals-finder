package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/miloszbo/meals-finder/internal/models"
	"github.com/miloszbo/meals-finder/internal/services"
)

type FinderHandler struct {
	FinderService services.FinderService
}

func (f *FinderHandler) FindRecipes(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	queries := r.URL.Query()

	page64, err := strconv.ParseInt(r.URL.Query().Get("page"), 10, 32)
	page := int32(page64)
	if err != nil && page < 1 {
		page = 1
	}
	limit64, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)
	limit := int32(limit64)
	if err != nil && limit < 2 {
		limit = 2
	}

	offset := (page - 1) * limit

	minTime, _ := strconv.ParseInt(queries["minTime"][0], 10, 32)
	maxTime, _ := strconv.ParseInt(queries["maxTime"][0], 10, 32)
	minDifficulty, _ := strconv.ParseInt(queries["minDifficulty"][0], 10, 32)
	maxDifficulty, _ := strconv.ParseInt(queries["maxDifficulty"][0], 10, 32)

	recipeParams := models.RecipeFinderParams{
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
		Limit:         limit,
		Offset:        offset,
	}

	recipes, err := f.FinderService.FindRecipe(ctx, recipeParams)
	if err != nil {
		http.Error(w, err.Error(), StatusFromError(err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(recipes)
}
