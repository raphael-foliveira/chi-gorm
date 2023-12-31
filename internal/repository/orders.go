package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
)

var Orders OrdersRepository = &orders{&repository[entities.Order]{}}

type OrdersRepository interface {
	Repository[entities.Order]
}

type orders struct {
	*repository[entities.Order]
}
