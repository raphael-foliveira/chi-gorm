package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
	"github.com/raphael-foliveira/chi-gorm/internal/http/routes"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

func clientsRoutes() *chi.Mux {
	return routes.Clients(clientsController())
}
func clientsController() *controller.Clients {
	return controller.NewClients(clientsService())
}
func clientsService() *service.Clients {
	return service.NewClients(clientsRepository(), ordersRepository())
}
func clientsRepository() *repository.Clients {
	return repository.NewClients(database.Db())
}
func ordersRepository() *repository.Orders {
	return repository.NewOrders(database.Db())
}
func productsRepository() *repository.Products {
	return repository.NewProducts(database.Db())
}
func productsService() *service.Products {
	return service.NewProducts(productsRepository())
}
func productsController() *controller.Products {
	return controller.NewProducts(productsService())
}
func productsRoutes() *chi.Mux {
	return routes.Products(productsController())
}
func ordersService() *service.Orders {
	return service.NewOrders(ordersRepository())
}
func ordersController() *controller.Orders {
	return controller.NewOrders(ordersService())
}
func ordersRoutes() *chi.Mux {
	return routes.Orders(ordersController())
}

func healthCheckRoutes() *chi.Mux {
	return routes.HealthCheck()
}
