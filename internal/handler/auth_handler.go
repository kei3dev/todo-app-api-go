package handler

import (
	"net/http"

	"github.com/kei3dev/todo-app-api-go/internal/errors"
	"github.com/kei3dev/todo-app-api-go/internal/handler/utils"
	"github.com/kei3dev/todo-app-api-go/internal/usecase"
	"github.com/kei3dev/todo-app-api-go/pkg/middleware"
)

type AuthHandler struct {
	UserUsecase usecase.UserUsecase
	JWTConfig   *middleware.JWTConfig
}

func NewAuthHandler(userUsecase usecase.UserUsecase, jwtConfig *middleware.JWTConfig) *AuthHandler {
	return &AuthHandler{
		UserUsecase: userUsecase,
		JWTConfig:   jwtConfig,
	}
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

	if err := utils.ValidateLogin(req.Email, req.Password); err != nil {
		utils.RespondWithError(w, err, http.StatusBadRequest)
		return
	}

	user, err := h.UserUsecase.VerifyPassword(req.Email, req.Password)
	if err != nil {
		utils.RespondWithError(w, errors.ErrInvalidCredentials, http.StatusUnauthorized)
		return
	}

	token, err := h.JWTConfig.GenerateJWT(user)
	if err != nil {
		utils.RespondWithError(w, errors.ErrTokenGenerationFailed, http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, map[string]string{"token": token}, http.StatusOK)
}
