package mocks

import (
	"errors"

	"github.com/raphael-foliveira/chi-gorm/internal/entities"
)

type store[T entities.Entity] struct {
	Store []T
	Error error
}

func (cr *store[T]) List(conds ...interface{}) ([]T, error) {
	return cr.Store, cr.Error
}

func (cr *store[T]) Get(id uint) (*T, error) {
	for _, entity := range cr.Store {
		if entity.GetId() == id {
			return &entity, cr.Error
		}
	}
	return nil, errors.New("not found")
}

func (cr *store[T]) Create(client *T) error {
	cr.Store = append(cr.Store, *client)
	return cr.Error
}

func (cr *store[T]) Update(client *T) error {
	for i, c := range cr.Store {
		if c.GetId() == (*client).GetId() {
			cr.Store[i] = (*client)
			return cr.Error
		}
	}
	return errors.New("not found")
}

func (cr *store[T]) Delete(client *T) error {
	for i, c := range cr.Store {
		if c.GetId() == (*client).GetId() {
			cr.Store = append(cr.Store[:i], cr.Store[i+1:]...)
			return cr.Error
		}
	}
	return errors.New("not found")
}
