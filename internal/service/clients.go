package service

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
)

type Clients struct {
	repository       ClientsRepository
	ordersRepository OrdersRepository
}

func NewClients(repository ClientsRepository, ordersRepository OrdersRepository) *Clients {
	return &Clients{repository, ordersRepository}
}

func (c *Clients) Create(schema *schemas.CreateClient) (*entities.Client, error) {
	newClient := schema.ToModel()
	err := c.repository.Create(newClient)
	return newClient, err
}

func (c *Clients) Update(id uint, schema *schemas.UpdateClient) (*entities.Client, error) {
	entity, err := c.Get(id)
	if err != nil {
		return nil, err
	}
	entity.Name = schema.Name
	entity.Email = schema.Email
	err = c.repository.Update(entity)
	return entity, err
}

func (c *Clients) Delete(id uint) error {
	client, err := c.Get(id)
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

func (c *Clients) Get(id uint) (*entities.Client, error) {
	client, err := c.repository.Get(id)
	if err != nil {
		return nil, errClientNotFound
	}
	return client, nil
}

func (c *Clients) GetOrders(clientId uint) ([]entities.Order, error) {
	client, err := c.Get(clientId)
	if err != nil {
		return nil, err
	}
	return c.ordersRepository.FindManyByClientId(client.ID)
}
