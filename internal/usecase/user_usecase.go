package usecase

import (
	"errors"

	"github.com/kei3dev/todo-app-api-go/internal/entity"
	"github.com/kei3dev/todo-app-api-go/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserDTO struct {
	Name     string
	Email    string
	Password string
}

type UserUsecase interface {
	RegisterUser(dto *UserDTO) error
	GetUserByID(id uint) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	VerifyPassword(email, password string) (*entity.User, error)
}

type userUsecaseImpl struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecaseImpl{userRepo: userRepo}
}

func (u *userUsecaseImpl) RegisterUser(dto *UserDTO) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &entity.User{
		Name:         dto.Name,
		Email:        dto.Email,
		PasswordHash: string(hashedPassword),
	}

	return u.userRepo.Create(user)
}

func (u *userUsecaseImpl) GetUserByID(id uint) (*entity.User, error) {
	return u.userRepo.FindByID(id)
}

func (u *userUsecaseImpl) GetUserByEmail(email string) (*entity.User, error) {
	return u.userRepo.FindByEmail(email)
}

func (u *userUsecaseImpl) VerifyPassword(email, password string) (*entity.User, error) {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
