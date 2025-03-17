package repository

import "github.com/kei3dev/todo-app-api-go/internal/entity"

type UserRepository interface {
	IRepository
	Create(user *entity.User) error
	FindByID(id uint) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
}
