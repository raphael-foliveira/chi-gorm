package service

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
)

type clients struct{}

func NewClients() *clients {
	return &clients{}
}

func (c *clients) Create(schema *schemas.CreateClient) (*entities.Client, error) {
	newClient := schema.ToModel()
	err := repository.Clients.Create(newClient)
	return newClient, err
}

func (c *clients) Update(id uint, schema *schemas.UpdateClient) (*entities.Client, error) {
	entity, err := c.Get(id)
	if err != nil {
		return nil, err
	}
	c.updatePopulatedFields(entity, schema)
	err = repository.Clients.Update(entity)
	return entity, err
}

func (c *clients) updatePopulatedFields(client *entities.Client, schema *schemas.UpdateClient) {
	if schema.Name != "" {
		client.Name = schema.Name
	}
	if schema.Email != "" {
		client.Email = schema.Email
	}
}

func (c *clients) Delete(id uint) error {
	client, err := c.Get(id)
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
		return nil, errClientNotFound
	}
	return client, nil
}

func (c *clients) GetOrders(clientId uint) ([]entities.Order, error) {
	client, err := c.Get(clientId)
	if err != nil {
		return nil, err
	}
	return repository.Orders.FindManyByClientId(client.ID)
}
