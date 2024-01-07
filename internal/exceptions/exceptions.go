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

func BadRequest(message string) *ApiError {
	return NewApiError(message, http.StatusBadRequest)
}

func UnprocessableEntity(message string) *ApiError {
	return NewApiError(message, http.StatusUnprocessableEntity)
}

func InternalServerError(message string) *ApiError {
	return NewApiError(message, http.StatusInternalServerError)
}

func NotFound(message string) *ApiError {
	return NewApiError(message, http.StatusNotFound)
}

type MultipleApiError struct {
	Errors []string `json:"errors"`
	Status int      `json:"status"`
}

func (ae *MultipleApiError) Error() string {
	return strings.Join(ae.Errors, ", ")
}
