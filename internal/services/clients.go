package services

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/exceptions"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
)

var clientNotFoundErr = &exceptions.NotFoundError{Entity: "client"}

type Clients struct {
	repository repository.Clients
}

func NewClients(clientsRepo repository.Clients) *Clients {
	return &Clients{clientsRepo}
}

func (c *Clients) Create(schema *schemas.CreateClient) (*entities.Client, error) {
	validationErr := schema.Validate()
	if validationErr != nil {
		return nil, validationErr
	}
	newClient := schema.ToModel()
	err := c.repository.Create(newClient)
	return newClient, err
}

func (c *Clients) Update(id int64, schema *schemas.UpdateClient) (*entities.Client, error) {
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

func (c *Clients) Delete(id int64) error {
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

func (c *Clients) List() ([]entities.Client, error) {
	return c.repository.List()
}

func (c *Clients) Get(id int64) (*entities.Client, error) {
	client, err := c.repository.Get(id)
	if err != nil {
		return nil, clientNotFoundErr
	}
	return client, nil
}
