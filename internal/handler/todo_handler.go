package handler

import (
	"net/http"

	"github.com/kei3dev/todo-app-api-go/internal/entity"
	"github.com/kei3dev/todo-app-api-go/internal/handler/utils"
	"github.com/kei3dev/todo-app-api-go/internal/usecase"
)

type TodoHandler struct {
	TodoUsecase usecase.TodoUsecase
}

func NewTodoHandler(todoUsecase usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{TodoUsecase: todoUsecase}
}

func (h *TodoHandler) checkTodoOwnership(todoID uint, userID uint) (*entity.Todo, error) {
	todo, err := h.TodoUsecase.GetTodoByID(todoID)
	if err != nil {
		return nil, utils.ErrTodoNotFound
	}

	if todo.UserID != userID {
		return nil, utils.ErrUnauthorized
	}

	return todo, nil
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	userID, err := utils.GetUserIDFromContext(r)
	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}

	var todo entity.Todo
	if err := utils.DecodeRequestBody(r, &todo); err != nil {
		utils.RespondWithError(w, err, http.StatusBadRequest)
		return
	}

	if err := utils.ValidateTodo(&todo); err != nil {
		utils.RespondWithError(w, err, http.StatusBadRequest)
		return
	}

	todo.UserID = userID

	if err := h.TodoUsecase.CreateTodo(&todo); err != nil {
		utils.RespondWithError(w, utils.ErrCreateTodoFailed, http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, todo, http.StatusCreated)
}

func (h *TodoHandler) GetTodoByID(w http.ResponseWriter, r *http.Request) {
	todoID, err := utils.GetIDFromURL(r, "id")
	if err != nil {
		utils.RespondWithError(w, err, http.StatusBadRequest)
		return
	}

	todo, err := h.TodoUsecase.GetTodoByID(todoID)
	if err != nil {
		utils.RespondWithError(w, utils.ErrTodoNotFound, http.StatusNotFound)
		return
	}

	utils.RespondWithJSON(w, todo, http.StatusOK)
}

func (h *TodoHandler) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	userID, err := utils.GetUserIDFromContext(r)
	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}

	todos, err := h.TodoUsecase.GetTodosByUserID(userID)
	if err != nil {
		utils.RespondWithError(w, utils.ErrGetTodosFailed, http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, todos, http.StatusOK)
}

func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	userID, err := utils.GetUserIDFromContext(r)
	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}

	todoID, err := utils.GetIDFromURL(r, "id")
	if err != nil {
		utils.RespondWithError(w, err, http.StatusBadRequest)
		return
	}

	existingTodo, err := h.checkTodoOwnership(todoID, userID)
	if err != nil {
		statusCode := http.StatusNotFound
		if err == utils.ErrUnauthorized {
			statusCode = http.StatusForbidden
		}
		utils.RespondWithError(w, err, statusCode)
		return
	}

	var updatedTodo entity.Todo
	if err := utils.DecodeRequestBody(r, &updatedTodo); err != nil {
		utils.RespondWithError(w, err, http.StatusBadRequest)
		return
	}

	if err := utils.ValidateTodo(&updatedTodo); err != nil {
		utils.RespondWithError(w, err, http.StatusBadRequest)
		return
	}

	updatedTodo.ID = todoID
	updatedTodo.UserID = userID
	updatedTodo.CreatedAt = existingTodo.CreatedAt

	if err := h.TodoUsecase.UpdateTodo(&updatedTodo); err != nil {
		utils.RespondWithError(w, utils.ErrUpdateTodoFailed, http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, updatedTodo, http.StatusOK)
}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	userID, err := utils.GetUserIDFromContext(r)
	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}

	todoID, err := utils.GetIDFromURL(r, "id")
	if err != nil {
		utils.RespondWithError(w, err, http.StatusBadRequest)
		return
	}

	_, err = h.checkTodoOwnership(todoID, userID)
	if err != nil {
		statusCode := http.StatusNotFound
		if err == utils.ErrUnauthorized {
			statusCode = http.StatusForbidden
		}
		utils.RespondWithError(w, err, statusCode)
		return
	}

	if err := h.TodoUsecase.DeleteTodo(todoID); err != nil {
		utils.RespondWithError(w, utils.ErrDeleteTodoFailed, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
