package utils

import (
	"strings"

	"github.com/kei3dev/todo-app-api-go/internal/entity"
)

func ValidateTodo(todo *entity.Todo) error {
	if strings.TrimSpace(todo.Title) == "" {
		return ErrEmptyTitle
	}
	return nil
}

func ValidateUserRegistration(name, email, password string) error {
	if strings.TrimSpace(name) == "" {
		return ErrEmptyName
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
		return ErrEmptyPassword
	}

	return nil
}

func ValidateEmail(email string) error {
	if strings.TrimSpace(email) == "" {
		return ErrEmptyEmail
	}

	if !strings.Contains(email, "@") {
		return ErrInvalidEmail
	}

	return nil
}

func ValidatePassword(password string) error {
	if strings.TrimSpace(password) == "" {
		return ErrEmptyPassword
	}

	if len(password) < 6 {
		return ErrPasswordTooShort
	}

	return nil
}
