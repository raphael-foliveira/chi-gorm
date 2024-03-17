package container

import (
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
	"github.com/raphael-foliveira/chi-gorm/internal/http/routes"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

var ClientsRoutes = func() *routes.Router {
	return routes.Clients(ClientsController())
}

var ClientsController = func() *controller.Clients {
	return controller.NewClients(ClientsService())
}

var ClientsService = func() *service.Clients {
	return service.NewClients(ClientsRepository(), OrdersRepository())
}

var OrdersRepository = func() repository.OrdersRepository {
	return repository.NewOrders(database.Db())
}

var ProductsRepository = func() repository.ProductsRepository {
	return repository.NewProducts(database.Db())
}

var ProductsService = func() *service.Products {
	return service.NewProducts(ProductsRepository())
}

var ProductsController = func() *controller.Products {
	return controller.NewProducts(ProductsService())
}

var ProductsRoutes = func() *routes.Router {
	return routes.Products(ProductsController())
}

var ClientsRepository = func() repository.ClientsRepository {
	return repository.NewClients(database.Db())
}

var OrdersService = func() *service.Orders {
	return service.NewOrders(OrdersRepository())
}

var OrdersController = func() *controller.Orders {
	return controller.NewOrders(OrdersService())
}

var OrdersRoutes = func() *routes.Router {
	return routes.Orders(OrdersController())
}

var HealthCheckRoutes = func() *routes.Router {
	return routes.HealthCheck()
}
