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
	Clients  repository.Clients
	Products repository.Products
	Orders   repository.Orders
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Clients:  repository.NewClients(db),
		Products: repository.NewProducts(db),
		Orders:   repository.NewOrders(db),
	}
}

type Services struct {
	Clients  *service.Clients
	Products *service.Products
	Orders   *service.Orders
	Jwt      *service.Jwt
}

func NewServices(repositories *Repositories) *Services {
	return &Services{
		Clients:  service.NewClients(repositories.Clients, repositories.Orders),
		Products: service.NewProducts(repositories.Products),
		Orders:   service.NewOrders(repositories.Orders),
		Jwt:      service.NewJwt(),
	}
}

type Controllers struct {
	Clients     *controller.Clients
	Products    *controller.Products
	Orders      *controller.Orders
	Healthcheck *controller.HealthCheck
}

func NewControllers(services *Services) *Controllers {
	return &Controllers{
		Clients:     controller.NewClients(services.Clients),
		Products:    controller.NewProducts(services.Products),
		Orders:      controller.NewOrders(services.Orders),
		Healthcheck: controller.NewHealthCheck(),
	}
}

func (c *Controllers) Mount(app *chi.Mux) {
	c.Clients.Mount(app)
	c.Products.Mount(app)
	c.Orders.Mount(app)
	c.Healthcheck.Mount(app)
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
