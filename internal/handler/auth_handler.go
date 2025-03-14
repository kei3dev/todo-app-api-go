package handler

import (
	"net/http"

	"github.com/kei3dev/todo-app-api-go/internal/handler/utils"
	"github.com/kei3dev/todo-app-api-go/internal/usecase"
	"github.com/kei3dev/todo-app-api-go/pkg/middleware"
)

type AuthHandler struct {
	UserUsecase usecase.UserUsecase
}

func NewAuthHandler(userUsecase usecase.UserUsecase) *AuthHandler {
	return &AuthHandler{UserUsecase: userUsecase}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := utils.DecodeRequestBody(r, &req); err != nil {
		utils.RespondWithError(w, err, http.StatusBadRequest)
		return
	}

	user, err := h.UserUsecase.VerifyPassword(req.Email, req.Password)
	if err != nil {
		utils.RespondWithError(w, utils.ErrAuthenticationFailed, http.StatusUnauthorized)
		return
	}

	token, err := middleware.GenerateJWT(user)
	if err != nil {
		utils.RespondWithError(w, utils.ErrTokenGenerationFailed, http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, map[string]string{"token": token}, http.StatusOK)
}
