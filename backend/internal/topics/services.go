package topics

import (
	"context"

	repo "github.com/haobuhaoo/gossip-with-go/internal/postgresql/sqlc"
)

// Service defines the domain logic for topic related operations.
// It is responsible for enforcing application rules and making database calls.
type Service interface {
	ListTopics(ctx context.Context) ([]repo.Topic, error)
}

// svc implements the Service interface.
// It depends on the sql generated Queries type to interact with the PostgreSQL database.
type svc struct {
	repo *repo.Queries
}

// NewService creates a new topic service.
func NewService(repo *repo.Queries) Service {
	return &svc{
		repo: repo,
	}
}

// ListTopics returns all topics from the database.
func (s *svc) ListTopics(ctx context.Context) ([]repo.Topic, error) {
	return s.repo.ListTopics(ctx)
}
