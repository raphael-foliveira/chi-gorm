package container

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/config"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

type Controllers struct {
	HealthCheck *controller.HealthCheck
	Clients     *controller.Clients
	Products    *controller.Products
	Orders      *controller.Orders
}

func NewControllers() *Controllers {
	return &Controllers{
		HealthCheck: controller.NewHealthCheck(),
		Clients:     controller.NewClients(),
		Products:    controller.NewProducts(),
		Orders:      controller.NewOrders(),
	}
}

func (c *Controllers) Mount() {
	c.HealthCheck.Mount()
	c.Clients.Mount()
	c.Products.Mount()
	c.Orders.Mount()
}

func InitializeDependencies(cfg *config.Cfg) *chi.Mux {
	database.Initialize(cfg.DatabaseURL)

	repository.Initialize(database.DB)

	service.Initialize(&service.Config{
		ClientsRepository:  repository.NewClients(),
		OrdersRepository:   repository.NewOrders(),
		ProductsRepository: repository.NewProducts(),
	})

	mainMux := chi.NewRouter()

	controller.Initialize(&controller.Config{
		Router:          mainMux,
		ClientsService:  service.NewClients(),
		OrdersService:   service.NewOrders(),
		ProductsService: service.NewProducts(),
	})

	c := NewControllers()
	c.Mount()

	return mainMux
}
