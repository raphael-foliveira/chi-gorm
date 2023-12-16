package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"gorm.io/gorm"
)

type Products interface {
	Repository[entities.Product]
	FindMany(ids []int64) ([]entities.Product, error)
}

type products struct {
	*repository[entities.Product]
}

func NewProducts(db *gorm.DB) Products {
	db.AutoMigrate(&entities.Product{})
	return &products{newRepository[entities.Product](db)}
}

func (r *products) FindMany(ids []int64) ([]entities.Product, error) {
	products := []entities.Product{}
	return products, r.db.Find(&products, ids).Error
}
