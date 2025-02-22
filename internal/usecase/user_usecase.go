package usecase

import (
	"github.com/kei3dev/todo-app-api-go/internal/entity"
	"github.com/kei3dev/todo-app-api-go/internal/repository"
)

type UserUsecase interface {
	RegisterUser(user *entity.User) error
	GetUserByID(id uint) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
}

type userUsecaseImpl struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecaseImpl{userRepo: userRepo}
}

func (u *userUsecaseImpl) RegisterUser(user *entity.User) error {
	// ここでパスワードのハッシュ化などの処理を行う（仮実装）
	return u.userRepo.Create(user)
}

func (u *userUsecaseImpl) GetUserByID(id uint) (*entity.User, error) {
	return u.userRepo.FindByID(id)
}

func (u *userUsecaseImpl) GetUserByEmail(email string) (*entity.User, error) {
	return u.userRepo.FindByEmail(email)
}
