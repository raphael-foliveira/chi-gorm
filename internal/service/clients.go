package service

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/exceptions"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
)

var clientNotFoundErr = &exceptions.NotFoundError{Entity: "client"}

type Clients interface {
	Create(schema *schemas.CreateClient) (*entities.Client, error)
	Update(id uint, schema *schemas.UpdateClient) (*entities.Client, error)
	Delete(id uint) error
	List() ([]entities.Client, error)
	Get(id uint) (*entities.Client, error)
}

type clients struct {
	repository repository.Clients
}

func NewClients(repository repository.Clients) Clients {
	return &clients{repository}
}

func (c *clients) Create(schema *schemas.CreateClient) (*entities.Client, error) {
	validationErr := schema.Validate()
	if validationErr != nil {
		return nil, validationErr
	}
	newClient := schema.ToModel()
	err := c.repository.Create(newClient)
	return newClient, err
}

func (c *clients) Update(id uint, schema *schemas.UpdateClient) (*entities.Client, error) {
	validationErr := schema.Validate()
	if validationErr != nil {
		return nil, validationErr
	}
	entity, err := c.repository.Get(id)
	if err != nil {
		return nil, err
	}
	entity.Name = schema.Name
	entity.Email = schema.Email
	err = c.repository.Update(entity)
	return entity, err
}

func (c *clients) Delete(id uint) error {
	client, err := c.repository.Get(id)
	if err != nil {
		return err
	}
	err = c.repository.Delete(client)
	if err != nil {
		return err
	}
	return nil
}

func (c *clients) List() ([]entities.Client, error) {
	return c.repository.List()
}

func (c *clients) Get(id uint) (*entities.Client, error) {
	client, err := c.repository.Get(id)
	if err != nil {
		return nil, clientNotFoundErr
	}
	return client, nil
}
