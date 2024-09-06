package service

var (
	Clients  *clients
	Orders   *orders
	Products *products
	Jwt      *jwtImpl
)

func Initialize() {
	Clients = NewClients()
	Orders = NewOrders()
	Products = NewProducts()
	Jwt = NewJwt()
}
