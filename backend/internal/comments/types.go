package comments

import (
	"context"
	"time"

	repo "github.com/haobuhaoo/gossip-with-go/internal/postgresql/sqlc"
)

// Service defines the domain logic for comment related operations.
// It is responsible for enforcing application rules and making database calls.
type Service interface {
	FindCommentsByPost(ctx context.Context, arg repo.FindPostByIDParams) ([]Comment, error)
	CreateComment(ctx context.Context, arg repo.CreateCommentParams) (repo.Comment, error)
	UpdateComment(ctx context.Context, arg repo.UpdateCommentParams) (repo.Comment, error)
	DeleteComment(ctx context.Context, arg repo.DeleteCommentParams) error
}

// Comment model that is passed to the frontend.
type Comment struct {
	CommentID   int64     `json:"comment_id"`
	PostID      int64     `json:"post_id"`
	UserID      int64     `json:"user_id"`
	Username    string    `json:"username"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateCommentRequest handles the comment related HTTP request body for creation of a new comment.
type CreateCommentRequest struct {
	PostID      int64  `json:"postId" validate:"required,min=1"`
	Description string `json:"description" validate:"required"`
}

// UpdateCommentRequest handles the comment related HTTP request body for updating of existing comment.
type UpdateCommentRequest struct {
	PostID      int64  `json:"postId" validate:"required,min=1"`
	Description string `json:"description" validate:"required"`
}
