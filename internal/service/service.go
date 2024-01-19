package service

import (
	"errors"

	"github.com/raphael-foliveira/chi-gorm/internal/repository"
)

type Services struct {
	Clients  *Clients
	Products *Products
	Orders   *Orders
	Jwt      *Jwt
}

func NewServices(repositories *repository.Repositories) *Services {
	return &Services{
		Products: NewProducts(repositories.Products),
		Orders:   NewOrders(repositories.Orders),
		Clients:  NewClients(repositories.Clients, repositories.Orders),
		Jwt:      NewJwt(),
	}
}

var ErrNotFound = errors.New("not found")
