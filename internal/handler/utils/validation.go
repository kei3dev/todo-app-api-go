package utils

import (
	"strings"

	"github.com/kei3dev/todo-app-api-go/internal/entity"
	"github.com/kei3dev/todo-app-api-go/internal/errors"
)

func ValidateTodo(todo *entity.Todo) error {
	if strings.TrimSpace(todo.Title) == "" {
		return errors.ErrEmptyTitle
	}
	return nil
}

func ValidateUserRegistration(name, email, password string) error {
	if strings.TrimSpace(name) == "" {
		return errors.ErrEmptyName
	}

	if err := ValidateEmail(email); err != nil {
		return err
	}

	if err := ValidatePassword(password); err != nil {
		return err
	}

	return nil
}

func ValidateLogin(email, password string) error {
	if err := ValidateEmail(email); err != nil {
		return err
	}

	if strings.TrimSpace(password) == "" {
		return errors.ErrEmptyPassword
	}

	return nil
}

func ValidateEmail(email string) error {
	if strings.TrimSpace(email) == "" {
		return errors.ErrEmptyEmail
	}

	if !strings.Contains(email, "@") {
		return errors.ErrInvalidEmailFormat
	}

	return nil
}

func ValidatePassword(password string) error {
	if strings.TrimSpace(password) == "" {
		return errors.ErrEmptyPassword
	}

	if len(password) < 6 {
		return errors.ErrPasswordTooShort
	}

	return nil
}
