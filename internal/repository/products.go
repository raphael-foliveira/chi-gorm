package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"gorm.io/gorm"
)

type ProductsRepository interface {
	Repository[entities.Product]
	FindMany([]uint) ([]entities.Product, error)
}

type productsRepository struct {
	*repository[entities.Product]
}

func NewProductsRepository(db *gorm.DB) *productsRepository {
	return &productsRepository{newRepository[entities.Product](db)}
}

func (r *productsRepository) FindMany(ids []uint) ([]entities.Product, error) {
	products := []entities.Product{}
	return products, r.db.Find(&products, ids).Error
}
