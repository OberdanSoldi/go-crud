package util

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AppError implements the error interface
type AppError struct {
	StatusCode int
	Message    string
}

// Error implements the error interface
func (e *AppError) Error() string {
	return e.Message
}

func NewBadRequestError(message string) error {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		Message:    message,
	}
}

func NewNotFoundError(message string) error {
	return &AppError{
		StatusCode: http.StatusNotFound,
		Message:    message,
	}
}

func NewInternalServerError(message string) error {
	return &AppError{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
	}
}

func RespondWithError(c *gin.Context, err error) {
	var appErr *AppError
	ok := errors.As(err, &appErr)
	if !ok {
		appErr = &AppError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	c.JSON(appErr.StatusCode, gin.H{"error": appErr.Message})
}
