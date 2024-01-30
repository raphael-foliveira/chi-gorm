package validate

import (
	"fmt"
	"strings"
)

type ValidationError struct {
	Errors map[string][]string `json:"errors"`
}

func Rules(err ...error) error {
	validationError := &ValidationError{
		Errors: make(map[string][]string),
	}
	for _, m := range err {
		if m != nil {
			splitMessage := strings.Split(m.Error(), ":")
			key := strings.TrimSpace(splitMessage[0])
			message := strings.TrimSpace(splitMessage[1])
			validationError.Errors[key] = append(validationError.Errors[key], message)
		}
	}
	if len(validationError.Errors) == 0 {
		return nil
	}
	return validationError
}

func (ve *ValidationError) Error() string {
	return fmt.Sprintf("%v", ve.Errors)
}
