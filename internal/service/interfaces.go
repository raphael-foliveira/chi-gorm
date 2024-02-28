package service

import "github.com/raphael-foliveira/chi-gorm/internal/entities"

type Repository[T interface{}] interface {
	List(...interface{}) ([]T, error)
	Get(uint) (*T, error)
	Create(*T) error
	Update(*T) error
	Delete(*T) error
}

type ClientsRepository interface {
	Repository[entities.Client]
}
type OrdersRepository interface {
	Repository[entities.Order]
	FindManyByClientId(uint) ([]entities.Order, error)
}

type ProductsRepository interface {
	Repository[entities.Product]
	FindMany([]uint) ([]entities.Product, error)
}

type UsersRepository interface {
	Repository[entities.User]
	FindOneByEmail(string) (*entities.User, error)
}
