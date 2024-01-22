package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/service"
	"gorm.io/gorm"
)

type repository[T interface{}] struct {
	db *gorm.DB
}

func NewRepository[T interface{}](db *gorm.DB) *repository[T] {
	return &repository[T]{db}
}

func (r *repository[T]) List() ([]T, error) {
	entities := []T{}
	return entities, r.db.Find(&entities).Error
}

func (r *repository[T]) Get(id uint) (*T, error) {
	entity := new(T)
	return entity, r.db.Model(new(T)).First(entity, id).Error
}

func (r *repository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *repository[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *repository[T]) Delete(entity *T) error {
	return r.db.Delete(entity).Error
}

type Repositories struct {
	clients  *clients
	products *products
	orders   *orders
}

func (r *Repositories) Clients() service.ClientsRepository {
	return r.clients
}

func (r *Repositories) Products() service.ProductsRepository {
	return r.products
}

func (r *Repositories) Orders() service.OrdersRepository {
	return r.orders
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		clients:  NewClients(db),
		products: NewProducts(db),
		orders:   NewOrders(db),
	}
}
