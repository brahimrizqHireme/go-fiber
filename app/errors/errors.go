package errors

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type AppError struct {
	Message    string
	StatusCode int
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(statusCode int, message string) *AppError {
	return &AppError{
		Message:    message,
		StatusCode: statusCode,
	}
}

var (
	ErrUserAlreadyExists  = NewAppError(fiber.StatusUnprocessableEntity, "User already exists")
	ErrInvalidCredentials = NewAppError(fiber.StatusUnprocessableEntity, "Invalid credentials")
	ErrNotFoundError      = NewAppError(fiber.StatusNotFound, "Entity not found")
)

func WrapAppError(statusCode int, err error, message string) *AppError {
	return &AppError{
		Message:    fmt.Sprintf("%s: %v", message, err),
		StatusCode: statusCode,
	}
}
