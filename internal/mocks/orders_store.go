package mocks

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
)

type OrdersStore struct {
	store[entities.Order]
}

func NewOrdersStore() *OrdersStore {
	return &OrdersStore{newStore[entities.Order]()}
}
