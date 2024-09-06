package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
)

type Orders struct {
	*Repository[entities.Order]
}

func NewOrders() *Orders {
	return &Orders{New[entities.Order]()}
}

func (o *Orders) FindManyByClientId(clientId uint) ([]entities.Order, error) {
	orders := []entities.Order{}
	return orders, db.Where("client_id = ?", clientId).Find(&orders).Error
}
