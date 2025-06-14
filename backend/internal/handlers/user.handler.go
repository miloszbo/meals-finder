package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/miloszbo/meals-finder/internal/models"
	"github.com/miloszbo/meals-finder/internal/services"
)

type UserHandler struct {
	UserService services.UserService
}

func (u *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	loginData := models.LoginUserRequest{}

	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		log.Println(err.Error())
		err := ErrBadRequest
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := loginData.Validate(); err != nil {
		http.Error(w, ErrBadRequest.Error(), http.StatusBadRequest)
	}

	token, err := u.UserService.LoginUser(ctx, &loginData)
	if err != nil {
		http.Error(w, err.Error(), StatusFromError(err))
		return
	}

	cookie := &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		MaxAge:   24 * 3600,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		Secure:   false,
	}

	http.SetCookie(w, cookie)

	w.WriteHeader(http.StatusOK)
}

func (uh *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     "auth_token",
		MaxAge:   0,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   false,
	}
	http.SetCookie(w, cookie)
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Validate() != nil {
		log.Println(err.Error())
		http.Error(w, ErrBadRequest.Error(), http.StatusBadRequest)
		return
	}

	if err := uh.UserService.CreateUser(r.Context(), &req); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), StatusFromError(err))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message":"user created"}`))
}

func (uh *UserHandler) IsLogged(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (uh *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	claims, ok := ctx.Value("claims").(jwt.MapClaims)
	if !ok {
		http.Error(w, "token was empty", http.StatusUnauthorized)
	}

	user, err := uh.UserService.GetUser(ctx, claims["sub"].(string))

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "user was not found", http.StatusInternalServerError)
	}

	jsonUser, _ := json.Marshal(user)

	w.WriteHeader(http.StatusOK)
	w.Write(jsonUser)
}

func (uh *UserHandler) UpdateUserSettings(w http.ResponseWriter, r *http.Request) {
	var req models.UpdateUserSettingsRequest

	// Decode JSON input
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err.Error())
		http.Error(w, ErrBadRequest.Error(), http.StatusBadRequest)
		return
	}

	// Validate request
	if err := req.Validate(); err != nil {
		http.Error(w, ErrBadRequest.Error(), http.StatusBadRequest)
		return
	}

	// Call service
	if err := uh.UserService.UpdateUserSettings(r.Context(), &req); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), StatusFromError(err))
		return
	}

	// Success
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"user settings updated"}`))
}
