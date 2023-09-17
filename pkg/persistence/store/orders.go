package store

import (
	"github.com/raphael-foliveira/chi-gorm/pkg/interfaces"
	"github.com/raphael-foliveira/chi-gorm/pkg/models"
	"gorm.io/gorm"
)

type Orders interface {
	interfaces.Store[models.Order]
}

type orders struct {
	db *gorm.DB
}

func NewOrders(db *gorm.DB) Orders {
	db.AutoMigrate(&models.Order{})
	return &orders{db}
}

func (r *orders) List() ([]models.Order, error) {
	orders := []models.Order{}
	return orders, r.db.Model(models.Order{}).Preload("Product").Find(&orders).Error
}

func (r *orders) Get(id int64) (*models.Order, error) {
	order := models.Order{}
	return &order, r.db.First(&order, id).Error
}

func (r *orders) Create(order *models.Order) error {
	return r.db.Create(order).Error
}

func (r *orders) Update(order *models.Order) error {
	return r.db.Save(order).Error
}

func (r *orders) Delete(order *models.Order) error {
	return r.db.Delete(order).Error
}
