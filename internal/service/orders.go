package service

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
)

type OrdersService struct {
	repository OrdersRepository
}

func NewOrdersService(repository OrdersRepository) *OrdersService {
	return &OrdersService{repository}
}

func (o *OrdersService) Create(schema *schemas.CreateOrder) (*entities.Order, error) {
	newOrder := schema.ToModel()
	err := o.repository.Create(newOrder)
	return newOrder, err
}

func (o *OrdersService) Update(id uint, schema *schemas.UpdateOrder) (*entities.Order, error) {
	entity, err := o.Get(id)
	if err != nil {
		return nil, err
	}
	entity.Quantity = schema.Quantity
	err = o.repository.Update(entity)
	return entity, err
}

func (o *OrdersService) Delete(id uint) error {
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

func (o *OrdersService) List() ([]entities.Order, error) {
	return o.repository.List()
}

func (o *OrdersService) Get(id uint) (*entities.Order, error) {
	order, err := o.repository.Get(id)
	if err != nil || order == nil {
		return nil, errOrderNotFound
	}
	return order, nil
}
