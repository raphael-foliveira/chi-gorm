package service

import "github.com/raphael-foliveira/chi-gorm/internal/ports"

var (
	Clients  ports.ClientsService  = NewClients()
	Orders   ports.OrdersService   = NewOrders()
	Products ports.ProductsService = NewProducts()
	Jwt      ports.JwtService      = NewJwt()
)
