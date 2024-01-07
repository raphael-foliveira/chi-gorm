package service

import "github.com/raphael-foliveira/chi-gorm/internal/repository"

type Services struct {
	Clients  Clients
	Products Products
	Orders   Orders
}

func NewServices(repositories *repository.Repositories) *Services {
	orders := NewOrders(repositories.Orders)
	clients := NewClients(repositories.Clients, orders)
	products := NewProducts(repositories.Products)
	return &Services{
		Clients:  clients,
		Products: products,
		Orders:   orders,
	}
}
