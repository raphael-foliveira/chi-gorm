package service

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
)

type Clients struct{}

func NewClients() *Clients {
	return &Clients{}
}

func (c *Clients) Create(schema *schemas.CreateClient) (*entities.Client, error) {
	newClient := schema.ToModel()
	err := clientsRepository.Create(newClient)
	return newClient, err
}

func (c *Clients) Update(id uint, schema *schemas.UpdateClient) (*entities.Client, error) {
	entity, err := c.Get(id)
	if err != nil {
		return nil, err
	}
	c.updatePopulatedFields(entity, schema)
	err = clientsRepository.Update(entity)
	return entity, err
}

func (c *Clients) updatePopulatedFields(client *entities.Client, schema *schemas.UpdateClient) {
	if schema.Name != "" {
		client.Name = schema.Name
	}
	if schema.Email != "" {
		client.Email = schema.Email
	}
}

func (c *Clients) Delete(id uint) error {
	client, err := c.Get(id)
	if err != nil {
		return err
	}
	err = clientsRepository.Delete(client)
	if err != nil {
		return err
	}
	return nil
}

func (c *Clients) List() ([]entities.Client, error) {
	return clientsRepository.List()
}

func (c *Clients) Get(id uint) (*entities.Client, error) {
	client, err := clientsRepository.Get(id)
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
	return ordersRepository.FindManyByClientId(client.ID)
}
