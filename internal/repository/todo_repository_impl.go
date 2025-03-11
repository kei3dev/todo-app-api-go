package repository

import (
	"time"

	"github.com/kei3dev/todo-app-api-go/internal/entity"
	"github.com/kei3dev/todo-app-api-go/pkg/db"

	"gorm.io/gorm"
)

type todoRepositoryImpl struct {
	db *gorm.DB
}

func NewTodoRepository() TodoRepository {
	return &todoRepositoryImpl{db: db.DB}
}

func (r *todoRepositoryImpl) Create(todo *entity.Todo) error {
	return r.db.Create(todo).Error
}

func (r *todoRepositoryImpl) FindByID(id uint) (*entity.Todo, error) {
	var todo entity.Todo
	err := r.db.First(&todo, id).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *todoRepositoryImpl) FindAllByUserID(userID uint) ([]entity.Todo, error) {
	var todos []entity.Todo
	err := r.db.Where("user_id = ?", userID).Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *todoRepositoryImpl) Update(todo *entity.Todo) error {
	return r.db.Model(todo).
		Updates(map[string]interface{}{
			"title":      todo.Title,
			"completed":  todo.Completed,
			"updated_at": time.Now(),
		}).Error
}

func (r *todoRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&entity.Todo{}, id).Error
}
