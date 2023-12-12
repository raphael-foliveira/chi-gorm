package mocks

import (
	"errors"

	"github.com/raphael-foliveira/chi-gorm/internal/entities"
)

type OrdersStore struct {
	Store       []entities.Order
	ShouldError bool
}

func (cr *OrdersStore) List() ([]entities.Order, error) {
	if cr.ShouldError {
		return nil, errors.New("")
	}
	return cr.Store, nil
}

func (os *OrdersStore) Get(id int64) (*entities.Order, error) {
	if os.ShouldError {
		return nil, errors.New("")
	}
	for _, order := range os.Store {
		if order.ID == id {
			return &order, nil
		}
	}
	return nil, errors.New("not found")
}

func (cr *OrdersStore) Create(client *entities.Order) error {
	if cr.ShouldError {
		return errors.New("")
	}
	cr.Store = append(cr.Store, *client)
	return nil
}

func (cr *OrdersStore) Update(client *entities.Order) error {
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

func (cr *OrdersStore) Delete(client *entities.Order) error {
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
