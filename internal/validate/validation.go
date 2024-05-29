package validate

import (
	"encoding/json"
	"strings"
)

type ValidationError map[string][]string

func (ve *ValidationError) Error() string {
	errBytes, _ := json.Marshal(ve)
	return string(errBytes)
}

func Rules(err ...error) error {
	validationError := ValidationError{}
	for _, m := range err {
		if m != nil {
			splitMessage := strings.SplitN(m.Error(), ":", 2)
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
