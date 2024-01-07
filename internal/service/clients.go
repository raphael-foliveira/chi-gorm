package service

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/exceptions"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
)

type Clients interface {
	Create(*schemas.CreateClient) (*entities.Client, error)
	Update(uint, *schemas.UpdateClient) (*entities.Client, error)
	Delete(uint) error
	List() ([]entities.Client, error)
	Get(uint) (*entities.Client, error)
	GetProducts(uint) ([]entities.Order, error)
}

type clients struct {
	repository       repository.Clients
	ordersRepository repository.Orders
}

func NewClients(repository repository.Clients, ordersRepository repository.Orders) Clients {
	return &clients{repository, ordersRepository}
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
		return nil, exceptions.NotFound("client not found")
	}
	return client, nil
}

func (c *clients) GetProducts(clientId uint) ([]entities.Order, error) {
	client, err := c.Get(clientId)
	if err != nil {
		return nil, err
	}
	return c.ordersRepository.FindManyByClientId(client.ID)
}
