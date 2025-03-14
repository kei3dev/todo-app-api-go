package handler

import (
	"net/http"

	"github.com/kei3dev/todo-app-api-go/internal/handler/utils"
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

	if err := utils.DecodeRequestBody(r, &req); err != nil {
		utils.RespondWithError(w, err, http.StatusBadRequest)
		return
	}

	if err := utils.ValidateUserRegistration(req.Name, req.Email, req.Password); err != nil {
		utils.RespondWithError(w, err, http.StatusBadRequest)
		return
	}

	userDTO := &usecase.UserDTO{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	err := h.UserUsecase.RegisterUser(userDTO)
	if err != nil {
		utils.RespondWithError(w, utils.ErrRegisterUserFailed, http.StatusInternalServerError)
		return
	}

	responseUser := map[string]interface{}{
		"name":  req.Name,
		"email": req.Email,
	}

	utils.RespondWithJSON(w, responseUser, http.StatusCreated)
}
