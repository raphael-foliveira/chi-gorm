package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"gorm.io/gorm"
)

type Orders Repository[entities.Order]

type orders struct {
	db *gorm.DB
}

func NewOrders(db *gorm.DB) Orders {
	db.AutoMigrate(&entities.Order{})
	return &orders{db}
}

func (r *orders) List() ([]entities.Order, error) {
	orders := []entities.Order{}
	return orders, r.db.Model(&entities.Order{}).Preload("Client").Preload("Product").Find(&orders).Error
}

func (r *orders) Get(id int64) (*entities.Order, error) {
	order := entities.Order{}
	return &order, r.db.Model(&entities.Order{}).Preload("Client").Preload("Product").First(&order, id).Error
}

func (r *orders) Create(order *entities.Order) error {
	return r.db.Create(order).Error
}

func (r *orders) Update(order *entities.Order) error {
	return r.db.Save(order).Error
}

func (r *orders) Delete(order *entities.Order) error {
	return r.db.Delete(order).Error
}
