package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
)

var Products Repository[entities.Product] = &products{&repository[entities.Product]{}}

type products struct {
	*repository[entities.Product]
}

func (r *products) FindMany(ids []int64) ([]entities.Product, error) {
	products := []entities.Product{}
	return products, database.Db.Find(&products, ids).Error
}
