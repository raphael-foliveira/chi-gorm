package services

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/exceptions"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
)

var orderNotFoundErr = &exceptions.NotFoundError{Entity: "order"}

var Orders = &orders{}

type orders struct{}

func (c *orders) Create(schema *schemas.CreateOrder) (*entities.Order, error) {
	validationErr := schema.Validate()
	if validationErr != nil {
		return nil, validationErr
	}
	newOrder := schema.ToModel()
	err := repository.Orders.Create(newOrder)
	return newOrder, err
}

func (c *orders) Update(id uint, schema *schemas.UpdateOrder) (*entities.Order, error) {
	validationErr := schema.Validate()
	if validationErr != nil {
		return nil, validationErr
	}
	entity, err := repository.Orders.Get(id)
	if err != nil {
		return nil, err
	}
	entity.Quantity = schema.Quantity
	err = repository.Orders.Update(entity)
	return entity, err
}

func (c *orders) Delete(id uint) error {
	client, err := repository.Orders.Get(id)
	if err != nil {
		return err
	}
	err = repository.Orders.Delete(client)
	if err != nil {
		return err
	}
	return nil
}

func (c *orders) List() ([]entities.Order, error) {
	return repository.Orders.List()
}

func (c *orders) Get(id uint) (*entities.Order, error) {
	order, err := repository.Orders.Get(id)
	if err != nil || order == nil {
		return nil, orderNotFoundErr
	}
	return order, nil
}
