package container

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/config"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
	"github.com/raphael-foliveira/chi-gorm/internal/http/server"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

func InitializeDependencies(cfg *config.Cfg) *chi.Mux {
	db := database.Initialize(cfg.DatabaseURL)

	healthcheckController := controller.NewHealthCheck()
	clientsRepository := repository.NewClients(db)
	productsRepository := repository.NewProducts(db)
	ordersRepository := repository.NewOrders(db)
	clientsService := service.NewClients(clientsRepository, ordersRepository)
	productsService := service.NewProducts(productsRepository)
	ordersService := service.NewOrders(ordersRepository)
	clientsController := controller.NewClients(clientsService)
	productsController := controller.NewProducts(productsService)
	ordersController := controller.NewOrders(ordersService)

	app := server.CreateMainRouter()

	clientsController.Mount(app)
	productsController.Mount(app)
	ordersController.Mount(app)
	healthcheckController.Mount(app)
	return app
}
