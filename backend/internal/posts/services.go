package posts

import (
	"context"

	"github.com/haobuhaoo/gossip-with-go/internal/helper"
	repo "github.com/haobuhaoo/gossip-with-go/internal/postgresql/sqlc"
	"github.com/haobuhaoo/gossip-with-go/internal/topics"
	"github.com/jackc/pgx/v5"
)

// svc implements the Service interface.
// It depends on the sql generated Queries type to interact with the PostgreSQL database.
type svc struct {
	repo *repo.Queries
}

// NewService creates a new post service.
func NewService(repo *repo.Queries) Service {
	return &svc{
		repo: repo,
	}
}

// FindPostsByTopic returns all posts of the given topic id from the database.
func (s *svc) FindPostsByTopic(ctx context.Context, id int64) ([]repo.Post, error) {
	_, err := s.repo.FindTopicByID(ctx, id)
	if err != nil {
		return []repo.Post{}, topics.ErrTopicNotFound
	}

	return s.repo.FindPostsByTopic(ctx, id)
}

// FindPostByID returns a specific post identified by id from the database.
func (s *svc) FindPostByID(ctx context.Context, id int64) (repo.Post, error) {
	post, err := s.repo.FindPostByID(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return repo.Post{}, ErrPostNotFound
		}
		return repo.Post{}, err
	}

	return post, nil
}

// CreatePost creates and returns a new post with the given arg params.
func (s *svc) CreatePost(ctx context.Context, arg repo.CreatePostParams) (repo.Post, error) {
	post, err := s.repo.CreatePost(ctx, arg)
	if err != nil {
		if helper.IsUniqueViolation(err) {
			return repo.Post{}, ErrPostAlreadyExists
		}
		return repo.Post{}, err
	}

	return post, nil
}

// UpdatePost updates an existing post with the given arg params and returns it.
// It calls the corresponding functions (UpdatePost, UpdatePostTitle, UpdatePostDescription) depending on
// the given arg params.
func (s *svc) UpdatePost(ctx context.Context, arg repo.UpdatePostParams) (repo.Post, error) {
	var post repo.Post
	var err error
	if arg.Title != "" && arg.Description != "" {
		post, err = s.repo.UpdatePost(ctx, arg)
	} else if arg.Title != "" {
		updateTitle := repo.UpdatePostTitleParams{
			PostID: arg.PostID,
			Title:  arg.Title,
		}
		post, err = s.UpdatePostTitle(ctx, updateTitle)
	} else if arg.Description != "" {
		updateDesc := repo.UpdatePostDescriptionParams{
			PostID:      arg.PostID,
			Description: arg.Description,
		}
		post, err = s.UpdatePostDescription(ctx, updateDesc)
	} else {
		return repo.Post{}, InvalidRequstBody
	}

	if err != nil {
		if err == pgx.ErrNoRows {
			return repo.Post{}, ErrPostNotFound
		}
		if helper.IsUniqueViolation(err) {
			return repo.Post{}, ErrPostAlreadyExists
		}
		return repo.Post{}, err
	}

	return post, nil
}

// UpdatePostTitle updates an existing post title with the given arg params and returns it.
func (s *svc) UpdatePostTitle(ctx context.Context, arg repo.UpdatePostTitleParams) (repo.Post, error) {
	return s.repo.UpdatePostTitle(ctx, arg)
}

// UpdatePostDescription updates an existing post description with the given arg params and returns it.
func (s *svc) UpdatePostDescription(ctx context.Context, arg repo.UpdatePostDescriptionParams) (repo.Post, error) {
	return s.repo.UpdatePostDescription(ctx, arg)
}

// DeletePost deletes the post given by the id from the database.
// It deletes all comments under that post too.
func (s *svc) DeletePost(ctx context.Context, id int64) error {
	delRows, err := s.repo.DeletePost(ctx, id)
	if err != nil {
		return err
	}

	if delRows == 0 {
		return ErrPostNotFound
	}

	return nil
}
