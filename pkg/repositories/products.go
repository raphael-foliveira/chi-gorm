package repositories

import (
	"github.com/raphael-foliveira/chi-gorm/pkg/db"
	"github.com/raphael-foliveira/chi-gorm/pkg/models"
)

type Products struct {
	db *db.DB
}

func NewProducts(db *db.DB) *Products {
	return &Products{db}
}

func (r *Products) List() ([]models.Product, error) {
	c := []models.Product{}
	return c, r.db.Find(&c).Error
}

func (r *Products) Get(id int64) (models.Product, error) {
	product := models.Product{}
	return product, r.db.First(&product, id).Error
}

func (r *Products) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *Products) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *Products) Delete(product *models.Product) error {
	return r.db.Delete(product).Error
}
