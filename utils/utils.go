package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ErrorResponse represents the JSON structure for error messages.
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// ParseJson
func ParseJson(r *http.Request, payload any) error {

	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}

// for write json
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// for json error
func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}
