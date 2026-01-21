package auth

import (
	"context"

	repo "github.com/haobuhaoo/gossip-with-go/internal/postgresql/sqlc"
)

// Service defines the domain logic for authentication related operations.
// It is responsible for enforcing application rules and making database calls.
type Service interface {
	Login(ctx context.Context, name string) (repo.User, error)
	AuthenticateUser(ctx context.Context, id int64) (repo.User, error)
}

// LoginRequest handles the authentication related HTTP request body for login of user.
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
}

// LoginResponse handles the authentication HTTP response body after login.
type LoginResponse struct {
	Token string `json:"token"`
}
