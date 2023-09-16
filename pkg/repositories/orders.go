package repositories

import (
	"github.com/raphael-foliveira/chi-gorm/pkg/db"
	"github.com/raphael-foliveira/chi-gorm/pkg/interfaces"
	"github.com/raphael-foliveira/chi-gorm/pkg/models"
)

type Orders interface {
	interfaces.Repository[models.Order]
}

type orders struct {
	db *db.DB
}

func NewOrders(db *db.DB) *orders {
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
