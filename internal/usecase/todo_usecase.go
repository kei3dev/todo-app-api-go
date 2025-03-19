package usecase

import (
	"fmt"

	"github.com/kei3dev/todo-app-api-go/internal/entity"
	"github.com/kei3dev/todo-app-api-go/internal/errors"
	"github.com/kei3dev/todo-app-api-go/internal/repository"
)

type TodoUsecase interface {
	CreateTodo(todo *entity.Todo) error
	GetTodoByID(id uint) (*entity.Todo, error)
	GetTodosByUserID(userID uint) ([]entity.Todo, error)
	UpdateTodo(todo *entity.Todo) error
	DeleteTodo(id uint) error
}

type todoUsecaseImpl struct {
	todoRepo repository.TodoRepository
}

func NewTodoUsecase(todoRepo repository.TodoRepository) TodoUsecase {
	return &todoUsecaseImpl{todoRepo: todoRepo}
}

func (u *todoUsecaseImpl) CreateTodo(todo *entity.Todo) error {
	return u.todoRepo.Create(todo)
}

func (u *todoUsecaseImpl) GetTodoByID(id uint) (*entity.Todo, error) {
	todo, err := u.todoRepo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find todo: %w", err)
	}
	if todo == nil {
		return nil, errors.ErrTodoNotFound
	}
	return todo, nil
}

func (u *todoUsecaseImpl) GetTodosByUserID(userID uint) ([]entity.Todo, error) {
	return u.todoRepo.FindAllByUserID(userID)
}

func (u *todoUsecaseImpl) UpdateTodo(todo *entity.Todo) error {
	existingTodo, err := u.todoRepo.FindByID(todo.ID)
	if err != nil {
		return fmt.Errorf("failed to find todo: %w", err)
	}
	if existingTodo == nil {
		return errors.ErrTodoNotFound
	}

	if existingTodo.UserID != todo.UserID {
		return errors.ErrUnauthorized
	}

	return u.todoRepo.Update(todo)
}

func (u *todoUsecaseImpl) DeleteTodo(id uint) error {
	todo, err := u.todoRepo.FindByID(id)
	if err != nil {
		return fmt.Errorf("failed to find todo: %w", err)
	}
	if todo == nil {
		return errors.ErrTodoNotFound
	}

	return u.todoRepo.Delete(id)
}
