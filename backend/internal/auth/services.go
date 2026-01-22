package auth

import (
	"context"

	repo "github.com/haobuhaoo/gossip-with-go/internal/postgresql/sqlc"
	"github.com/jackc/pgx/v5"
)

// svc implements the Service interface.
// It depends on the sql generated Queries type to interact with the PostgreSQL database.
type svc struct {
	repo *repo.Queries
}

// NewService creates a new authentication service.
func NewService(repo *repo.Queries) Service {
	return &svc{
		repo: repo,
	}
}

// Login finds and returns the user identified by the name in the database.
func (s *svc) Login(ctx context.Context, name string) (repo.User, error) {
	user, err := s.repo.FindUserByName(ctx, name)
	if err != nil {
		if err == pgx.ErrNoRows {
			return repo.User{}, ErrUserNotFound
		}
		return repo.User{}, err
	}

	return user, nil
}

// AuthenticateUser finds and returns the user identified by the ID in the database.
func (s *svc) AuthenticateUser(ctx context.Context, id int64) (repo.User, error) {
	user, err := s.repo.FindUserByID(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return repo.User{}, ErrUserNotFound
		}
		return repo.User{}, err
	}

	return user, nil
}
