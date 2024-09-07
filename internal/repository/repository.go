package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/ports"
)

var (
	Clients  ports.ClientsRepository  = NewClients()
	Products ports.ProductsRepository = NewProducts()
	Orders   ports.OrdersRepository   = NewOrders()
)

type repository[T any] struct{}

func newRepository[T any]() *repository[T] {
	return &repository[T]{}
}

func (r *repository[T]) List(conds ...any) ([]T, error) {
	entities := []T{}
	return entities, database.DB.Find(&entities, conds...).Error
}

func (r *repository[T]) Get(id uint) (*T, error) {
	entity := new(T)
	return entity, database.DB.Model(new(T)).First(entity, id).Error
}

func (r *repository[T]) Create(entity *T) error {
	return database.DB.Create(entity).Error
}

func (r *repository[T]) Update(entity *T) error {
	return database.DB.Save(entity).Error
}

func (r *repository[T]) Delete(entity *T) error {
	return database.DB.Delete(entity).Error
}
