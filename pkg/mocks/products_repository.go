package mocks

import (
	"errors"

	"github.com/raphael-foliveira/chi-gorm/pkg/models"
)

type ProductsRepository struct {
	Store       []models.Product
	ShouldError bool
}

func (cr *ProductsRepository) List() ([]models.Product, error) {
	if cr.ShouldError {
		return nil, errors.New("")
	}
	return cr.Store, nil
}

func (cr *ProductsRepository) Get(id int64) (*models.Product, error) {
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

func (cr *ProductsRepository) Create(client *models.Product) error {
	if cr.ShouldError {
		return errors.New("")
	}
	cr.Store = append(cr.Store, *client)
	return nil
}

func (cr *ProductsRepository) Update(client *models.Product) error {
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

func (cr *ProductsRepository) Delete(client *models.Product) error {
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
