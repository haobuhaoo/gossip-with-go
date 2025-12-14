package topics

import (
	"context"

	repo "github.com/haobuhaoo/gossip-with-go/internal/postgresql/sqlc"
)

type Service interface {
	ListTopics(ctx context.Context) ([]repo.Topic, error)
}

type svc struct {
	repo *repo.Queries
}

func NewService(repo *repo.Queries) Service {
	return &svc{
		repo: repo,
	}
}

func (s *svc) ListTopics(ctx context.Context) ([]repo.Topic, error) {
	return s.repo.ListTopics(ctx)
}
