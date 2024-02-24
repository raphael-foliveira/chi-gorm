package service

import "github.com/raphael-foliveira/chi-gorm/internal/repository"

func Orders() *OrdersService {
	return NewOrders(repository.Orders())
}

func Products() *ProductsService {
	return NewProducts(repository.Products())
}

func Clients() *ClientsService {
	return NewClients(repository.Clients(), repository.Orders())
}

func Jwt() *JwtService {
	return NewJwt()
}
