package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"gorm.io/gorm"
)

type OrdersRepository interface {
	Repository[entities.Order]
	FindManyByClientId(uint) ([]entities.Order, error)
}

type orders struct {
	*repository[entities.Order]
}

func NewOrders(db *gorm.DB) *orders {
	return &orders{&repository[entities.Order]{db}}
}

func (o *orders) FindManyByClientId(clientId uint) ([]entities.Order, error) {
	orders := []entities.Order{}
	return orders, o.db.Where("client_id = ?", clientId).Find(&orders).Error
}
