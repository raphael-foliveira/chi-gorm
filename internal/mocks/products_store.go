package mocks

import (
	"errors"

	"github.com/raphael-foliveira/chi-gorm/internal/entities"
)

type ProductsStore struct {
	Store       []entities.Product
	ShouldError bool
}

func (cr *ProductsStore) List() ([]entities.Product, error) {
	if cr.ShouldError {
		return nil, errors.New("")
	}
	return cr.Store, nil
}

func (cr *ProductsStore) Get(id int64) (*entities.Product, error) {
	if cr.ShouldError {
		return nil, errors.New("")
	}
	for _, client := range cr.Store {
		if client.ID == id {
			return &client, nil
		}
	}
	return nil, errors.New("not found")
}

func (cr *ProductsStore) Create(client *entities.Product) error {
	if cr.ShouldError {
		return errors.New("")
	}
	cr.Store = append(cr.Store, *client)
	return nil
}

func (cr *ProductsStore) Update(client *entities.Product) error {
	if cr.ShouldError {
		return errors.New("")
	}
	for i, c := range cr.Store {
		if c.ID == client.ID {
			cr.Store[i] = *client
			return nil
		}
	}
	return errors.New("not found")
}

func (cr *ProductsStore) Delete(client *entities.Product) error {
	if cr.ShouldError {
		return errors.New("")
	}
	for i, c := range cr.Store {
		if c.ID == client.ID {
			cr.Store = append(cr.Store[:i], cr.Store[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}
