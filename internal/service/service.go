package service

import (
	"github.com/raphael-foliveira/chi-gorm/internal/config"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
)

type Services struct {
	ClientsService  *Clients
	ProductsService *Products
	OrdersService   *Orders
	JwtService      *Jwt
}

func NewServices(repositories *repository.Repositories, cfg *config.Cfg) *Services {
	return &Services{
		ClientsService:  NewClients(repositories.ClientsRepository, repositories.OrdersRepository),
		ProductsService: NewProducts(repositories.ProductsRepository),
		OrdersService:   NewOrders(repositories.OrdersRepository),
		JwtService:      NewJwt(cfg.JwtSecret),
	}
}
