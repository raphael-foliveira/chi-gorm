package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"gorm.io/gorm"
)

type Orders interface {
	Repository[entities.Order]
}

type orders struct {
	*repository[entities.Order]
}

func NewOrders(db *gorm.DB) Orders {
	return &orders{&repository[entities.Order]{db}}
}
