package api

import (
	"encoding/json"
	"net/http"
)

type ApiError struct {
	Err    any `json:"error"`
	Status int `json:"status"`
}

func NewApiError(status int, errors any) *ApiError {
	return &ApiError{
		Err:    errors,
		Status: status,
	}
}

func (ae *ApiError) Error() string {
	data, _ := json.Marshal(ae.Err)
	return string(data)
}

func BadRequest(message string) *ApiError {
	return NewApiError(http.StatusBadRequest, message)
}

func UnprocessableEntity(message string) *ApiError {
	return NewApiError(http.StatusUnprocessableEntity, message)
}

func InternalServerError(message string) *ApiError {
	return NewApiError(http.StatusInternalServerError, message)
}

func NotFound(message string) *ApiError {
	return NewApiError(http.StatusNotFound, message)
}

func Unauthorized(message string) *ApiError {
	return NewApiError(http.StatusUnauthorized, message)
}

func Forbidden(message string) *ApiError {
	return NewApiError(http.StatusForbidden, message)
}
