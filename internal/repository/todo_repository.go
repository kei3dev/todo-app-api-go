package repository

import "github.com/kei3dev/todo-app-api-go/internal/entity"

type TodoRepository interface {
	IRepository
	Create(todo *entity.Todo) error
	FindByID(id uint) (*entity.Todo, error)
	FindAllByUserID(userID uint) ([]entity.Todo, error)
	Update(todo *entity.Todo) error
	Delete(id uint) error
}
