package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"gorm.io/gorm"
)


type ordersRepository struct {
	*repository[entities.Order]
}

func NewOrdersRepository(db *gorm.DB) *ordersRepository {
	return &ordersRepository{newRepository[entities.Order](db)}
}

func (o *ordersRepository) FindManyByClientId(clientId uint) ([]entities.Order, error) {
	orders := []entities.Order{}
	return orders, o.db.Where("client_id = ?", clientId).Find(&orders).Error
}
