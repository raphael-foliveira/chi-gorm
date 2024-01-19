package service

import (
	"fmt"

	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
)

type Orders struct {
	repository repository.Orders
}

func NewOrders(repository repository.Orders) *Orders {
	return &Orders{repository}
}

func (c *Orders) Create(schema *schemas.CreateOrder) (*entities.Order, error) {
	newOrder := schema.ToModel()
	err := c.repository.Create(newOrder)
	return newOrder, err
}

func (c *Orders) Update(id uint, schema *schemas.UpdateOrder) (*entities.Order, error) {
	entity, err := c.Get(id)
	if err != nil {
		return nil, err
	}
	entity.Quantity = schema.Quantity
	err = c.repository.Update(entity)
	return entity, err
}

func (c *Orders) Delete(id uint) error {
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

func (c *Orders) List() ([]entities.Order, error) {
	return c.repository.List()
}

func (c *Orders) Get(id uint) (*entities.Order, error) {
	order, err := c.repository.Get(id)
	if err != nil || order == nil {
		return nil, errOrderNotFound
	}
	return order, nil
}

var errOrderNotFound = fmt.Errorf("order %w", ErrNotFound)
