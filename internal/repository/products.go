package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"gorm.io/gorm"
)

type ProductsRepository interface {
	Repository[entities.Product]
	FindMany([]uint) ([]entities.Product, error)
}

func Products() *products {
	return NewProducts(database.Db())
}

type products struct {
	*repository[entities.Product]
}

func NewProducts(db *gorm.DB) *products {
	return &products{newRepository[entities.Product](db)}
}

func (r *products) FindMany(ids []uint) ([]entities.Product, error) {
	products := []entities.Product{}
	return products, r.db.Find(&products, ids).Error
}
