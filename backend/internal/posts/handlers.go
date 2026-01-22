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
	"github.com/jackc/pgx/v5/pgtype"
)

const (
	InvalidTopicIdMessage              = "Invalid topic id"
	InvalidPostIdMessage               = "Invalid post id"
	InvalidRequestBodyMessage          = "Required fields missing"
	InvalidQueryMessage                = "Query string missing"
	MissingUserIDMessage               = "Missing userID"
	SuccessfulFindPostByTopicMessage   = "Successfully listed all posts"
	SuccessfulFindPostByIdMessage      = "Successfully find post"
	SuccessfulCreatePostMessage        = "Successfully created post"
	SuccessfulUpdatePostMessage        = "Successfully updated post"
	SuccessfulDeletePostMessage        = "Successfully deleted post"
	SuccessfulSearchPostByTopicMessage = "Successfully searched post"
	SuccessfulLikePostMessage          = "Successfully liked post"
	SuccessfulDislikePostMessage       = "Successfully disliked post"
	SuccessfulRemovePostVoteMessage    = "Successfully removed vote"
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

// FindPostsByTopic handles GET /api/posts/all/{topicId} requests.
// It parses the topicId string, and passes it to the post service to return all posts for that topic
// and serializes the result into a JSON HTTP response.
func (h *handler) FindPostsByTopic(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "topicId")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		helper.WriteError(w, topics.InvalidTopicIdMessage, http.StatusBadRequest)
		return
	}

	userId, ok := r.Context().Value("userID").(int64)
	if !ok {
		helper.WriteError(w, MissingUserIDMessage, http.StatusBadRequest)
		return
	}

	arg := repo.FindPostsByTopicParams{
		TopicID: id,
		UserID:  userId,
	}
	posts, err := h.service.FindPostsByTopic(r.Context(), arg)
	if err != nil {
		if err == topics.ErrTopicNotFound {
			helper.WriteError(w, topics.InvalidTopicIdMessage, http.StatusBadRequest)
			return
		}

		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonPost, err := json.Marshal(posts)
	if err != nil {
		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := helper.ParseResponseDataAndMessage(jsonPost, SuccessfulFindPostByTopicMessage)
	helper.Write(w, response)
}

// FindPostByID handles GET /api/posts/{topicId}/{postId} requests.
// It parses the topicId and postId string, and passes it to the post service to return the
// specified post, which then serializes the result into a JSON HTTP response.
func (h *handler) FindPostByID(w http.ResponseWriter, r *http.Request) {
	topicIdStr := chi.URLParam(r, "topicId")
	topicId, err := strconv.ParseInt(topicIdStr, 10, 64)
	if err != nil {
		helper.WriteError(w, InvalidTopicIdMessage, http.StatusBadRequest)
		return
	}

	postIdStr := chi.URLParam(r, "postId")
	postId, err := strconv.ParseInt(postIdStr, 10, 64)
	if err != nil {
		helper.WriteError(w, InvalidPostIdMessage, http.StatusBadRequest)
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
	post, err := h.service.FindPostByID(r.Context(), req)
	if err != nil {
		if err == ErrPostNotFound {
			helper.WriteError(w, ErrPostNotFound.Error(), http.StatusNotFound)
			return
		}

		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonPost, err := json.Marshal(post)
	if err != nil {
		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := helper.ParseResponseDataAndMessage(jsonPost, SuccessfulFindPostByIdMessage)
	helper.Write(w, response)
}

// CreatePost handles POST /api/posts requests.
// It reads and validates the request body, and passes it to the post service to create the new
// post with a unique title and description. It then serializes the result into a JSON HTTP response.
func (h *handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var req CreatePostRequest
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

	newPost := repo.CreatePostParams{
		TopicID:     req.TopicID,
		UserID:      userId,
		Title:       req.Title,
		Description: req.Description,
	}
	post, err := h.service.CreatePost(r.Context(), newPost)
	if err != nil {
		if err == ErrPostAlreadyExists {
			helper.WriteError(w, ErrPostAlreadyExists.Error(), http.StatusConflict)
			return
		}

		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonPost, err := json.Marshal(post)
	if err != nil {
		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := helper.ParseResponseDataAndMessage(jsonPost, SuccessfulCreatePostMessage)
	helper.Write(w, response)
}

// UpdatePost handles PUT /api/posts/{id} requests.
// It parses the id string, reads and validates the request body, and passes it to the post
// service to update the existing post. It then serializes the result into a JSON HTTP response.
func (h *handler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		helper.WriteError(w, InvalidPostIdMessage, http.StatusBadRequest)
		return
	}

	var req UpdatePostRequest
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

	newPost := repo.UpdatePostParams{
		PostID:      id,
		UserID:      userId,
		Title:       req.Title,
		Description: req.Description,
	}
	post, err := h.service.UpdatePost(r.Context(), newPost)
	if err != nil {
		if err == ErrPostNotFound {
			helper.WriteError(w, ErrPostNotFound.Error(), http.StatusNotFound)
			return
		}
		if err == ErrPostAlreadyExists {
			helper.WriteError(w, ErrPostAlreadyExists.Error(), http.StatusConflict)
			return
		}

		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonPost, err := json.Marshal(post)
	if err != nil {
		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := helper.ParseResponseDataAndMessage(jsonPost, SuccessfulUpdatePostMessage)
	helper.Write(w, response)
}

// DeletePost handles DELETE /api/posts/{id} requests.
// It parses the id string, and passes it to the post service to delete the specified post,
// which then serializes the result into a JSON HTTP response.
func (h *handler) DeletePost(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		helper.WriteError(w, InvalidPostIdMessage, http.StatusBadRequest)
		return
	}

	userId, ok := r.Context().Value("userID").(int64)
	if !ok {
		helper.WriteError(w, MissingUserIDMessage, http.StatusBadRequest)
		return
	}

	arg := repo.DeletePostParams{
		PostID: id,
		UserID: userId,
	}
	err = h.service.DeletePost(r.Context(), arg)
	if err != nil {
		if err == ErrPostNotFound {
			helper.WriteError(w, ErrPostNotFound.Error(), http.StatusNotFound)
			return
		}

		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := helper.ParseResponseMessage(SuccessfulDeletePostMessage)
	helper.Write(w, response)
}

// SearchPost handles GET /api/posts/{topicId}/search requests.
// It parses the topicId and query string, and passes it to the post service to search for all
// post titles and descriptions under the specified topic that contains the query string
// (case-insensitive), which then serializes the result into a JSON HTTP response.
func (h *handler) SearchPost(w http.ResponseWriter, r *http.Request) {
	topicIdStr := chi.URLParam(r, "topicId")
	topicId, err := strconv.ParseInt(topicIdStr, 10, 64)
	if err != nil {
		helper.WriteError(w, InvalidTopicIdMessage, http.StatusBadRequest)
		return
	}

	rawQuery := r.URL.Query().Get("q")
	if rawQuery == "" {
		helper.WriteError(w, InvalidQueryMessage, http.StatusBadRequest)
		return
	}

	userId, ok := r.Context().Value("userID").(int64)
	if !ok {
		helper.WriteError(w, MissingUserIDMessage, http.StatusBadRequest)
		return
	}

	arg := repo.SearchPostParams{
		TopicID: topicId,
		Column2: pgtype.Text{
			String: rawQuery,
			Valid:  true,
		},
		UserID: userId,
	}
	posts, err := h.service.SearchPost(r.Context(), arg)
	if err != nil {
		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonPost, err := json.Marshal(posts)
	if err != nil {
		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := helper.ParseResponseDataAndMessage(jsonPost, SuccessfulSearchPostByTopicMessage)
	helper.Write(w, response)
}

// LikesPost handles POST /api/posts/{id}/likes requests.
// It parses the id string and passes it to the post service to increment a like count for that
// specified post, which then serializes the result into a JSON HTTP response.
func (h *handler) LikesPost(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		helper.WriteError(w, InvalidPostIdMessage, http.StatusBadRequest)
		return
	}

	userId, ok := r.Context().Value("userID").(int64)
	if !ok {
		helper.WriteError(w, MissingUserIDMessage, http.StatusBadRequest)
		return
	}

	arg := repo.LikesPostParams{
		PostID: id,
		UserID: userId,
	}
	err = h.service.LikesPost(r.Context(), arg)
	if err != nil {
		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := helper.ParseResponseMessage(SuccessfulLikePostMessage)
	helper.Write(w, response)
}

// DislikesPost handles POST /api/posts/{id}/dislikes requests.
// It parses the id string and passes it to the post service to increment a dislike count for that
// specified post, which then serializes the result into a JSON HTTP response.
func (h *handler) DislikesPost(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		helper.WriteError(w, InvalidPostIdMessage, http.StatusBadRequest)
		return
	}

	userId, ok := r.Context().Value("userID").(int64)
	if !ok {
		helper.WriteError(w, MissingUserIDMessage, http.StatusBadRequest)
		return
	}

	arg := repo.DislikesPostParams{
		PostID: id,
		UserID: userId,
	}
	err = h.service.DislikesPost(r.Context(), arg)
	if err != nil {
		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := helper.ParseResponseMessage(SuccessfulDislikePostMessage)
	helper.Write(w, response)
}

// RemovePostVote handles DELETE /api/posts/{id}/remove requests.
// It parses the id string and passes it to the post service to remove the vote count for that
// specified post, which then serializes the result into a JSON HTTP response.
func (h *handler) RemovePostVote(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		helper.WriteError(w, InvalidPostIdMessage, http.StatusBadRequest)
		return
	}

	userId, ok := r.Context().Value("userID").(int64)
	if !ok {
		helper.WriteError(w, MissingUserIDMessage, http.StatusBadRequest)
		return
	}

	arg := repo.RemovePostVoteParams{
		PostID: id,
		UserID: userId,
	}
	err = h.service.RemovePostVote(r.Context(), arg)
	if err != nil {
		if err == ErrVoteNotFound {
			helper.WriteError(w, ErrVoteNotFound.Error(), http.StatusNotFound)
			return
		}

		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := helper.ParseResponseMessage(SuccessfulRemovePostVoteMessage)
	helper.Write(w, response)
}
