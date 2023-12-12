package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/interfaces"
	"gorm.io/gorm"
)

type Orders interface {
	interfaces.Repository[entities.Order]
	GetByClientId(clientId int64) ([]entities.Order, error)
}

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

func (r *orders) GetByClientId(clientId int64) ([]entities.Order, error) {
	orders := []entities.Order{}
	return orders, r.db.Where("client_id = ?", clientId).Find(&orders).Error
}
