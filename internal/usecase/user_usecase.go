package usecase

import (
	"fmt"
	"regexp"

	"github.com/kei3dev/todo-app-api-go/internal/entity"
	"github.com/kei3dev/todo-app-api-go/internal/errors"
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
	if err := validateUserDTO(dto); err != nil {
		return err
	}

	existingUser, err := u.userRepo.FindByEmail(dto.Email)
	if err == nil && existingUser != nil {
		return errors.ErrEmailAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
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
		return nil, errors.ErrInvalidCredentials
	}

	return user, nil
}

func validateUserDTO(dto *UserDTO) error {
	if len(dto.Name) < 2 {
		return errors.ErrNameTooShort
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(dto.Email) {
		return errors.ErrInvalidEmailFormat
	}

	if len(dto.Password) < 8 {
		return errors.ErrPasswordTooShort
	}

	return nil
}
