package service

import "github.com/raphael-foliveira/chi-gorm/internal/ports"

var (
	Clients  ports.ClientsService
	Orders   ports.OrdersService
	Products ports.ProductsService
	Jwt      ports.JwtService
)

func Initialize() {
	Clients = NewClients()
	Orders = NewOrders()
	Products = NewProducts()
	Jwt = NewJwt()
}
