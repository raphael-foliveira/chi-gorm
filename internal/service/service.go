package service

type Repository[T interface{}] interface {
	List() ([]T, error)
	Get(uint) (*T, error)
	Create(*T) error
	Update(*T) error
	Delete(*T) error
}

type Repositories struct {
	Clients  ClientsRepository
	Products ProductsRepository
	Orders   OrdersRepository
}

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
