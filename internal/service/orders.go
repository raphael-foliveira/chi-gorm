package service

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/interfaces"
)

type Orders struct {
	repository interfaces.OrdersRepository
}

func NewOrders(repository interfaces.OrdersRepository) *Orders {
	return &Orders{repository}
}

func (o *Orders) Create(schema *schemas.CreateOrder) (*entities.Order, error) {
	newOrder := schema.ToModel()
	err := o.repository.Create(newOrder)
	return newOrder, err
}

func (o *Orders) Update(id uint, schema *schemas.UpdateOrder) (*entities.Order, error) {
	entity, err := o.Get(id)
	if err != nil {
		return nil, err
	}
	entity.Quantity = schema.Quantity
	err = o.repository.Update(entity)
	return entity, err
}

func (o *Orders) Delete(id uint) error {
	client, err := o.Get(id)
	if err != nil {
		return err
	}
	err = o.repository.Delete(client)
	if err != nil {
		return err
	}
	return nil
}

func (o *Orders) List() ([]entities.Order, error) {
	return o.repository.List()
}

func (o *Orders) Get(id uint) (*entities.Order, error) {
	order, err := o.repository.Get(id)
	if err != nil || order == nil {
		return nil, errOrderNotFound
	}
	return order, nil
}
