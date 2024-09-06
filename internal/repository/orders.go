package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
)

type orders struct {
	*repository[entities.Order]
}

func NewOrders() *orders {
	return &orders{newRepository[entities.Order]()}
}

func (o *orders) FindManyByClientId(clientId uint) ([]entities.Order, error) {
	orders := []entities.Order{}
	return orders, database.DB.Where("client_id = ?", clientId).Find(&orders).Error
}
