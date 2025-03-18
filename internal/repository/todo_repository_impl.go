package repository

import (
	"github.com/kei3dev/todo-app-api-go/internal/entity"
	"github.com/kei3dev/todo-app-api-go/pkg/db"
)

type todoRepositoryImpl struct {
	BaseRepository
}

func NewTodoRepository() TodoRepository {
	return &todoRepositoryImpl{
		BaseRepository: NewBaseRepository(db.DB),
	}
}

func (r *todoRepositoryImpl) Create(todo *entity.Todo) error {
	return r.DB.Create(todo).Error
}

func (r *todoRepositoryImpl) FindByID(id uint) (*entity.Todo, error) {
	var todo entity.Todo
	err := r.DB.First(&todo, id).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *todoRepositoryImpl) FindAllByUserID(userID uint) ([]entity.Todo, error) {
	var todos []entity.Todo
	err := r.DB.Where("user_id = ?", userID).Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *todoRepositoryImpl) Update(todo *entity.Todo) error {
	return r.DB.Model(todo).Updates(map[string]any{
		"title":     todo.Title,
		"completed": todo.Completed,
	}).Error
}

func (r *todoRepositoryImpl) Delete(id uint) error {
	return r.DB.Delete(&entity.Todo{}, id).Error
}
