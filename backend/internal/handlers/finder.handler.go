package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/miloszbo/meals-finder/internal/models"
	"github.com/miloszbo/meals-finder/internal/services"
)

type FinderHandler struct {
	FinderService services.FinderService
}

func (f *FinderHandler) GetRecipe(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	id64, err := strconv.ParseInt(r.PathValue("id"), 10, 32)
	id := int32(id64)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
	}
	recipe, err := f.FinderService.GetRecipe(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), StatusFromError(err))
		return
	}

	recipeJson, err := json.Marshal(recipe)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(recipeJson)
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
		limit = 100
	}

	offset := (page - 1) * limit

	//minTime, _ := strconv.ParseInt(queries["minTime"][0], 10, 32)
	//maxTime, _ := strconv.ParseInt(queries["maxTime"][0], 10, 32)
	//minDifficulty, _ := strconv.ParseInt(queries["minDifficulty"][0], 10, 32)
	//maxDifficulty, _ := strconv.ParseInt(queries["maxDifficulty"][0], 10, 32)

	minTime := 1
	maxTime := 1000
	minDifficulty := 1
	maxDifficulty := 5

	recipeParams := models.RecipesFinderParams{
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

	recipesJson, err := json.Marshal(recipes)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(recipesJson)
}
