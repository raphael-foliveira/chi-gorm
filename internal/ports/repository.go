package ports

import "github.com/raphael-foliveira/chi-gorm/internal/entities"

type Repository[T any] interface {
	Create(entity *T) error
	Delete(id uint) error
	Get(id uint) (*T, error)
	List(conds ...any) ([]T, error)
	Update(entity *T) error
}

type ClientsRepository interface {
	Repository[entities.Client]
}

type OrdersRepository interface {
	Repository[entities.Order]
	FindByClient(clientId uint) ([]entities.Order, error)
}

type ProductsRepository interface {
	Repository[entities.Product]
	FindMany(ids []uint) ([]entities.Product, error)
}
