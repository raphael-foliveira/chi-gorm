package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
)

type Products struct {
	*repository[entities.Product]
}

func NewProducts() *Products {
	return &Products{newRepository[entities.Product]()}
}

func (r *Products) FindMany(ids []uint) ([]entities.Product, error) {
	products := []entities.Product{}
	return products, db.Find(&products, ids).Error
}
