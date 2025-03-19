package utils

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kei3dev/todo-app-api-go/internal/errors"
	"github.com/kei3dev/todo-app-api-go/pkg/middleware"
)

func GetUserIDFromContext(r *http.Request) (uint, error) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(uint)
	if !ok {
		return 0, errors.ErrUserIDNotFound
	}
	return userID, nil
}

func GetIDFromURL(r *http.Request, paramName string) (uint, error) {
	idParam := chi.URLParam(r, paramName)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return 0, errors.ErrInvalidID
	}
	return uint(id), nil
}
