package topics

import (
	"context"

	"github.com/haobuhaoo/gossip-with-go/internal/helper"
	repo "github.com/haobuhaoo/gossip-with-go/internal/postgresql/sqlc"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

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

// FindTopicByID returns a specific topic identified by id from the database.
func (s *svc) FindTopicByID(ctx context.Context, id int64) (repo.Topic, error) {
	topic, err := s.repo.FindTopicByID(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return repo.Topic{}, ErrTopicNotFound
		}
		return repo.Topic{}, err
	}

	return topic, nil
}

// CreateTopic creates and returns a new topic with the given arg params.
func (s *svc) CreateTopic(ctx context.Context, arg repo.CreateTopicParams) (repo.Topic, error) {
	topic, err := s.repo.CreateTopic(ctx, arg)
	if err != nil {
		if helper.IsUniqueViolation(err) {
			return repo.Topic{}, ErrTopicAlreadyExists
		}
		return repo.Topic{}, err
	}

	return topic, nil
}

// UpdateTopic updates an existing topic with the given arg params and returns it.
func (s *svc) UpdateTopic(ctx context.Context, arg repo.UpdateTopicParams) (repo.Topic, error) {
	topic, err := s.repo.UpdateTopic(ctx, arg)
	if err != nil {
		if err == pgx.ErrNoRows {
			return repo.Topic{}, ErrTopicNotFound
		}
		if helper.IsUniqueViolation(err) {
			return repo.Topic{}, ErrTopicAlreadyExists
		}
		return repo.Topic{}, err
	}

	return topic, nil
}

// DeleteTopic deletes the topic given by the id from the database.
// It deletes all posts under that topic too.
func (s *svc) DeleteTopic(ctx context.Context, arg repo.DeleteTopicParams) error {
	delRows, err := s.repo.DeleteTopic(ctx, arg)
	if err != nil {
		return err
	}

	if delRows == 0 {
		return ErrTopicNotFound
	}

	return nil
}

// SearchTopic searches all topic titles that contains the search query (case-insensitive)
// and returns all matched topics.
func (s *svc) SearchTopic(ctx context.Context, query pgtype.Text) ([]repo.Topic, error) {
	return s.repo.SearchTopic(ctx, query)
}
