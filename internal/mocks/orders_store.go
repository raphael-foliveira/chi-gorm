package mocks

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
)

var OrdersStore = &OrdersStoreImpl{store[entities.Order]{}}

type OrdersStoreImpl struct {
	store[entities.Order]
}
