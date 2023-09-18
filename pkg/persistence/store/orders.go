package store

import (
	"github.com/raphael-foliveira/chi-gorm/pkg/interfaces"
	"github.com/raphael-foliveira/chi-gorm/pkg/models"
	"github.com/raphael-foliveira/chi-gorm/pkg/persistence/db"
)

type Orders interface {
	interfaces.Store[models.Order]
	GetByClientId(clientId int64) ([]models.Order, error)
}

type orders struct{}

func NewOrders() Orders {
	db.Db.AutoMigrate(&models.Order{})
	return &orders{}
}

func (r *orders) List() ([]models.Order, error) {
	orders := []models.Order{}
	return orders, db.Db.Find(&orders).Error
}

func (r *orders) Get(id int64) (*models.Order, error) {
	order := models.Order{}
	return &order, db.Db.First(&order, id).Error
}

func (r *orders) Create(order *models.Order) error {
	return db.Db.Create(order).Error
}

func (r *orders) Update(order *models.Order) error {
	return db.Db.Save(order).Error
}

func (r *orders) Delete(order *models.Order) error {
	return db.Db.Delete(order).Error
}

func (r *orders) GetByClientId(clientId int64) ([]models.Order, error) {
	orders := []models.Order{}
	return orders, db.Db.Where("client_id = ?", clientId).Find(&orders).Error
}
