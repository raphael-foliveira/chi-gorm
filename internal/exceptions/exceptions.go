package exceptions

import "strings"

type NotFoundError struct {
	Entity string
}

func (e *NotFoundError) Error() string {
	return e.Entity + " not found"
}

type ApiError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (ae *ApiError) Error() string {
	return ae.Message
}

type MultipleApiError struct {
	Errors []string `json:"errors"`
	Status int      `json:"status"`
}

func (ae *MultipleApiError) Error() string {
	return strings.Join(ae.Errors, ", ")
}

type ValidationError struct {
	Message string `json:"message"`
}

func (ve *ValidationError) Error() string {
	return ve.Message
}
