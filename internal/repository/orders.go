package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"gorm.io/gorm"
)

type OrdersRepository interface {
	Repository[entities.Order]
	FindManyByClientId(uint) ([]entities.Order, error)
}

type Orders struct {
	*repository[entities.Order]
}

func NewOrders(db *gorm.DB) *Orders {
	return &Orders{newRepository[entities.Order](db)}
}

func (o *Orders) FindManyByClientId(clientId uint) ([]entities.Order, error) {
	orders := []entities.Order{}
	return orders, o.db.Where("client_id = ?", clientId).Find(&orders).Error
}
