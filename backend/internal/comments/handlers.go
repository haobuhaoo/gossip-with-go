package comments

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/haobuhaoo/gossip-with-go/internal/helper"
	repo "github.com/haobuhaoo/gossip-with-go/internal/postgresql/sqlc"
	"github.com/haobuhaoo/gossip-with-go/internal/posts"
)

const (
	InvalidCommentIdMessage            = "Invalid comment id"
	InvalidRequestBodyMessage          = "Required fields missing"
	MissingUserIDMessage               = "Missing userID"
	SuccessfulFindCommentByPostMessage = "Successfully listed all comments"
	SuccessfulCreateCommentMessage     = "Successfully created comment"
	SuccessfulUpdateCommentMessage     = "Successfully updated comment"
	SuccessfulDeleteCommentMessage     = "Successfully deleted comment"
	SuccessfulLikeCommentMessage       = "Successfully liked comment"
	SuccessfulDislikeCommentMessage    = "Successfully disliked comment"
	SuccessfulRemoveCommentVoteMessage = "Successfully removed vote"
)

// handler handles the comment related HTTP requests.
// It is responsible for translating HTTP requests into service calls and formatting service
// responses into HTTP responses.
type handler struct {
	service Service
}

// NewHandler creates a new comment handler.
func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

// FindCommentsByPost handles GET /api/comments/all/{topicId}/{postId} requests.
// It parses the topicId and postId string, and passes it to the comment service to return all
// comments for that post, and serializes the result into a JSON HTTP response.
func (h *handler) FindCommentsByPost(w http.ResponseWriter, r *http.Request) {
	topicIdStr := chi.URLParam(r, "topicId")
	topicId, err := strconv.ParseInt(topicIdStr, 10, 64)
	if err != nil {
		helper.WriteError(w, posts.InvalidTopicIdMessage, http.StatusBadRequest)
		return
	}

	postIdStr := chi.URLParam(r, "postId")
	postId, err := strconv.ParseInt(postIdStr, 10, 64)
	if err != nil {
		helper.WriteError(w, posts.InvalidPostIdMessage, http.StatusBadRequest)
		return
	}

	userId, ok := r.Context().Value("userID").(int64)
	if !ok {
		helper.WriteError(w, MissingUserIDMessage, http.StatusBadRequest)
		return
	}

	req := repo.FindPostByIDParams{
		PostID:  postId,
		TopicID: topicId,
		UserID:  userId,
	}
	comments, err := h.service.FindCommentsByPost(r.Context(), req)
	if err != nil {
		if err == posts.ErrPostNotFound {
			helper.WriteError(w, posts.InvalidPostIdMessage, http.StatusBadRequest)
			return
		}

		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonComment, err := json.Marshal(comments)
	if err != nil {
		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := helper.ParseResponseDataAndMessage(jsonComment, SuccessfulFindCommentByPostMessage)
	helper.Write(w, response)
}

// CreateComment handles POST /api/comments requests.
// It reads and validates the request body, and passes it to the comment service to create the new
// comment with a description. It then serializes the result into a JSON HTTP response.
func (h *handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	var req CreateCommentRequest
	err := helper.Read(r, &req)
	if err != nil {
		helper.WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validator.New().Struct(req)
	if err != nil {
		helper.WriteError(w, InvalidRequestBodyMessage, http.StatusBadRequest)
		return
	}

	userId, ok := r.Context().Value("userID").(int64)
	if !ok {
		helper.WriteError(w, MissingUserIDMessage, http.StatusBadRequest)
		return
	}

	newComment := repo.CreateCommentParams{
		UserID:      userId,
		PostID:      req.PostID,
		Description: req.Description,
	}
	comment, err := h.service.CreateComment(r.Context(), newComment)
	if err != nil {
		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonComment, err := json.Marshal(comment)
	if err != nil {
		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := helper.ParseResponseDataAndMessage(jsonComment, SuccessfulCreateCommentMessage)
	helper.Write(w, response)
}

// UpdateComment handles PUT /api/comments/{id} requests.
// It parses the id string, reads and validates the request body, and passes it to the comment
// service to update the existing comment with the new description. It then serializes the result
// into a JSON HTTP response.
func (h *handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		helper.WriteError(w, InvalidCommentIdMessage, http.StatusBadRequest)
		return
	}

	var req UpdateCommentRequest
	err = helper.Read(r, &req)
	if err != nil {
		helper.WriteError(w, InvalidRequestBodyMessage, http.StatusBadRequest)
		return
	}

	err = validator.New().Struct(req)
	if err != nil {
		helper.WriteError(w, InvalidRequestBodyMessage, http.StatusBadRequest)
		return
	}

	userId, ok := r.Context().Value("userID").(int64)
	if !ok {
		helper.WriteError(w, MissingUserIDMessage, http.StatusBadRequest)
		return
	}

	newComment := repo.UpdateCommentParams{
		CommentID:   id,
		PostID:      req.PostID,
		UserID:      userId,
		Description: req.Description,
	}
	comment, err := h.service.UpdateComment(r.Context(), newComment)
	if err != nil {
		if err == ErrCommentNotFound {
			helper.WriteError(w, ErrCommentNotFound.Error(), http.StatusNotFound)
			return
		}

		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonComment, err := json.Marshal(comment)
	if err != nil {
		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := helper.ParseResponseDataAndMessage(jsonComment, SuccessfulUpdateCommentMessage)
	helper.Write(w, response)
}

// DeleteComment handles DELETE /api/comments/{id} requests.
// It parses the id string, and passes it to the comment service to delete the specified comment,
// which then serializes the result into a JSON HTTP response.
func (h *handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		helper.WriteError(w, InvalidCommentIdMessage, http.StatusBadRequest)
		return
	}

	userId, ok := r.Context().Value("userID").(int64)
	if !ok {
		helper.WriteError(w, MissingUserIDMessage, http.StatusBadRequest)
		return
	}

	arg := repo.DeleteCommentParams{
		CommentID: id,
		UserID:    userId,
	}
	err = h.service.DeleteComment(r.Context(), arg)
	if err != nil {
		if err == ErrCommentNotFound {
			helper.WriteError(w, ErrCommentNotFound.Error(), http.StatusNotFound)
			return
		}

		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := helper.ParseResponseMessage(SuccessfulDeleteCommentMessage)
	helper.Write(w, response)
}

// LikesComment handles POST /api/comments/{id}/likes requests.
// It parses the id string and passes it to the comment service to increment a like count for that
// specified comment, which then serializes the result into a JSON HTTP response.
func (h *handler) LikesComment(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		helper.WriteError(w, InvalidCommentIdMessage, http.StatusBadRequest)
		return
	}

	userId, ok := r.Context().Value("userID").(int64)
	if !ok {
		helper.WriteError(w, MissingUserIDMessage, http.StatusBadRequest)
		return
	}

	arg := repo.LikesCommentParams{
		CommentID: id,
		UserID:    userId,
	}
	err = h.service.LikesComment(r.Context(), arg)
	if err != nil {
		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := helper.ParseResponseMessage(SuccessfulLikeCommentMessage)
	helper.Write(w, response)
}

// DislikesComment handles POST /api/comments/{id}/dislikes requests.
// It parses the id string and passes it to the comment service to increment a dislike count for that
// specified comment, which then serializes the result into a JSON HTTP response.
func (h *handler) DislikesComment(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		helper.WriteError(w, InvalidCommentIdMessage, http.StatusBadRequest)
		return
	}

	userId, ok := r.Context().Value("userID").(int64)
	if !ok {
		helper.WriteError(w, MissingUserIDMessage, http.StatusBadRequest)
		return
	}

	arg := repo.DislikesCommentParams{
		CommentID: id,
		UserID:    userId,
	}
	err = h.service.DislikesComment(r.Context(), arg)
	if err != nil {
		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := helper.ParseResponseMessage(SuccessfulDislikeCommentMessage)
	helper.Write(w, response)
}

// RemoveCommentVote handles DELETE /api/comments/{id}/remove requests.
// It parses the id string and passes it to the comment service to remove the vote count for that
// specified comment, which then serializes the result into a JSON HTTP response.
func (h *handler) RemoveCommentVote(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		helper.WriteError(w, InvalidCommentIdMessage, http.StatusBadRequest)
		return
	}

	userId, ok := r.Context().Value("userID").(int64)
	if !ok {
		helper.WriteError(w, MissingUserIDMessage, http.StatusBadRequest)
		return
	}

	arg := repo.RemoveCommentVoteParams{
		CommentID: id,
		UserID:    userId,
	}
	err = h.service.RemoveCommentVote(r.Context(), arg)
	if err != nil {
		if err == ErrVoteNotFound {
			helper.WriteError(w, ErrVoteNotFound.Error(), http.StatusNotFound)
			return
		}

		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := helper.ParseResponseMessage(SuccessfulRemoveCommentVoteMessage)
	helper.Write(w, response)
}
