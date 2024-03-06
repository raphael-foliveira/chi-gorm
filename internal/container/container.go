package container

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
	"github.com/raphael-foliveira/chi-gorm/internal/http/routes"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

func ClientsRoutes() *chi.Mux {
	return routes.Clients(ClientsController())
}

func ClientsController() *controller.Clients {
	return controller.NewClients(ClientsService())
}

func ClientsService() *service.Clients {
	return service.NewClients(ClientsRepository(), OrdersRepository())
}

func OrdersRepository() *repository.Orders {
	return repository.NewOrders(database.Db())
}

func ProductsRepository() *repository.Products {
	return repository.NewProducts(database.Db())
}

func ProductsService() *service.Products {
	return service.NewProducts(ProductsRepository())
}

func ProductsController() *controller.Products {
	return controller.NewProducts(ProductsService())
}

func ProductsRoutes() *chi.Mux {
	return routes.Products(ProductsController())
}

func ClientsRepository() *repository.Clients {
	return repository.NewClients(database.Db())
}

func OrdersService() *service.Orders {
	return service.NewOrders(OrdersRepository())
}

func OrdersController() *controller.Orders {
	return controller.NewOrders(OrdersService())
}

func OrdersRoutes() *chi.Mux {
	return routes.Orders(OrdersController())
}

func HealthCheckRoutes() *chi.Mux {
	return routes.HealthCheck()
}
