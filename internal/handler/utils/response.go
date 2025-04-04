package utils

import (
	"encoding/json"
	"net/http"

	"github.com/kei3dev/todo-app-api-go/internal/errors"
)

func DecodeRequestBody(r *http.Request, v any) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return errors.ErrInvalidRequestPayload
	}
	return nil
}

func RespondWithError(w http.ResponseWriter, err error, statusCode int) {
	http.Error(w, err.Error(), statusCode)
}

func RespondWithJSON(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
