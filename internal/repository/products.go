package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/interfaces"
	"github.com/raphael-foliveira/chi-gorm/internal/models"
	"gorm.io/gorm"
)

type Products interface {
	interfaces.Repository[models.Product]
	FindMany(ids []int64) ([]models.Product, error)
}

type products struct {
	db *gorm.DB
}

func NewProducts(db *gorm.DB) Products {
	db.AutoMigrate(&models.Product{})
	return &products{db}
}

func (r *products) List() ([]models.Product, error) {
	c := []models.Product{}
	return c, r.db.Find(&c).Error
}

func (r *products) Get(id int64) (*models.Product, error) {
	product := models.Product{}
	return &product, r.db.First(&product, id).Error
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

func (r *products) FindMany(ids []int64) ([]models.Product, error) {
	products := []models.Product{}
	return products, r.db.Find(&products, ids).Error
}