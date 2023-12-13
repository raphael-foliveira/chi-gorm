package services

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/exceptions"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
)

var orderNotFoundErr = &exceptions.NotFoundError{Entity: "order"}

type Orders struct {
	repository repository.Orders
}

func NewOrders(productsRepo repository.Orders) *Orders {
	return &Orders{productsRepo}
}

func (c *Orders) Create(schema *schemas.CreateOrder) (*entities.Order, error) {
	newOrder := schema.ToModel()
	err := c.repository.Create(newOrder)
	return newOrder, err
}

func (c *Orders) Update(id int64, schema *schemas.UpdateOrder) (*entities.Order, error) {
	entity, err := c.repository.Get(id)
	if err != nil {
		return nil, err
	}
	entity.Quantity = schema.Quantity
	err = c.repository.Update(entity)
	return entity, err
}

func (c *Orders) Delete(id int64) error {
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

func (c *Orders) List() ([]entities.Order, error) {
	return c.repository.List()
}

func (c *Orders) Get(id int64) (*entities.Order, error) {
	client, err := c.repository.Get(id)
	if err != nil {
		return nil, orderNotFoundErr
	}
	return client, nil
}
