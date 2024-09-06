package container

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/config"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
	"github.com/raphael-foliveira/chi-gorm/internal/http/server"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
	"gorm.io/gorm"
)

type Repositories struct {
	clients  repository.Clients
	products repository.Products
	orders   repository.Orders
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		clients:  repository.NewClients(db),
		products: repository.NewProducts(db),
		orders:   repository.NewOrders(db),
	}
}

type Services struct {
	clients  *service.Clients
	products *service.Products
	orders   *service.Orders
	jwt      *service.Jwt
}

func NewServices(repositories *Repositories) *Services {
	return &Services{
		clients:  service.NewClients(repositories.clients, repositories.orders),
		products: service.NewProducts(repositories.products),
		orders:   service.NewOrders(repositories.orders),
		jwt:      service.NewJwt(),
	}
}

type Controllers struct {
	clients     *controller.Clients
	products    *controller.Products
	orders      *controller.Orders
	healthcheck *controller.HealthCheck
}

func NewControllers(services *Services) *Controllers {
	return &Controllers{
		clients:     controller.NewClients(services.clients),
		products:    controller.NewProducts(services.products),
		orders:      controller.NewOrders(services.orders),
		healthcheck: controller.NewHealthCheck(),
	}
}

func (c *Controllers) Mount(app *chi.Mux) {
	c.clients.Mount(app)
	c.products.Mount(app)
	c.orders.Mount(app)
	c.healthcheck.Mount(app)
}

func InitializeDependencies(cfg *config.Cfg) (*chi.Mux, func()) {
	db := database.New(cfg.DatabaseURL)

	repositories := NewRepositories(db)
	services := NewServices(repositories)
	controllers := NewControllers(services)

	app := server.CreateMainRouter()

	controllers.Mount(app)

	return app, func() {
		database.Close(db)
	}
}
