package mocks

import (
	"errors"

	"github.com/raphael-foliveira/chi-gorm/internal/entities"
)

type store[T entities.Entity] struct {
	Store       []T
	ShouldError bool
}

func newStore[T entities.Entity]() store[T] {
	return store[T]{Store: []T{}}
}

func (cr *store[T]) List() ([]T, error) {
	if cr.ShouldError {
		return nil, errors.New("")
	}
	return cr.Store, nil
}

func (cr *store[T]) Get(id int64) (*T, error) {
	if cr.ShouldError {
		return nil, errors.New("")
	}

	for _, entity := range cr.Store {
		if entity.GetId() == id {
			return &entity, nil
		}
	}
	return nil, errors.New("not found")
}

func (cr *store[T]) Create(client *T) error {
	if cr.ShouldError {
		return errors.New("")
	}
	cr.Store = append(cr.Store, *client)
	return nil
}

func (cr *store[T]) Update(client *T) error {
	if cr.ShouldError {
		return errors.New("")
	}
	for i, c := range cr.Store {
		if c.GetId() == (*client).GetId() {
			cr.Store[i] = (*client)
			return nil
		}
	}
	return errors.New("not found")
}

func (cr *store[T]) Delete(client *T) error {
	if cr.ShouldError {
		return errors.New("")
	}
	for i, c := range cr.Store {
		if c.GetId() == (*client).GetId() {
			cr.Store = append(cr.Store[:i], cr.Store[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}
