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

	stack := middlewares.CreateStack(
		middlewares.CorsMiddleware,
		middlewares.Logging,
	)

	authMux := http.NewServeMux()
	authMux.HandleFunc("GET /profile", userHandler.GetProfile)

	mux.Handle("/", middlewares.Authentication(authMux))

	return stack(mux)
}
