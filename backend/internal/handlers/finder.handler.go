package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/miloszbo/meals-finder/internal/services"
)

type FinderHandler struct {
	FinderService services.FinderService
}

func (f *FinderHandler) FindRecipes(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	queries := r.URL.Query()

	recipes, err := f.FinderService.FindRecipe(ctx, queries)
	if err != nil {
		http.Error(w, err.Error(), StatusFromError(err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(recipes)
}
