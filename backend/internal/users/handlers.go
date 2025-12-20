package users

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/haobuhaoo/gossip-with-go/internal/helper"
)

const (
	InvalidRequestBodyMessage   = "Required fields missing"
	SuccessfulFindUserMessage   = "Successfully find user"
	SuccessfulCreateUserMessage = "Successfully created user"
)

// handler handles the user related HTTP requests.
// It is responsible for translating HTTP requests into service calls and formatting service
// responses into HTTP responses.
type handler struct {
	service Service
}

// NewHandler creates a new user handler.
func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

// FindUserByName handles GET /users/{name} requests.
// It passes the name string to the user service to return the specified user, which then
// serializes the result into a JSON HTTP response.
func (h *handler) FindUserByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	user, err := h.service.FindUserByName(r.Context(), name)
	if err != nil {
		if err == ErrUserNotFound {
			http.Error(w, ErrUserNotFound.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := helper.ParseResponseDataAndMessage(jsonUser, SuccessfulFindUserMessage)
	helper.Write(w, response)
}

// CreateUser handles POST /users requests.
// It reads and validates the request body, and passes it to the user service to create the new
// user with a unique name. It then serializes the result into a JSON HTTP response.
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
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

	user, err := h.service.CreateUser(r.Context(), req.Name)
	if err != nil {
		if err == ErrUserAlreadyExists {
			http.Error(w, ErrUserAlreadyExists.Error(), http.StatusConflict)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := helper.ParseResponseDataAndMessage(jsonUser, SuccessfulCreateUserMessage)
	helper.Write(w, response)
}
