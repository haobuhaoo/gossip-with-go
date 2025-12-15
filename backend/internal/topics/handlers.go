package topics

import (
	"encoding/json"
	"net/http"

	"github.com/haobuhaoo/gossip-with-go/internal/api"
	"github.com/haobuhaoo/gossip-with-go/internal/helpers"
)

const (
	SuccessFulListTopicMessage = "Successfully listed all topics"
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

// ListTopics handles GET /topics endpoints.
// It calls the topic service to return all topics and serializes the result into a JSON HTTP
// response.
func (h *handler) ListTopics(w http.ResponseWriter, r *http.Request) {
	topics, err := h.service.ListTopics(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonTopic, err := json.Marshal(topics)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := api.Response{
		Payload: api.Payload{
			Data: jsonTopic,
		},
		Messages: []string{SuccessFulListTopicMessage},
	}

	helpers.Write(w, response)
}
