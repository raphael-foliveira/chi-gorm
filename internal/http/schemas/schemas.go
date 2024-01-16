package schemas

import "errors"

type CreateSchema interface {
	ToModel() interface{}
}

type ValidateableSchema interface {
	Validate() error
}

var ErrValidation = errors.New("validation error")
