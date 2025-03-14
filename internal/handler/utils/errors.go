package utils

import "errors"

var (
	ErrInvalidRequestPayload = errors.New("Invalid request payload")
	ErrInvalidID             = errors.New("Invalid ID")
	ErrUserIDNotFound        = errors.New("User ID not found in context")

	ErrInvalidTodoID    = errors.New("Invalid Todo ID")
	ErrTodoNotFound     = errors.New("Todo not found")
	ErrUnauthorized     = errors.New("Unauthorized to access this resource")
	ErrCreateTodoFailed = errors.New("Failed to create todo")
	ErrGetTodosFailed   = errors.New("Failed to get todos")
	ErrUpdateTodoFailed = errors.New("Failed to update todo")
	ErrDeleteTodoFailed = errors.New("Failed to delete todo")

	ErrRegisterUserFailed = errors.New("Failed to register user")

	ErrAuthenticationFailed  = errors.New("Authentication failed")
	ErrTokenGenerationFailed = errors.New("Failed to generate token")

	ErrEmptyTitle       = errors.New("Title cannot be empty")
	ErrEmptyEmail       = errors.New("Email cannot be empty")
	ErrInvalidEmail     = errors.New("Invalid email format")
	ErrEmptyPassword    = errors.New("Password cannot be empty")
	ErrPasswordTooShort = errors.New("Password must be at least 6 characters")
	ErrEmptyName        = errors.New("Name cannot be empty")
)
