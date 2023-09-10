package repositories

import (
	"github.com/raphael-foliveira/chi-gorm/pkg/db"
	"github.com/raphael-foliveira/chi-gorm/pkg/interfaces"
	"github.com/raphael-foliveira/chi-gorm/pkg/models"
)

type Products interface {
	interfaces.Repository[models.Product]
}

type products struct {
	db *db.DB
}

func NewProducts(db *db.DB) *products {
	return &products{db}
}

func (r *products) List() ([]models.Product, error) {
	c := []models.Product{}
	return c, r.db.Find(&c).Error
}

func (r *products) Get(id int64) (models.Product, error) {
	product := models.Product{}
	return product, r.db.First(&product, id).Error
}

func (r *products) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *products) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *products) Delete(product *models.Product) error {
	return r.db.Delete(product).Error
}
