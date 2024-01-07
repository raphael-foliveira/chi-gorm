package service

import "github.com/raphael-foliveira/chi-gorm/internal/repository"

type Services struct {
	Clients  Clients
	Products Products
	Orders   Orders
}

func NewServices(repositories *repository.Repositories) *Services {
	return &Services{
		Clients:  NewClients(repositories.Clients),
		Products: NewProducts(repositories.Products),
		Orders:   NewOrders(repositories.Orders),
	}
}
