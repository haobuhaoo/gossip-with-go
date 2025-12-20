package comments

import (
	"context"

	repo "github.com/haobuhaoo/gossip-with-go/internal/postgresql/sqlc"
)

// Service defines the domain logic for comment related operations.
// It is responsible for enforcing application rules and making database calls.
type Service interface {
	FindCommentsByPost(ctx context.Context, id int64) ([]repo.Comment, error)
	CreateComment(ctx context.Context, arg repo.CreateCommentParams) (repo.Comment, error)
	UpdateComment(ctx context.Context, arg repo.UpdateCommentParams) (repo.Comment, error)
	DeleteComment(ctx context.Context, id int64) error
}

// CreateCommentRequest handles the comment related HTTP request body for creation of a new comment.
type CreateCommentRequest struct {
	UserID      int64  `json:"userId" validate:"required,min=1"`
	PostID      int64  `json:"postId" validate:"required,min=1"`
	Description string `json:"description" validate:"required"`
}

// UpdateCommentRequest handles the comment related HTTP request body for updating of existing comment.
type UpdateCommentRequest struct {
	PostID      int64  `json:"postId" validate:"required,min=1"`
	Description string `json:"description" validate:"required"`
}
