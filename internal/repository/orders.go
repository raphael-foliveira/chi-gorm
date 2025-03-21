package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/ports"
	"gorm.io/gorm"
)

var _ ports.OrdersRepository = &Orders{}

type Orders struct {
	*Repository[entities.Order]
}

func NewOrders(db *gorm.DB) *Orders {
	return &Orders{NewRepository[entities.Order](db)}
}

func (o *Orders) FindByClient(clientId uint) ([]entities.Order, error) {
	orders := []entities.Order{}
	return orders, database.DB.Where("client_id = ?", clientId).Find(&orders).Error
}
