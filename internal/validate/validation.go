package validate

import (
	"fmt"
	"strings"
)

type ValidationError map[string][]string

func Rules(err ...error) error {
	validationError := ValidationError{}
	for _, m := range err {
		if m != nil {
			splitMessage := strings.Split(m.Error(), ":")
			key := strings.TrimSpace(splitMessage[0])
			message := strings.TrimSpace(splitMessage[1])
			validationError[key] = append(validationError[key], message)
		}
	}
	if len(validationError) == 0 {
		return nil
	}
	return &validationError
}

func (ve *ValidationError) Error() string {
	return fmt.Sprintf("%v", *ve)
}
