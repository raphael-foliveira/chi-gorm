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

func (r *products) List() ([]entities.Product, error) {
	c := []entities.Product{}
	return c, r.db.Find(&c).Error
}

func (r *products) Get(id int64) (*entities.Product, error) {
	product := entities.Product{}
	return &product, r.db.First(&product, id).Error
}

func (r *products) Create(product *entities.Product) error {
	return r.db.Create(product).Error
}

func (r *products) Update(product *entities.Product) error {
	return r.db.Save(product).Error
}

func (r *products) Delete(product *entities.Product) error {
	return r.db.Delete(product).Error
}

func (r *products) FindMany(ids []int64) ([]entities.Product, error) {
	products := []entities.Product{}
	return products, r.db.Find(&products, ids).Error
}
