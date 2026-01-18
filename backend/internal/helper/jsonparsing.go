package helper

import (
	"encoding/json"
	"net/http"
)

// Write encodes the data as JSON and writes it to the HTTP response.
func Write(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// Read decodes the JSON body of the HTTP request into the provided data structure.
// It disallows unknown fields to prevent unexpected inputs and returns any error encountered.
func Read(r *http.Request, data any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(data)
}

// WriteError encodes the error message and status code into a HTTP response.
func WriteError(w http.ResponseWriter, msg string, code int) {
	w.WriteHeader(code)
	response := ParseErrorResponseMessage(msg, code)
	Write(w, response)
}
