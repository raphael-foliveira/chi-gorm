package exceptions

import (
	"net/http"
	"strings"
)

type ApiError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (ae *ApiError) Error() string {
	return ae.Message
}

func NewApiError(message string, status int) *ApiError {
	return &ApiError{
		Message: message,
		Status:  status,
	}
}

type MultipleApiError struct {
	Errors []string `json:"errors"`
	Status int      `json:"status"`
}

func (ae *MultipleApiError) Error() string {
	return strings.Join(ae.Errors, ", ")
}

func NewValidationError(message string) *ApiError {
	return &ApiError{
		Message: message,
		Status:  http.StatusBadRequest,
	}
}

func NewNotFoundError(message string) *ApiError {
	return &ApiError{
		Message: message,
		Status:  http.StatusNotFound,
	}
}
