package users

import (
	"context"

	repo "github.com/haobuhaoo/gossip-with-go/internal/postgresql/sqlc"
)

// Service defines the domain logic for user related operations.
// It is responsible for enforcing application rules and making database calls.
type Service interface {
	FindUserByName(ctx context.Context, name string) (repo.User, error)
	CreateUser(ctx context.Context, name string) (repo.User, error)
}

// CreateUserRequest handles the user related HTTP request body for creation of a new user.
type CreateUserRequest struct {
	Name string `json:"name" validate:"required"`
}
