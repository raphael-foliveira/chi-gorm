package mocks

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
)

var OrdersStore = &ordersStore{store[entities.Order]{}}

type ordersStore struct {
	store[entities.Order]
}

func (os *ordersStore) FindManyByClientId(clientId uint) ([]entities.Order, error) {
	orders := []entities.Order{}
	for _, order := range os.Store {
		if order.ClientID == clientId {
			orders = append(orders, order)
		}
	}
	return orders, os.Error
}
