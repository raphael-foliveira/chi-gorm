package ports

import "github.com/raphael-foliveira/chi-gorm/internal/entities"

type ClientsRepository interface {
	Create(entity *entities.Client) error
	Delete(entity *entities.Client) error
	Get(id uint) (*entities.Client, error)
	List(conds ...any) ([]entities.Client, error)
	Update(entity *entities.Client) error
}

type OrdersRepository interface {
	Create(entity *entities.Order) error
	Delete(entity *entities.Order) error
	FindManyByClientId(clientId uint) ([]entities.Order, error)
	Get(id uint) (*entities.Order, error)
	List(conds ...any) ([]entities.Order, error)
	Update(entity *entities.Order) error
}

type ProductsRepository interface {
	Create(entity *entities.Product) error
	Delete(entity *entities.Product) error
	FindMany(ids []uint) ([]entities.Product, error)
	Get(id uint) (*entities.Product, error)
	List(conds ...any) ([]entities.Product, error)
	Update(entity *entities.Product) error
}
