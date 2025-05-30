package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/miloszbo/meals-finder/internal/handlers"
	"github.com/miloszbo/meals-finder/internal/services"
)

func TestFindRecipes(t *testing.T) {
	handler := handlers.FinderHandler{
		FinderService: &services.MockFinderService{},
	}

	req := httptest.NewRequest(http.MethodGet, "/finder?region=polska&region=wlochy&dieta=keto", nil)
	res := httptest.NewRecorder()

	handler.FindRecipes(res, req)
}
