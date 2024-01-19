package controller

import "net/http"

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

func ErrBadRequest(message string) *ApiError {
	return NewApiError(message, http.StatusBadRequest)
}

func ErrUnprocessableEntity(message string) *ApiError {
	return NewApiError(message, http.StatusUnprocessableEntity)
}

func InternalServerError(message string) *ApiError {
	return NewApiError(message, http.StatusInternalServerError)
}

func NotFound(message string) *ApiError {
	return NewApiError(message, http.StatusNotFound)
}
