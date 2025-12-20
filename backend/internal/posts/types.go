package posts

import (
	"context"

	repo "github.com/haobuhaoo/gossip-with-go/internal/postgresql/sqlc"
)

// Service defines the domain logic for post related operations.
// It is responsible for enforcing application rules and making database calls.
type Service interface {
	FindPostsByTopic(ctx context.Context, id int64) ([]repo.Post, error)
	FindPostByID(ctx context.Context, id int64) (repo.Post, error)
	CreatePost(ctx context.Context, arg repo.CreatePostParams) (repo.Post, error)
	UpdatePost(ctx context.Context, arg repo.UpdatePostParams) (repo.Post, error)
	UpdatePostTitle(ctx context.Context, arg repo.UpdatePostTitleParams) (repo.Post, error)
	UpdatePostDescription(ctx context.Context, arg repo.UpdatePostDescriptionParams) (repo.Post, error)
	DeletePost(ctx context.Context, id int64) error
}

// CreatePostRequest handles the post related HTTP request body for creation of a new post.
type CreatePostRequest struct {
	TopicID     int64  `json:"topicId" validate:"required,min=1"`
	UserID      int64  `json:"userId" validate:"required,min=1"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

// UpdatePostRequest handles the post related HTTP request body for updating of existing post.
type UpdatePostRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
