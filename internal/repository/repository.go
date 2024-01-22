package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/service"
	"gorm.io/gorm"
)

type Repository[T interface{}] struct {
	db *gorm.DB
}

func NewRepository[T interface{}](db *gorm.DB) *Repository[T] {
	return &Repository[T]{db}
}

func (r *Repository[T]) List() ([]T, error) {
	entities := []T{}
	return entities, r.db.Find(&entities).Error
}

func (r *Repository[T]) Get(id uint) (*T, error) {
	entity := new(T)
	return entity, r.db.Model(new(T)).First(entity, id).Error
}

func (r *Repository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *Repository[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *Repository[T]) Delete(entity *T) error {
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
