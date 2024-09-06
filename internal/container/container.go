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

var (
	db                    *gorm.DB
	clientsController     *controller.Clients
	productsController    *controller.Products
	ordersController      *controller.Orders
	healthcheckController *controller.HealthCheck
	clientsService        *service.Clients
	productsService       *service.Products
	ordersService         *service.Orders
	clientsRepository     repository.Clients
	productsRepository    repository.Products
	ordersRepository      repository.Orders
	app                   *chi.Mux
)

func InitializeDependencies(cfg *config.Cfg) *chi.Mux {
	db = database.Db(cfg.DatabaseURL)

	clientsRepository = repository.NewClients(db)
	productsRepository = repository.NewProducts(db)
	ordersRepository = repository.NewOrders(db)
	clientsService = service.NewClients(clientsRepository, ordersRepository)
	productsService = service.NewProducts(productsRepository)
	ordersService = service.NewOrders(ordersRepository)
	clientsController = controller.NewClients(clientsService)
	productsController = controller.NewProducts(productsService)
	ordersController = controller.NewOrders(ordersService)
	healthcheckController = controller.NewHealthCheck()

	app = server.CreateMainRouter()

	clientsController.Mount(app)
	productsController.Mount(app)
	ordersController.Mount(app)
	healthcheckController.Mount(app)
	return app
}
