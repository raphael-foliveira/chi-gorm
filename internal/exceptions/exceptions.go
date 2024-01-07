package exceptions

import "strings"

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

type ValidationError struct {
	ApiError
}

func NewValidationError(message string) *ValidationError {
	return &ValidationError{
		ApiError: ApiError{
			Message: message,
			Status:  422,
		},
	}
}

type NotFoundError struct {
	ApiError
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{
		ApiError: ApiError{
			Message: message,
			Status:  404,
		},
	}
}
