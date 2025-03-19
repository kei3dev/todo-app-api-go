package errors

import "errors"

var (
	// Common errors
	ErrInvalidRequestPayload = errors.New("Invalid request payload")
	ErrInvalidID             = errors.New("Invalid ID")
	ErrUserIDNotFound        = errors.New("User ID not found in context")

	// User related errors
	ErrInvalidEmailFormat = errors.New("invalid email format")
	ErrPasswordTooShort   = errors.New("password must be at least 8 characters")
	ErrNameTooShort       = errors.New("name must be at least 2 characters")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrRegisterUserFailed = errors.New("Failed to register user")

	// Todo related errors
	ErrInvalidTodoID    = errors.New("Invalid Todo ID")
	ErrTodoNotFound     = errors.New("Todo not found")
	ErrUnauthorized     = errors.New("Unauthorized to access this resource")
	ErrCreateTodoFailed = errors.New("Failed to create todo")
	ErrGetTodosFailed   = errors.New("Failed to get todos")
	ErrUpdateTodoFailed = errors.New("Failed to update todo")
	ErrDeleteTodoFailed = errors.New("Failed to delete todo")

	// Authentication errors
	ErrAuthenticationFailed  = errors.New("Authentication failed")
	ErrTokenGenerationFailed = errors.New("Failed to generate token")

	// Validation errors
	ErrEmptyTitle    = errors.New("Title cannot be empty")
	ErrEmptyEmail    = errors.New("Email cannot be empty")
	ErrEmptyPassword = errors.New("Password cannot be empty")
	ErrEmptyName     = errors.New("Name cannot be empty")
)
