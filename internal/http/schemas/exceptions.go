package schemas

import (
	"errors"
	"fmt"
	"strings"
)

type ValidationError struct {
	Errors map[string][]string `json:"errors"`
}

func NewValidationError(err error) error {
	if err == nil {
		return nil
	}
	validationError := &ValidationError{
		Errors: make(map[string][]string),
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

func addFieldError(err error, key, message string) error {
	return errors.Join(err, fmt.Errorf("%s: %s", key, message))
}
