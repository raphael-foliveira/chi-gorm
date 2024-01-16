package schemas

import "strings"

type ValidationError struct {
	Errors []string `json:"errors"`
}

func NewValidationError(status ...int) *ValidationError {
	return &ValidationError{
		Errors: []string{},
	}
}

func (ae *ValidationError) Error() string {
	if ae.Errors == nil {
		return ""
	}
	return strings.Join(ae.Errors, ", ")
}

func (ae *ValidationError) Add(err error) {
	ae.Errors = append(ae.Errors, err.Error())
}

func (ae *ValidationError) ReturnIfError() error {
	if len(ae.Errors) > 0 {
		return ae
	}
	return nil
}
