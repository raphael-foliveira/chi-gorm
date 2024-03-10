package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
)

type OrdersRepository interface {
	Repository[entities.Order]
	FindManyByClientId(uint) ([]entities.Order, error)
}

type Orders struct {
	*repository[entities.Order]
}

func NewOrders(db *database.DB) *Orders {
	return &Orders{newRepository[entities.Order](db)}
}

func (o *Orders) FindManyByClientId(clientId uint) ([]entities.Order, error) {
	orders := []entities.Order{}
	return orders, o.db.Where("client_id = ?", clientId).Find(&orders).Error
}
