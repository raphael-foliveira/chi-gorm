package service

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
)

type Orders struct{}

func NewOrders() *Orders {
	return &Orders{}
}

func (o *Orders) Create(schema *schemas.CreateOrder) (*entities.Order, error) {
	newOrder := schema.ToModel()
	err := ordersRepository.Create(newOrder)
	return newOrder, err
}

func (o *Orders) Update(id uint, schema *schemas.UpdateOrder) (*entities.Order, error) {
	entity, err := o.Get(id)
	if err != nil {
		return nil, err
	}
	if schema.Quantity != 0 {
		entity.Quantity = schema.Quantity
	}
	err = ordersRepository.Update(entity)
	return entity, err
}

func (o *Orders) Delete(id uint) error {
	client, err := o.Get(id)
	if err != nil {
		return err
	}
	err = ordersRepository.Delete(client)
	if err != nil {
		return err
	}
	return nil
}

func (o *Orders) List() ([]entities.Order, error) {
	return ordersRepository.List()
}

func (o *Orders) Get(id uint) (*entities.Order, error) {
	order, err := ordersRepository.Get(id)
	if err != nil || order == nil {
		return nil, errOrderNotFound
	}
	return order, nil
}
