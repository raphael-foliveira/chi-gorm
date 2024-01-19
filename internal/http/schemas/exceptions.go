package schemas

import (
	"fmt"
	"strings"
)

type ValidationErrors map[string]string

func (vem ValidationErrors) Error() string {
	message := ""
	for key, value := range vem {
		message += fmt.Sprintf("%s %s\n", key, value)
	}
	return message
}

func NewValidationErrors(err error) ValidationErrors {
	ve := ValidationErrors{}
	errors := strings.Split(err.Error(), "\n")
	for _, e := range errors {
		split := strings.Split(e, ": ")
		ve[split[0]] = split[1]
	}
	return ve
}
