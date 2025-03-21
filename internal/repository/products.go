package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/ports"
	"gorm.io/gorm"
)

var _ ports.ProductsRepository = &Products{}

type Products struct {
	*Repository[entities.Product]
}

func NewProducts(db *gorm.DB) *Products {
	return &Products{NewRepository[entities.Product](db)}
}

func (r *Products) FindMany(ids []uint) ([]entities.Product, error) {
	products := []entities.Product{}
	return products, database.DB.Find(&products, ids).Error
}
