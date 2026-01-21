package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/haobuhaoo/gossip-with-go/internal/helper"
)

const (
	InvalidRequestBodyMessage         = "Required fields missing"
	InvalidCredentialMessage          = "Invalid credentials"
	FailedTokenGenerationMessage      = "Failed to generate token"
	MissingUserIDMessage              = "Missing userID"
	SuccessfulLoginMessage            = "Successfully login"
	SuccessfulAuthenticateUserMessage = "Successfully authenticate user"
)

// handler handles the authentication related HTTP requests.
// It is responsible for translating HTTP requests into service calls and formatting service
// responses into HTTP responses.
type handler struct {
	service   Service
	jwtSecret []byte
}

// NewHandler creates a new authentication handler.
func NewHandler(service Service, secret string) *handler {
	return &handler{
		service:   service,
		jwtSecret: []byte(secret),
	}
}

// Login handles POST /auth/login requests.
// It reads and validates the request body, and passes it to the authentication service to login
// the user. It then generates a JWT token with the userID that expires in 12 hours, and serializes
// it into a JSON HTTP response.
func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
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

	user, err := h.service.Login(r.Context(), req.Username)
	if err != nil {
		if err == ErrUserNotFound {
			helper.WriteError(w, InvalidCredentialMessage, http.StatusUnauthorized)
			return
		}
		helper.WriteError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	claims := jwt.MapClaims{
		"user_id": user.UserID,
		"exp":     time.Now().Add(12 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(h.jwtSecret)
	if err != nil {
		helper.WriteError(w, FailedTokenGenerationMessage, http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.Marshal(LoginResponse{Token: signedToken})
	if err != nil {
		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := helper.ParseResponseDataAndMessage(jsonResp, SuccessfulLoginMessage)
	helper.Write(w, response)
}

// AuthenticateUser handles GET /api/me requests.
// It gets the userId from the JWT context, and passes it to the authentication service to verify the
// user and returns it. It then serializes the result into a JSON HTTP response.
func (h *handler) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userID").(int64)
	if !ok {
		helper.WriteError(w, MissingUserIDMessage, http.StatusBadRequest)
		return
	}

	user, err := h.service.AuthenticateUser(r.Context(), userId)
	if err != nil {
		if err == ErrUserNotFound {
			helper.WriteError(w, ErrUserNotFound.Error(), http.StatusNotFound)
			return
		}

		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := helper.ParseResponseDataAndMessage(jsonUser, SuccessfulAuthenticateUserMessage)
	helper.Write(w, response)
}
