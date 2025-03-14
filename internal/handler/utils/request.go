package utils

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kei3dev/todo-app-api-go/pkg/middleware"
)

func GetUserIDFromContext(r *http.Request) (uint, error) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(uint)
	if !ok {
		return 0, ErrUserIDNotFound
	}
	return userID, nil
}

func GetIDFromURL(r *http.Request, paramName string) (uint, error) {
	idParam := chi.URLParam(r, paramName)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return 0, ErrInvalidID
	}
	return uint(id), nil
}
