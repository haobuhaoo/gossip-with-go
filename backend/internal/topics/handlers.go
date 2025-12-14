package topics

import (
	"encoding/json"
	"net/http"

	"github.com/haobuhaoo/gossip-with-go/internal/api"
	"github.com/haobuhaoo/gossip-with-go/internal/helpers"
)

const (
	SuccessFulListTopicMessage = "Successfully listed topics"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

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
