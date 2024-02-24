package service

import "github.com/raphael-foliveira/chi-gorm/internal/repository"

type Services struct {
	Clients  *ClientsService
	Products *ProductsService
	Orders   *OrdersService
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
