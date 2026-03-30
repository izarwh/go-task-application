package apperror

import (
	"errors"
	"fmt"
	"net/http"
)

type AppError struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	StatusCode int    `json:"-"`
	Internal   error  `json:"-"`
}

func (e *AppError) Error() string {
	if e.Internal != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Internal)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func (e *AppError) Unwrap() error {
	return e.Internal
}

func New(statusCode int, code, message string, internal error) *AppError {
	return &AppError{
		StatusCode: statusCode,
		Code:       code,
		Message:    message,
		Internal:   internal,
	}
}
func Extract(err error) *AppError {
	var appErr *AppError

	if errors.As(err, &appErr) {
		return appErr
	}

	return &AppError{
		Code:       ErrInternal,
		Message:    "An unexpected error occurred",
		StatusCode: http.StatusInternalServerError,
		Internal:   err,
	}
}
