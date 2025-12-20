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
	SuccessfulFindCommentByPostMessage = "Successfully listed all comments"
	SuccessfulCreateCommentMessage     = "Successfully created comment"
	SuccessfulUpdateCommentMessage     = "Successfully updated comment"
	SuccessfulDeleteCommentMessage     = "Successfully deleted comment"
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

// FindCommentsByPost handles GET /comments/all/{postId} requests.
// It parses the postId string, and passes it to the comment service to return all comments for that post,
// and serializes the result into a JSON HTTP response.
func (h *handler) FindCommentsByPost(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "postId")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, posts.InvalidPostIdMessage, http.StatusBadRequest)
		return
	}

	comments, err := h.service.FindCommentsByPost(r.Context(), id)
	if err != nil {
		if err == posts.ErrPostNotFound {
			http.Error(w, posts.InvalidPostIdMessage, http.StatusBadRequest)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonComment, err := json.Marshal(comments)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := helper.ParseResponseDataAndMessage(jsonComment, SuccessfulFindCommentByPostMessage)
	helper.Write(w, response)
}

// CreateComment handles POST /comments requests.
// It reads and validates the request body, and passes it to the comment service to create the new
// comment with a description. It then serializes the result into a JSON HTTP response.
func (h *handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	var req CreateCommentRequest
	err := helper.Read(r, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validator.New().Struct(req)
	if err != nil {
		http.Error(w, InvalidRequestBodyMessage, http.StatusBadRequest)
		return
	}

	newComment := repo.CreateCommentParams{
		UserID:      req.UserID,
		PostID:      req.PostID,
		Description: req.Description,
	}
	comment, err := h.service.CreateComment(r.Context(), newComment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonComment, err := json.Marshal(comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := helper.ParseResponseDataAndMessage(jsonComment, SuccessfulCreateCommentMessage)
	helper.Write(w, response)
}

// UpdateComment handles PUT /comments/{id} requests.
// It parses the id string, reads and validates the request body, and passes it to the comment
// service to update the existing comment with the new description. It then serializes the result
// into a JSON HTTP response.
func (h *handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, InvalidCommentIdMessage, http.StatusBadRequest)
		return
	}

	var req UpdateCommentRequest
	err = helper.Read(r, &req)
	if err != nil {
		http.Error(w, InvalidRequestBodyMessage, http.StatusBadRequest)
		return
	}

	err = validator.New().Struct(req)
	if err != nil {
		http.Error(w, InvalidRequestBodyMessage, http.StatusBadRequest)
		return
	}

	newComment := repo.UpdateCommentParams{
		CommentID:   id,
		PostID:      req.PostID,
		Description: req.Description,
	}
	comment, err := h.service.UpdateComment(r.Context(), newComment)
	if err != nil {
		if err == ErrCommentNotFound {
			http.Error(w, ErrCommentNotFound.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonComment, err := json.Marshal(comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := helper.ParseResponseDataAndMessage(jsonComment, SuccessfulUpdateCommentMessage)
	helper.Write(w, response)
}

// DeleteComment handles DELETE /comments/{id} requests.
// It parses the id string, and passes it to the comment service to delete the specified comment,
// which then serializes the result into a JSON HTTP response.
func (h *handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, InvalidCommentIdMessage, http.StatusBadRequest)
		return
	}

	err = h.service.DeleteComment(r.Context(), id)
	if err != nil {
		if err == ErrCommentNotFound {
			http.Error(w, ErrCommentNotFound.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := helper.ParseResponseMessage(SuccessfulDeleteCommentMessage)
	helper.Write(w, response)
}
