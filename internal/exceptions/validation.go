package exceptions

import (
	"fmt"
	"net/http"
	"strings"
)

type ValidationError struct {
	Errors map[string][]string `json:"errors"`
	Status int                 `json:"status"`
}

func NewValidationError(err error, status ...int) error {
	if err == nil {
		return nil
	}
	validationError := &ValidationError{
		Errors: make(map[string][]string),
		Status: http.StatusUnprocessableEntity,
	}
	if len(status) > 0 {
		validationError.Status = status[0]
	}
	errorMessages := strings.Split(err.Error(), "\n")
	for _, m := range errorMessages {
		if m != "" {
			splitMessage := strings.Split(m, ":")
			key := strings.TrimSpace(splitMessage[0])
			message := strings.TrimSpace(splitMessage[1])
			validationError.Errors[key] = append(validationError.Errors[key], message)
		}
	}
	return validationError
}

func (ve *ValidationError) Error() string {
	return fmt.Sprintf("%v", ve.Errors)
}
