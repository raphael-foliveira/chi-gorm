package service

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
)

type orders struct{}

func NewOrders() *orders {
	return &orders{}
}

func (o *orders) Create(schema *schemas.CreateOrder) (*entities.Order, error) {
	newOrder := schema.ToModel()
	err := repository.Orders.Create(newOrder)
	return newOrder, err
}

func (o *orders) Update(id uint, schema *schemas.UpdateOrder) (*entities.Order, error) {
	entity, err := o.Get(id)
	if err != nil {
		return nil, err
	}
	if schema.Quantity != 0 {
		entity.Quantity = schema.Quantity
	}
	err = repository.Orders.Update(entity)
	return entity, err
}

func (o *orders) Delete(id uint) error {
	client, err := o.Get(id)
	if err != nil {
		return err
	}
	err = repository.Orders.Delete(client)
	if err != nil {
		return err
	}
	return nil
}

func (o *orders) List() ([]entities.Order, error) {
	return repository.Orders.List()
}

func (o *orders) Get(id uint) (*entities.Order, error) {
	order, err := repository.Orders.Get(id)
	if err != nil || order == nil {
		return nil, errOrderNotFound
	}
	return order, nil
}
