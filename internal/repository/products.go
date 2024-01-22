package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"gorm.io/gorm"
)

type products struct {
	*Repository[entities.Product]
}

func NewProducts(db *gorm.DB) *products {
	return &products{&Repository[entities.Product]{db}}
}

func (r *products) FindMany(ids []uint) ([]entities.Product, error) {
	products := []entities.Product{}
	return products, r.db.Find(&products, ids).Error
}
