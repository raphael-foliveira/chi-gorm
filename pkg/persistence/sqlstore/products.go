package sqlstore

import (
	"github.com/raphael-foliveira/chi-gorm/pkg/interfaces"
	"github.com/raphael-foliveira/chi-gorm/pkg/models"
)

type Products interface {
	interfaces.Store[models.Product]
	FindMany(ids []int64) ([]models.Product, error)
}

type products struct{}

func NewProducts() Products {
	db.AutoMigrate(&models.Product{})
	return &products{}
}

func (r *products) List() ([]models.Product, error) {
	c := []models.Product{}
	return c, db.Find(&c).Error
}

func (r *products) Get(id int64) (*models.Product, error) {
	product := models.Product{}
	return &product, db.First(&product, id).Error
}

func (r *products) Create(product *models.Product) error {
	return db.Create(product).Error
}

func (r *products) Update(product *models.Product) error {
	return db.Save(product).Error
}

func (r *products) Delete(product *models.Product) error {
	return db.Delete(product).Error
}

func (r *products) FindMany(ids []int64) ([]models.Product, error) {
	products := []models.Product{}
	return products, db.Find(&products, ids).Error
}
