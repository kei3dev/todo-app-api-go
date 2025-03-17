package repository

import (
	"github.com/kei3dev/todo-app-api-go/internal/entity"
	"github.com/kei3dev/todo-app-api-go/pkg/db"
)

type userRepositoryImpl struct {
	BaseRepository
}

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{
		BaseRepository: NewBaseRepository(db.DB),
	}
}

func (r *userRepositoryImpl) Create(user *entity.User) error {
	return r.DB.Create(user).Error
}

func (r *userRepositoryImpl) FindByID(id uint) (*entity.User, error) {
	var user entity.User
	err := r.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
