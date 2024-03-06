package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"gorm.io/gorm"
)

type ProductsRepository interface {
	Repository[entities.Product]
	FindMany([]uint) ([]entities.Product, error)
}

type Products struct {
	*repository[entities.Product]
}

func NewProducts(db *gorm.DB) *Products {
	return &Products{newRepository[entities.Product](db)}
}

func (r *Products) FindMany(ids []uint) ([]entities.Product, error) {
	products := []entities.Product{}
	return products, r.db.Find(&products, ids).Error
}
