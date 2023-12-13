package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"gorm.io/gorm"
)

type Products Repository[entities.Product]

type products struct {
	db *gorm.DB
}

func NewProducts(db *gorm.DB) Products {
	db.AutoMigrate(&entities.Product{})
	return &products{db}
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
