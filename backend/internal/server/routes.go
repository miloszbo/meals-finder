package server

import (
	"net/http"

	"github.com/miloszbo/meals-finder/internal/handlers"
	"github.com/miloszbo/meals-finder/internal/middlewares"
	"github.com/miloszbo/meals-finder/internal/services"
)

func SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	conn := NewConnection()

	userService := services.NewBaseUserService(conn)
	userHandler := handlers.UserHandler{
		UserService: &userService,
	}

	mux.HandleFunc("POST /user/login", userHandler.LoginUser)
	mux.HandleFunc("POST /user/register", userHandler.CreateUser)
	mux.HandleFunc("GET /logout", userHandler.Logout)
	mux.HandleFunc("POST /user/settings", userHandler.UpdateUserSettings)

	stack := middlewares.CreateStack(
		middlewares.Logging,
		middlewares.CorsMiddleware,
	)

	finderService := services.NewBaseFinderService(conn)
	finderHandler := handlers.FinderHandler{
		FinderService: &finderService,
	}

	authMux := http.NewServeMux()
	authMux.HandleFunc("GET /profile", userHandler.GetProfile)
	authMux.HandleFunc("GET /verify", userHandler.IsLogged)
	authMux.HandleFunc("GET /browser", finderHandler.FindRecipes)
	authMux.HandleFunc("GET /re/{id}", finderHandler.GetRecipe)

	mux.Handle("/", middlewares.Authentication(authMux))

	return stack(mux)
}
