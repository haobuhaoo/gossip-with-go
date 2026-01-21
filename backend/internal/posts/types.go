package posts

import (
	"context"
	"time"

	repo "github.com/haobuhaoo/gossip-with-go/internal/postgresql/sqlc"
)

// Service defines the domain logic for post related operations.
// It is responsible for enforcing application rules and making database calls.
type Service interface {
	FindPostsByTopic(ctx context.Context, arg repo.FindPostsByTopicParams) ([]Post, error)
	FindPostByID(ctx context.Context, arg repo.FindPostByIDParams) (Post, error)
	CreatePost(ctx context.Context, arg repo.CreatePostParams) (repo.Post, error)
	UpdatePost(ctx context.Context, arg repo.UpdatePostParams) (repo.Post, error)
	DeletePost(ctx context.Context, arg repo.DeletePostParams) error
	SearchPost(ctx context.Context, arg repo.SearchPostParams) ([]Post, error)
}

// Post model that is passed to the frontend.
type Post struct {
	PostID      int64     `json:"post_id"`
	TopicID     int64     `json:"topic_id"`
	UserID      int64     `json:"user_id"`
	Username    string    `json:"username"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreatePostRequest handles the post related HTTP request body for creation of a new post.
type CreatePostRequest struct {
	TopicID     int64  `json:"topicId" validate:"required,min=1"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

// UpdatePostRequest handles the post related HTTP request body for updating of existing post.
type UpdatePostRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}
