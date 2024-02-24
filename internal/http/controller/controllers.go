package controller

import (
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

type Controllers struct {
	Clients  *ClientsController
	Orders   *OrdersController
	Products *ProductsController
}

func NewControllers(services *service.Services) *Controllers {
	return &Controllers{
		Clients:  NewClients(services.Clients),
		Orders:   NewOrders(services.Orders),
		Products: NewProducts(services.Products),
	}
}
