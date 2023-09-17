package mocks

import (
	"errors"

	"github.com/raphael-foliveira/chi-gorm/pkg/models"
)

type OrdersStore struct {
	Store       []models.Order
	ShouldError bool
}

func (cr *OrdersStore) List() ([]models.Order, error) {
	if cr.ShouldError {
		return nil, errors.New("")
	}
	return cr.Store, nil
}

func (cr *OrdersStore) Get(id int64) (*models.Order, error) {
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

func (cr *OrdersStore) Create(client *models.Order) error {
	if cr.ShouldError {
		return errors.New("")
	}
	cr.Store = append(cr.Store, *client)
	return nil
}

func (cr *OrdersStore) Update(client *models.Order) error {
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

func (cr *OrdersStore) Delete(client *models.Order) error {
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
