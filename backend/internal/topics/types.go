package topics

import (
	"context"

	repo "github.com/haobuhaoo/gossip-with-go/internal/postgresql/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

// Service defines the domain logic for topic related operations.
// It is responsible for enforcing application rules and making database calls.
type Service interface {
	ListTopics(ctx context.Context) ([]repo.Topic, error)
	FindTopicByID(ctx context.Context, id int64) (repo.Topic, error)
	CreateTopic(ctx context.Context, arg repo.CreateTopicParams) (repo.Topic, error)
	UpdateTopic(ctx context.Context, arg repo.UpdateTopicParams) (repo.Topic, error)
	DeleteTopic(ctx context.Context, arg repo.DeleteTopicParams) error
	SearchTopic(ctx context.Context, query pgtype.Text) ([]repo.Topic, error)
}

// CreateTopicRequest handles the topic related HTTP request body for creation of a new topic.
type CreateTopicRequest struct {
	Title string `json:"title" validate:"required"`
}

// UpdateTopicRequest handles the topic related HTTP request body for updating of existing topic.
type UpdateTopicRequest struct {
	Title string `json:"title" validate:"required"`
}
