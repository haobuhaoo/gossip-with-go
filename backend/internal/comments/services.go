package comments

import (
	"context"

	repo "github.com/haobuhaoo/gossip-with-go/internal/postgresql/sqlc"
	"github.com/haobuhaoo/gossip-with-go/internal/posts"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// svc implements the Service interface.
// It depends on the sql generated Queries type to interact with the PostgreSQL database.
type svc struct {
	repo *repo.Queries
	db   *pgxpool.Pool
}

// NewService creates a new comment service.
func NewService(repo *repo.Queries, db *pgxpool.Pool) Service {
	return &svc{
		repo: repo,
		db:   db,
	}
}

// FindCommentsByPost returns all comments of the given post id from the database.
func (s *svc) FindCommentsByPost(ctx context.Context, arg repo.FindPostByIDParams) ([]Comment, error) {
	_, err := s.repo.FindPostByID(ctx, arg)
	if err != nil {
		return []Comment{}, posts.ErrPostNotFound
	}

	commentArg := repo.FindCommentsByPostParams{
		PostID: arg.PostID,
		UserID: arg.UserID,
	}
	rows, err := s.repo.FindCommentsByPost(ctx, commentArg)
	if err != nil {
		return []Comment{}, err
	}

	comments := make([]Comment, 0, len(rows))
	for _, row := range rows {
		comments = append(comments, Comment{
			CommentID:   row.CommentID,
			PostID:      row.PostID,
			UserID:      row.UserID,
			Username:    row.Username,
			Description: row.Description,
			Likes:       row.Likes,
			Dislikes:    row.Dislikes,
			UserVote:    row.UserVote,
			CreatedAt:   row.CreatedAt.Time,
			UpdatedAt:   row.UpdatedAt.Time,
		})
	}
	return comments, nil
}

// CreateComment creates and returns a new comment with the given arg params. It then updates
// the post's updated status.
// If there is an error in between, the whole transaction is rolled back.
func (s *svc) CreateComment(ctx context.Context, arg repo.CreateCommentParams) (repo.Comment, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return repo.Comment{}, err
	}
	defer tx.Rollback(ctx)

	qtx := s.repo.WithTx(tx)

	comment, err := qtx.CreateComment(ctx, arg)
	if err != nil {
		return repo.Comment{}, err
	}

	statusArg := repo.UpdatePostStatusParams{
		PostID: arg.PostID,
		UserID: arg.UserID,
	}
	_, err = qtx.UpdatePostStatus(ctx, statusArg)
	if err != nil {
		return repo.Comment{}, ErrPostNotUpdated
	}

	tx.Commit(ctx)
	return comment, nil
}

// UpdateComment updates an existing comment with the given arg params and returns it. It then updates
// the post's updated status.
// If there is an error in between, the whole transaction is rolled back.
func (s *svc) UpdateComment(ctx context.Context, arg repo.UpdateCommentParams) (repo.Comment, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return repo.Comment{}, err
	}
	defer tx.Rollback(ctx)

	qtx := s.repo.WithTx(tx)

	comment, err := qtx.UpdateComment(ctx, arg)
	if err != nil {
		if err == pgx.ErrNoRows {
			return repo.Comment{}, ErrCommentNotFound
		}
		return repo.Comment{}, err
	}

	statusArg := repo.UpdatePostStatusParams{
		PostID: arg.PostID,
		UserID: arg.UserID,
	}
	_, err = qtx.UpdatePostStatus(ctx, statusArg)
	if err != nil {
		return repo.Comment{}, ErrPostNotUpdated
	}

	tx.Commit(ctx)
	return comment, nil
}

// DeleteComment deletes the comment given by the id from the database.
func (s *svc) DeleteComment(ctx context.Context, arg repo.DeleteCommentParams) error {
	delRows, err := s.repo.DeleteComment(ctx, arg)
	if err != nil {
		return err
	}

	if delRows == 0 {
		return ErrCommentNotFound
	}

	return nil
}

// LikesComment increments the like count for the specific comment by 1.
func (s *svc) LikesComment(ctx context.Context, arg repo.LikesCommentParams) (repo.CommentVote, error) {
	comment, err := s.repo.LikesComment(ctx, arg)
	if err != nil {
		return repo.CommentVote{}, err
	}

	return comment, nil
}

// DislikesComment increments the dislike count for the specific comment by 1.
func (s *svc) DislikesComment(ctx context.Context, arg repo.DislikesCommentParams) (repo.CommentVote, error) {
	comment, err := s.repo.DislikesComment(ctx, arg)
	if err != nil {
		return repo.CommentVote{}, err
	}

	return comment, nil
}

// RemoveCommentVote removes the user's vote for that specific comment.
func (s *svc) RemoveCommentVote(ctx context.Context, arg repo.RemoveCommentVoteParams) error {
	delRows, err := s.repo.RemoveCommentVote(ctx, arg)
	if err != nil {
		return err
	}

	if delRows == 0 {
		return ErrVoteNotFound
	}

	return nil
}
