package handlers

import (
	"context"
	"net/http"

	"github.com/miloszbo/meals-finder/internal/services"
)

type FinderHandler struct {
	FinderService services.FinderService
}

func (f *FinderHandler) FindRecipes(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	queries := r.URL.Query()

	f.FinderService.FindRecipe(ctx, &queries)
}
