package topics

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/haobuhaoo/gossip-with-go/internal/helper"
	repo "github.com/haobuhaoo/gossip-with-go/internal/postgresql/sqlc"
)

const (
	InvalidTopicIdMessage        = "Invalid topic id"
	InvalidRequestBodyMessage    = "Required fields missing"
	SuccessfulListTopicMessage   = "Successfully listed all topics"
	SuccessfulFindTopicMessage   = "Successfully find topic"
	SuccessfulCreateTopicMessage = "Successfully created topic"
	SuccessfulUpdateTopicMessage = "Successfully updated topic"
	SuccessfulDeleteTopicMessage = "Successfully deleted topic"
)

// handler handles the topic related HTTP requests.
// It is responsible for translating HTTP requests into service calls and formatting service
// responses into HTTP responses.
type handler struct {
	service Service
}

// NewHandler creates a new topic handler.
func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

// ListTopics handles GET /topics requests.
// It calls the topic service to return all topics and serializes the result into a JSON HTTP
// response.
func (h *handler) ListTopics(w http.ResponseWriter, r *http.Request) {
	topics, err := h.service.ListTopics(r.Context())
	if err != nil {
		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonTopic, err := json.Marshal(topics)
	if err != nil {
		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := helper.ParseResponseDataAndMessage(jsonTopic, SuccessfulListTopicMessage)
	helper.Write(w, response)
}

// FindTopicByID handles GET /topics/{id} requests.
// It parses the id string, and passes it to the topic service to return the specified topic,
// which then serializes the result into a JSON HTTP response.
func (h *handler) FindTopicByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		helper.WriteError(w, InvalidTopicIdMessage, http.StatusBadRequest)
		return
	}

	topic, err := h.service.FindTopicByID(r.Context(), id)
	if err != nil {
		if err == ErrTopicNotFound {
			helper.WriteError(w, ErrTopicNotFound.Error(), http.StatusNotFound)
			return
		}

		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonTopic, err := json.Marshal(topic)
	if err != nil {
		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := helper.ParseResponseDataAndMessage(jsonTopic, SuccessfulFindTopicMessage)
	helper.Write(w, response)
}

// CreateTopic handles POST /topics requests.
// It reads and validates the request body, and passes it to the topic service to create the new
// topic with a unique title. It then serializes the result into a JSON HTTP response.
func (h *handler) CreateTopic(w http.ResponseWriter, r *http.Request) {
	var req CreateTopicRequest
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

	newTopic := repo.CreateTopicParams{
		UserID: req.UserID,
		Title:  req.Title,
	}
	topic, err := h.service.CreateTopic(r.Context(), newTopic)
	if err != nil {
		if err == ErrTopicAlreadyExists {
			helper.WriteError(w, ErrTopicAlreadyExists.Error(), http.StatusConflict)
			return
		}

		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonTopic, err := json.Marshal(topic)
	if err != nil {
		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := helper.ParseResponseDataAndMessage(jsonTopic, SuccessfulCreateTopicMessage)
	helper.Write(w, response)
}

// UpdateTopic handles PUT /topics/{id} requests.
// It parses the id string, reads and validates the request body, and passes it to the topic
// service to update the existing topic with the new title. It then serializes the result into a
// JSON HTTP response.
func (h *handler) UpdateTopic(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		helper.WriteError(w, InvalidTopicIdMessage, http.StatusBadRequest)
		return
	}

	var req UpdateTopicRequest
	err = helper.Read(r, &req)
	if err != nil {
		helper.WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validator.New().Struct(req)
	if err != nil {
		helper.WriteError(w, InvalidRequestBodyMessage, http.StatusBadRequest)
		return
	}

	newTopic := repo.UpdateTopicParams{
		TopicID: id,
		Title:   req.Title,
	}
	topic, err := h.service.UpdateTopic(r.Context(), newTopic)
	if err != nil {
		if err == ErrTopicNotFound {
			helper.WriteError(w, ErrTopicNotFound.Error(), http.StatusNotFound)
			return
		}
		if err == ErrTopicAlreadyExists {
			helper.WriteError(w, ErrTopicAlreadyExists.Error(), http.StatusConflict)
			return
		}

		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonTopic, err := json.Marshal(topic)
	if err != nil {
		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := helper.ParseResponseDataAndMessage(jsonTopic, SuccessfulUpdateTopicMessage)
	helper.Write(w, response)
}

// DeleteTopic handles DELETE /topics/{id} requests.
// It parses the id string, and passes it to the topic service to delete the specified topic,
// which then serializes the result into a JSON HTTP response.
func (h *handler) DeleteTopic(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		helper.WriteError(w, InvalidTopicIdMessage, http.StatusBadRequest)
		return
	}

	err = h.service.DeleteTopic(r.Context(), id)
	if err != nil {
		if err == ErrTopicNotFound {
			helper.WriteError(w, ErrTopicNotFound.Error(), http.StatusNotFound)
			return
		}

		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := helper.ParseResponseMessage(SuccessfulDeleteTopicMessage)
	helper.Write(w, response)
}
