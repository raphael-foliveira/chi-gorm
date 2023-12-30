package services

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/exceptions"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
)

var clientNotFoundErr = &exceptions.NotFoundError{Entity: "client"}
var Clients = &clients{}

type clients struct{}

func (c *clients) Create(schema *schemas.CreateClient) (*entities.Client, error) {
	validationErr := schema.Validate()
	if validationErr != nil {
		return nil, validationErr
	}
	newClient := schema.ToModel()
	err := repository.Clients.Create(newClient)
	return newClient, err
}

func (c *clients) Update(id uint, schema *schemas.UpdateClient) (*entities.Client, error) {
	validationErr := schema.Validate()
	if validationErr != nil {
		return nil, validationErr
	}
	entity, err := repository.Clients.Get(id)
	if err != nil {
		return nil, err
	}
	entity.Name = schema.Name
	entity.Email = schema.Email
	err = repository.Clients.Update(entity)
	return entity, err
}

func (c *clients) Delete(id uint) error {
	client, err := repository.Clients.Get(id)
	if err != nil {
		return err
	}
	err = repository.Clients.Delete(client)
	if err != nil {
		return err
	}
	return nil
}

func (c *clients) List() ([]entities.Client, error) {
	return repository.Clients.List()
}

func (c *clients) Get(id uint) (*entities.Client, error) {
	client, err := repository.Clients.Get(id)
	if err != nil {
		return nil, clientNotFoundErr
	}
	return client, nil
}
