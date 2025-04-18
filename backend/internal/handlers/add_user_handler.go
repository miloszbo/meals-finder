package handlers

import (
    "encoding/json"
    "net/http"

    repository "github.com/miloszbo/meals-finder/internal/repositories"
)

// CreateUserRequest is the write‑side DTO for creating a user.
type CreateUserRequest struct {
    Username string `json:"username"`
    Passwd   string `json:"passwd"`
    Email    string `json:"email"`
    Sex      bool   `json:"sex"`
    Age      int32  `json:"age"`
}

// CreateUserResponse is returned after a successful CreateUser command.
type CreateUserResponse struct {
    ID       int32  `json:"id"`       // matches SQL 'id' INTEGER → int32
    Username string `json:"username"`
    Email    string `json:"email"`
    Sex      bool   `json:"sex"`
    Age      int32  `json:"age"`
}

// UserCommandHandler handles write‑side (commands) for users.
type UserCommandHandler struct {
    repo *repository.Queries
}

// NewUserCommandHandler constructs a new UserCommandHandler.
func NewUserCommandHandler(repo *repository.Queries) *UserCommandHandler {
    return &UserCommandHandler{repo: repo}
}

// CreateUser handles POST /users.
func (h *UserCommandHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    var req CreateUserRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "invalid request payload: "+err.Error(), http.StatusBadRequest)
        return
    }

    // Pass plain Go types directly to the SQLC-generated method.
    user, err := h.repo.CreateUser(r.Context(), repository.CreateUserParams{
        Username: req.Username,
        Passwd:   req.Passwd,
        Email:    req.Email,
        Sex:      req.Sex,
        Age:      req.Age,
    })
    if err != nil {
        http.Error(w, "failed to create user: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Build response using the native Go return types.
    resp := CreateUserResponse{
        ID:       user.ID,
        Username: user.Username,
        Email:    user.Email,
        Sex:      user.Sex,
        Age:      user.Age,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(resp); err != nil {
        http.Error(w, "failed to write response: "+err.Error(), http.StatusInternalServerError)
    }
}
