package ports

import (
	"github.com/raphael-foliveira/chi-gorm/internal/dto"
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
)

type ClientsService interface {
	Create(schema *schemas.CreateClient) (*entities.Client, error)
	Delete(id uint) error
	Get(id uint) (*entities.Client, error)
	GetOrders(clientId uint) ([]entities.Order, error)
	List() ([]entities.Client, error)
	Update(id uint, schema *schemas.UpdateClient) (*entities.Client, error)
}

type JwtService interface {
	Sign(payload *dto.JwtPayload) (string, error)
	Verify(token string) (*dto.JwtPayload, error)
}

type OrdersService interface {
	Create(schema *schemas.CreateOrder) (*entities.Order, error)
	Delete(id uint) error
	Get(id uint) (*entities.Order, error)
	List() ([]entities.Order, error)
	Update(id uint, schema *schemas.UpdateOrder) (*entities.Order, error)
}

type ProductsService interface {
	Create(schema *schemas.CreateProduct) (*entities.Product, error)
	Delete(id uint) error
	Get(id uint) (*entities.Product, error)
	List() ([]entities.Product, error)
	Update(id uint, schema *schemas.UpdateProduct) (*entities.Product, error)
}
