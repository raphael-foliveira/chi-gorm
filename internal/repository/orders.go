package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"gorm.io/gorm"
)

type Orders Repository[entities.Order]

type orders struct {
	*repository[entities.Order]
}

func NewOrders(db *gorm.DB) Orders {
	db.AutoMigrate(&entities.Order{})
	return &orders{newRepository[entities.Order](db)}
}
