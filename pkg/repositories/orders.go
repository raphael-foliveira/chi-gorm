package repositories

import (
	"github.com/raphael-foliveira/chi-gorm/pkg/db"
	"github.com/raphael-foliveira/chi-gorm/pkg/models"
)

type Orders struct {
	db *db.DB
}

func NewOrders(db *db.DB) *Orders {
	return &Orders{db}
}

func (r *Orders) List() ([]models.Order, error) {
	orders := []models.Order{}
	return orders, r.db.Model(models.Order{}).Preload("Product").Find(&orders).Error
}

func (r *Orders) ListByClient(id uint) ([]models.Order, error) {
	orders := []models.Order{}
	return orders, r.db.Model(models.Order{}).Preload("Product").Where("client_id = ?", id).Find(&orders).Error
}

func (r *Orders) Get(id uint64) (models.Order, error) {
	order := models.Order{}
	return order, r.db.First(&order, id).Error
}

func (r *Orders) Create(order *models.Order) error {
	return r.db.Create(order).Error
}

func (r *Orders) Update(order *models.Order) error {
	return r.db.Save(order).Error
}

func (r *Orders) Delete(order *models.Order) error {
	return r.db.Delete(order).Error
}
