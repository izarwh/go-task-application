package apperror

import (
	"errors"
	"net/http"

	"gorm.io/gorm"
)

// this is custome err that i usually use for entire my project
func FromDB(err error, entity string) *AppError {
	if err == nil {
		return nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return NotFound(entity+" not found", err)
	}

	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return New(http.StatusConflict, ErrConflict, entity+" already exists", err)
	}

	if errors.Is(err, gorm.ErrForeignKeyViolated) {
		return New(http.StatusForbidden, "DEPENDENCY_ERROR", "this "+entity+" is linked to other records", err)
	}

	if errors.Is(err, gorm.ErrCheckConstraintViolated) {
		return New(http.StatusBadRequest, "CONSTRAINT_VIOLATION", "invalid data format", err)
	}

	if errors.Is(err, gorm.ErrInvalidTransaction) {
		return Internal(err)
	}

	return Internal(err)
}

var (
	ErrNotFound     = "NOT_FOUND"
	ErrBadRequest   = "BAD_REQUEST"
	ErrInternal     = "INTERNAL_SERVER_ERROR"
	ErrUnauthorized = "UNAUTHORIZED"
	ErrConflict     = "CONFLICT"
)

func NotFound(message string, internal error) *AppError {
	return New(http.StatusNotFound, ErrNotFound, message, internal)
}

func Internal(internal error) *AppError {
	return New(http.StatusInternalServerError, ErrInternal, "An unexpected error occurred", internal)
}

func BadRequest(message string, internal error) *AppError {
	return New(http.StatusBadRequest, ErrBadRequest, message, internal)
}
