package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
)

type products struct {
	*repository[entities.Product]
}

func NewProducts() *products {
	return &products{newRepository[entities.Product]()}
}

func (r *products) FindMany(ids []uint) ([]entities.Product, error) {
	products := []entities.Product{}
	return products, database.DB.Find(&products, ids).Error
}
