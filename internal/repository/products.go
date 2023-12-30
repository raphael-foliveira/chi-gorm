package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
)

var Products productsRepository = &products{&repository[entities.Product]{}}

type productsRepository interface {
	Repository[entities.Product]
	FindMany([]uint) ([]entities.Product, error)
}
type products struct {
	*repository[entities.Product]
}

func (r *products) FindMany(ids []uint) ([]entities.Product, error) {
	products := []entities.Product{}
	return products, database.Db.Find(&products, ids).Error
}
