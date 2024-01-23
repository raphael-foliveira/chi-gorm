package service

import "github.com/raphael-foliveira/chi-gorm/internal/interfaces"

type Services struct {
	Clients  *Clients
	Products *Products
	Orders   *Orders
	Jwt      *Jwt
}

func NewServices(repositories *interfaces.Repositories) *Services {
	return &Services{
		Products: NewProducts(repositories.Products),
		Orders:   NewOrders(repositories.Orders),
		Clients:  NewClients(repositories.Clients, repositories.Orders),
		Jwt:      NewJwt(),
	}
}
