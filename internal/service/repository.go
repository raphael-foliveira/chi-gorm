package service

import "github.com/raphael-foliveira/chi-gorm/internal/entities"

type Repository[T interface{}] interface {
	List() ([]T, error)
	Get(uint) (*T, error)
	Create(*T) error
	Update(*T) error
	Delete(*T) error
}

type Repositories struct {
	Clients  ClientsRepository
	Products ProductsRepository
	Orders   OrdersRepository
}

type ClientsRepository interface {
	Repository[entities.Client]
}

type ProductsRepository interface {
	Repository[entities.Product]
	FindMany([]uint) ([]entities.Product, error)
}

type OrdersRepository interface {
	Repository[entities.Order]
	FindManyByClientId(uint) ([]entities.Order, error)
}
