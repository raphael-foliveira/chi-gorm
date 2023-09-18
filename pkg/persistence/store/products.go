package store

import (
	"github.com/raphael-foliveira/chi-gorm/pkg/interfaces"
	"github.com/raphael-foliveira/chi-gorm/pkg/models"
	"github.com/raphael-foliveira/chi-gorm/pkg/persistence/db"
)

type Products interface {
	interfaces.Store[models.Product]
	FindMany(ids []int64) ([]models.Product, error)
}

type products struct{}

func NewProducts() Products {
	db.Db.AutoMigrate(&models.Product{})
	return &products{}
}

func (r *products) List() ([]models.Product, error) {
	c := []models.Product{}
	return c, db.Db.Find(&c).Error
}

func (r *products) Get(id int64) (*models.Product, error) {
	product := models.Product{}
	return &product, db.Db.First(&product, id).Error
}

func (r *products) Create(product *models.Product) error {
	return db.Db.Create(product).Error
}

func (r *products) Update(product *models.Product) error {
	return db.Db.Save(product).Error
}

func (r *products) Delete(product *models.Product) error {
	return db.Db.Delete(product).Error
}

func (r *products) FindMany(ids []int64) ([]models.Product, error) {
	products := []models.Product{}
	return products, db.Db.Find(&products, ids).Error
}
