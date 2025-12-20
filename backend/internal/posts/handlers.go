package posts

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/haobuhaoo/gossip-with-go/internal/helper"
	repo "github.com/haobuhaoo/gossip-with-go/internal/postgresql/sqlc"
	"github.com/haobuhaoo/gossip-with-go/internal/topics"
)

const (
	InvalidPostIdMessage             = "Invalid post id"
	InvalidRequestBodyMessage        = "Required fields missing"
	SuccessfulFindPostByTopicMessage = "Successfully listed all posts"
	SuccessfulFindPostByIdMessage    = "Successfully find post"
	SuccessfulCreatePostMessage      = "Successfully created post"
	SuccessfulUpdatePostMessage      = "Successfully updated post"
	SuccessfulDeletePostMessage      = "Successfully deleted post"
)

// handler handles the post related HTTP requests.
// It is responsible for translating HTTP requests into service calls and formatting service
// responses into HTTP responses.
type handler struct {
	service Service
}

// NewHandler creates a new post handler.
func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

// FindPostsByTopic handles GET /posts/all/{topicId} requests.
// It parses the topicId string, and passes it to the post service to return all posts for that topic
// and serializes the result into a JSON HTTP response.
func (h *handler) FindPostsByTopic(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "topicId")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, topics.InvalidTopicIdMessage, http.StatusBadRequest)
		return
	}

	posts, err := h.service.FindPostsByTopic(r.Context(), id)
	if err != nil {
		if err == topics.ErrTopicNotFound {
			http.Error(w, topics.InvalidTopicIdMessage, http.StatusBadRequest)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonPost, err := json.Marshal(posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := helper.ParseResponseDataAndMessage(jsonPost, SuccessfulFindPostByTopicMessage)
	helper.Write(w, response)
}

// FindPostByID handles GET /posts/{id} requests.
// It parses the id string, and passes it to the post service to return the specified post,
// which then serializes the result into a JSON HTTP response.
func (h *handler) FindPostByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, InvalidPostIdMessage, http.StatusBadRequest)
		return
	}

	post, err := h.service.FindPostByID(r.Context(), id)
	if err != nil {
		if err == ErrPostNotFound {
			http.Error(w, ErrPostNotFound.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonPost, err := json.Marshal(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := helper.ParseResponseDataAndMessage(jsonPost, SuccessfulFindPostByIdMessage)
	helper.Write(w, response)
}

// CreatePost handles POST /posts requests.
// It reads and validates the request body, and passes it to the post service to create the new
// post with a unique title and description. It then serializes the result into a JSON HTTP response.
func (h *handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var req CreatePostRequest
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

	newPost := repo.CreatePostParams{
		TopicID:     req.TopicID,
		UserID:      req.UserID,
		Title:       req.Title,
		Description: req.Description,
	}
	post, err := h.service.CreatePost(r.Context(), newPost)
	if err != nil {
		if err == ErrPostAlreadyExists {
			http.Error(w, ErrPostAlreadyExists.Error(), http.StatusConflict)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonPost, err := json.Marshal(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := helper.ParseResponseDataAndMessage(jsonPost, SuccessfulCreatePostMessage)
	helper.Write(w, response)
}

// UpdatePost handles PUT /posts/{id} requests.
// It parses the id string, reads and validates the request body, and passes it to the post
// service to update the existing post. It then serializes the result into a JSON HTTP response.
func (h *handler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, InvalidPostIdMessage, http.StatusBadRequest)
		return
	}

	var req UpdatePostRequest
	err = helper.Read(r, &req)
	if err != nil {
		http.Error(w, InvalidRequestBodyMessage, http.StatusBadRequest)
		return
	}

	if !isValidPostReqBody(req) {
		http.Error(w, InvalidRequestBodyMessage, http.StatusBadRequest)
		return
	}

	newPost := repo.UpdatePostParams{
		PostID:      id,
		Title:       req.Title,
		Description: req.Description,
	}
	post, err := h.service.UpdatePost(r.Context(), newPost)
	if err != nil {
		if err == InvalidRequstBody {
			http.Error(w, InvalidRequstBody.Error(), http.StatusBadRequest)
			return
		}
		if err == ErrPostNotFound {
			http.Error(w, ErrPostNotFound.Error(), http.StatusNotFound)
			return
		}
		if err == ErrPostAlreadyExists {
			http.Error(w, ErrPostAlreadyExists.Error(), http.StatusConflict)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonPost, err := json.Marshal(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := helper.ParseResponseDataAndMessage(jsonPost, SuccessfulUpdatePostMessage)
	helper.Write(w, response)
}

// isValidPostReqBody returns true if both title and description of the request body are not empty.
func isValidPostReqBody(data UpdatePostRequest) bool {
	if data.Title == "" && data.Description == "" {
		return false
	}
	return true
}

// DeletePost handles DELETE /posts/{id} requests.
// It parses the id string, and passes it to the post service to delete the specified post,
// which then serializes the result into a JSON HTTP response.
func (h *handler) DeletePost(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, InvalidPostIdMessage, http.StatusBadRequest)
		return
	}

	err = h.service.DeletePost(r.Context(), id)
	if err != nil {
		if err == ErrPostNotFound {
			http.Error(w, ErrPostNotFound.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := helper.ParseResponseMessage(SuccessfulDeletePostMessage)
	helper.Write(w, response)
}
