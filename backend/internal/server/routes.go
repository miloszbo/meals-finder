package server

import (
	"net/http"

	"github.com/miloszbo/meals-finder/internal/handlers"
	"github.com/miloszbo/meals-finder/internal/middlewares"
	"github.com/miloszbo/meals-finder/internal/repositories"
)

func SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	conn := NewConnection()

	 // Hello‑world endpoint
	 helloWorldHandler := handlers.NewHelloWorldHandler(conn)
	 mux.HandleFunc("GET /hi", helloWorldHandler.Greetings)
 
	 // User‑creation endpoint
	 queries := db.New(conn)
	 userCmdHandler := handlers.NewUserCommandHandler(queries)
	 mux.HandleFunc("POST /create_user", userCmdHandler.CreateUser)

	return middlewares.Logging(mux)
}
