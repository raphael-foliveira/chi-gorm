package service

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/exceptions"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
)

type Orders interface {
	Create(schema *schemas.CreateOrder) (*entities.Order, error)
	Update(id uint, schema *schemas.UpdateOrder) (*entities.Order, error)
	Delete(id uint) error
	List() ([]entities.Order, error)
	Get(id uint) (*entities.Order, error)
	FindManyByClientId(clientId uint) ([]entities.Order, error)
}

type orders struct {
	repository repository.Orders
}

func NewOrders(repository repository.Orders) Orders {
	return &orders{repository}
}

func (c *orders) Create(schema *schemas.CreateOrder) (*entities.Order, error) {
	validationErr := schema.Validate()
	if validationErr != nil {
		return nil, validationErr
	}
	newOrder := schema.ToModel()
	err := c.repository.Create(newOrder)
	return newOrder, err
}

func (c *orders) Update(id uint, schema *schemas.UpdateOrder) (*entities.Order, error) {
	validationErr := schema.Validate()
	if validationErr != nil {
		return nil, validationErr
	}
	entity, err := c.repository.Get(id)
	if err != nil {
		return nil, err
	}
	entity.Quantity = schema.Quantity
	err = c.repository.Update(entity)
	return entity, err
}

func (c *orders) Delete(id uint) error {
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

func (c *orders) List() ([]entities.Order, error) {
	return c.repository.List()
}

func (c *orders) Get(id uint) (*entities.Order, error) {
	order, err := c.repository.Get(id)
	if err != nil || order == nil {
		return nil, exceptions.NewNotFoundError("order not found")
	}
	return order, nil
}

func (c *orders) FindManyByClientId(clientId uint) ([]entities.Order, error) {
	return c.repository.FindManyByClientId(clientId)
}
