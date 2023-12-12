package mocks

import (
	"errors"

	"github.com/raphael-foliveira/chi-gorm/internal/entities"
)

type ClientsStore struct {
	Store       []entities.Client
	ShouldError bool
}

func (cr *ClientsStore) List() ([]entities.Client, error) {
	if cr.ShouldError {
		return nil, errors.New("Error")
	}
	return cr.Store, nil
}

func (cr *ClientsStore) Get(id int64) (*entities.Client, error) {
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

func (cr *ClientsStore) Create(client *entities.Client) error {
	if cr.ShouldError {
		return errors.New("")
	}
	cr.Store = append(cr.Store, *client)
	return nil
}

func (cr *ClientsStore) Update(client *entities.Client) error {
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

func (cr *ClientsStore) Delete(client *entities.Client) error {
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
