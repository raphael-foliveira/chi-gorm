package container

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
	"github.com/raphael-foliveira/chi-gorm/internal/http/middleware"
	"github.com/raphael-foliveira/chi-gorm/internal/http/routes"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

var ClientsRoutes = func() *chi.Mux {
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

var ProductsRoutes = func() *chi.Mux {
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

var OrdersRoutes = func() *chi.Mux {
	return routes.Orders(OrdersController())
}

var HealthCheckRoutes = func() *chi.Mux {
	return routes.HealthCheck()
}

var JwtService = func() *service.Jwt {
	return service.NewJwt()
}

var AuthMiddleware = func() *middleware.AuthMiddleware {
	return middleware.NewAuthMiddleware(JwtService())
}
