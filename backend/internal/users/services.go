package users

import (
	"context"

	"github.com/haobuhaoo/gossip-with-go/internal/helper"
	repo "github.com/haobuhaoo/gossip-with-go/internal/postgresql/sqlc"
	"github.com/jackc/pgx/v5"
)

// svc implements the Service interface.
// It depends on the sql generated Queries type to interact with the PostgreSQL database.
type svc struct {
	repo *repo.Queries
}

// NewService creates a new user service.
func NewService(repo *repo.Queries) Service {
	return &svc{
		repo: repo,
	}
}

// FindUserByName returns a specific user identified by the name from the database.
func (s *svc) FindUserByName(ctx context.Context, name string) (repo.User, error) {
	user, err := s.repo.FindUserByName(ctx, name)
	if err != nil {
		if err == pgx.ErrNoRows {
			return repo.User{}, ErrUserNotFound
		}
		return repo.User{}, err
	}

	return user, nil
}

// CreateUser creates and returns a new user with the given name.
func (s *svc) CreateUser(ctx context.Context, name string) (repo.User, error) {
	user, err := s.repo.CreateUser(ctx, name)
	if err != nil {
		if helper.IsUniqueViolation(err) {
			return repo.User{}, ErrUserAlreadyExists
		}
		return repo.User{}, err
	}

	return user, nil
}
