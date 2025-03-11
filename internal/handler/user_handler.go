package handler

import (
	"encoding/json"
	"net/http"

	"github.com/kei3dev/todo-app-api-go/internal/usecase"
)

type UserHandler struct {
	UserUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{UserUsecase: userUsecase}
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	userDTO := &usecase.UserDTO{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	err := h.UserUsecase.RegisterUser(userDTO)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	responseUser := map[string]interface{}{
		"name":  req.Name,
		"email": req.Email,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responseUser)
}
