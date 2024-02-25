package controller

import (
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

func Clients() *clients {
	return NewClients(service.Clients())
}

func Orders() *orders {
	return NewOrders(service.Orders())
}

func Products() *products {
	return NewProducts(service.Products())
}

func Users() *users {
	return NewUsers(service.Users())
}
