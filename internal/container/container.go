package container

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/config"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
)

func InitializeDependencies(cfg *config.Cfg) *chi.Mux {
	db := database.Initialize(cfg.DatabaseURL)
	repository.SetDB(db)

	healthcheckController := controller.NewHealthCheck()
	clientsController := controller.NewClients()
	productsController := controller.NewProducts()
	ordersController := controller.NewOrders()

	clientsController.Mount()
	productsController.Mount()
	ordersController.Mount()
	healthcheckController.Mount()
	return controller.GetApp()
}
