package exceptions

import "net/http"

type ApiError struct {
	Message string `json:"error"`
	Status  int    `json:"status"`
}

func (ae *ApiError) Error() string {
	return ae.Message
}

func NewApiError(status int, message string) *ApiError {
	return &ApiError{
		Message: message,
		Status:  status,
	}
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
