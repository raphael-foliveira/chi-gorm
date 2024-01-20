package schemas

import (
	"errors"
	"fmt"
)

type CreateSchema interface {
	ToModel() interface{}
}

type ValidateableSchema interface {
	Validate() error
}

func addFieldError(err error, key, message string) error {
	return errors.Join(err, fmt.Errorf("%s: %s", key, message))
}
