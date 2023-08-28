package order

import (
	"github.com/raphael-foliveira/chi-gorm/internal/db"
)

type Repository struct {
	db *db.DB
}

func NewRepository(db *db.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) List() ([]Order, error) {
	orders := []Order{}
	return orders, r.db.Model(Order{}).Preload("Product").Find(&orders).Error
}

func (r *Repository) ListByClient(id uint) ([]Order, error) {
	orders := []Order{}
	return orders, r.db.Model(Order{}).Preload("Product").Where("client_id = ?", id).Find(&orders).Error
}

func (r *Repository) Get(id uint64) (Order, error) {
	order := Order{}
	return order, r.db.First(&order, id).Error
}

func (r *Repository) Create(order *Order) error {
	return r.db.Create(order).Error
}

func (r *Repository) Update(order *Order) error {
	return r.db.Save(order).Error
}

func (r *Repository) Delete(order *Order) error {
	return r.db.Delete(order).Error
}
