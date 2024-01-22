package service

type Services struct {
	Clients  *Clients
	Products *Products
	Orders   *Orders
	Jwt      *Jwt
}

func NewServices(repositories *Repositories) *Services {
	return &Services{
		Products: NewProducts(repositories.Products),
		Orders:   NewOrders(repositories.Orders),
		Clients:  NewClients(repositories.Clients, repositories.Orders),
		Jwt:      NewJwt(),
	}
}

