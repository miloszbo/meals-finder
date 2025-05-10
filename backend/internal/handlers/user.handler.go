package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/miloszbo/meals-finder/internal/models"
	"github.com/miloszbo/meals-finder/internal/services"
)

type UserHandler struct {
	UserService services.UserService
}

func (u *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	loginData := &models.LoginUserRequest{}

	if err := json.NewDecoder(r.Body).Decode(loginData); err != nil {
		log.Println(err.Error())
		err := ErrBadRequest
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := loginData.Validate(); err != nil {
		http.Error(w, ErrBadRequest.Error(), http.StatusBadRequest)
	}

	token, err := u.UserService.LoginUser(ctx, loginData)
	if err != nil {
		http.Error(w, err.Error(), StatusFromError(err))
		return
	}

	cookie := &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, cookie)

	w.WriteHeader(http.StatusOK)
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Validate() != nil {
		http.Error(w, ErrBadRequest.Error(), http.StatusBadRequest)
		return
	}

	if err := uh.UserService.CreateUser(r.Context(), &req); err != nil {
		http.Error(w, err.Error(), StatusFromError(err))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message":"user created"}`))
}