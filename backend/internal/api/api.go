package api

import (
	"encoding/json"
)

// Payload represents an API data format.
type Payload struct {
	Meta json.RawMessage `json:"meta,omitempty"`
	Data json.RawMessage `json:"data,omitempty"`
}

// Response represents an API response format.
type Response struct {
	Payload   Payload  `json:"payload"`
	Messages  []string `json:"messages"`
	ErrorCode int      `json:"errorCode"`
}