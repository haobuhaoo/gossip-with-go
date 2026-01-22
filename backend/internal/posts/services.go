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
func (s *svc) FindPostsByTopic(ctx context.Context, arg repo.FindPostsByTopicParams) ([]Post, error) {
	_, err := s.repo.FindTopicByID(ctx, arg.TopicID)
	if err != nil {
		return []Post{}, topics.ErrTopicNotFound
	}

	rows, err := s.repo.FindPostsByTopic(ctx, arg)
	if err != nil {
		return []Post{}, err
	}

	posts := make([]Post, 0, len(rows))
	for _, row := range rows {
		posts = append(posts, Post{
			PostID:      row.PostID,
			TopicID:     row.TopicID,
			UserID:      row.UserID,
			Username:    row.Username,
			Title:       row.Title,
			Description: row.Description,
			Likes:       row.Likes,
			Dislikes:    row.Dislikes,
			UserVote:    row.UserVote,
			CreatedAt:   row.CreatedAt.Time,
			UpdatedAt:   row.UpdatedAt.Time,
		})
	}
	return posts, nil
}

// FindPostByID returns a specific post identified by id from the database.
func (s *svc) FindPostByID(ctx context.Context, arg repo.FindPostByIDParams) (Post, error) {
	rows, err := s.repo.FindPostByID(ctx, arg)
	if err != nil {
		if err == pgx.ErrNoRows {
			return Post{}, ErrPostNotFound
		}
		return Post{}, err
	}

	posts := Post{
		PostID:      rows.PostID,
		TopicID:     rows.TopicID,
		UserID:      rows.UserID,
		Username:    rows.Username,
		Title:       rows.Title,
		Description: rows.Description,
		Likes:       rows.Likes,
		Dislikes:    rows.Dislikes,
		UserVote:    rows.UserVote,
		CreatedAt:   rows.CreatedAt.Time,
		UpdatedAt:   rows.UpdatedAt.Time,
	}
	return posts, nil
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
func (s *svc) UpdatePost(ctx context.Context, arg repo.UpdatePostParams) (repo.Post, error) {
	post, err := s.repo.UpdatePost(ctx, arg)
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

// DeletePost deletes the post given by the id from the database.
// It deletes all comments under that post too.
func (s *svc) DeletePost(ctx context.Context, arg repo.DeletePostParams) error {
	delRows, err := s.repo.DeletePost(ctx, arg)
	if err != nil {
		return err
	}

	if delRows == 0 {
		return ErrPostNotFound
	}

	return nil
}

// SearchPost searches all post titles and descriptions under the specific topic that contains the
// search query (case-insensitive) and returns all matched posts.
func (s *svc) SearchPost(ctx context.Context, arg repo.SearchPostParams) ([]Post, error) {
	rows, err := s.repo.SearchPost(ctx, arg)
	if err != nil {
		return []Post{}, err
	}

	posts := make([]Post, 0, len(rows))
	for _, row := range rows {
		posts = append(posts, Post{
			PostID:      row.PostID,
			TopicID:     row.TopicID,
			UserID:      row.UserID,
			Username:    row.Username,
			Title:       row.Title,
			Description: row.Description,
			Likes:       row.Likes,
			Dislikes:    row.Dislikes,
			UserVote:    row.UserVote,
			CreatedAt:   row.CreatedAt.Time,
			UpdatedAt:   row.UpdatedAt.Time,
		})
	}
	return posts, nil
}

// LikesPost increments the like count for the specific post by 1.
func (s *svc) LikesPost(ctx context.Context, arg repo.LikesPostParams) error {
	err := s.repo.LikesPost(ctx, arg)
	if err != nil {
		return err
	}

	return nil
}

// DislikesPost increments the dislike count for the specific post by 1.
func (s *svc) DislikesPost(ctx context.Context, arg repo.DislikesPostParams) error {
	err := s.repo.DislikesPost(ctx, arg)
	if err != nil {
		return err
	}

	return nil
}

// RemovePostVote removes the user's vote for that specific post.
func (s *svc) RemovePostVote(ctx context.Context, arg repo.RemovePostVoteParams) error {
	delRows, err := s.repo.RemovePostVote(ctx, arg)
	if err != nil {
		return err
	}

	if delRows == 0 {
		return ErrVoteNotFound
	}

	return nil
}
