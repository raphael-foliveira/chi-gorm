package container

import (
	"github.com/raphael-foliveira/chi-gorm/internal/config"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
	"github.com/raphael-foliveira/chi-gorm/internal/http/middleware"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

var ClientsController = func() *controller.Clients {
	return controller.NewClients(ClientsService())
}

var ClientsService = func() *service.Clients {
	return service.NewClients(ClientsRepository(), OrdersRepository())
}

var OrdersRepository = func() repository.OrdersRepository {
	return repository.NewOrders(Db())
}

var ProductsRepository = func() repository.ProductsRepository {
	return repository.NewProducts(Db())
}

var ProductsService = func() *service.Products {
	return service.NewProducts(ProductsRepository())
}

var ProductsController = func() *controller.Products {
	return controller.NewProducts(ProductsService())
}

var ClientsRepository = func() repository.ClientsRepository {
	return repository.NewClients(Db())
}

var OrdersService = func() *service.Orders {
	return service.NewOrders(OrdersRepository())
}

var OrdersController = func() *controller.Orders {
	return controller.NewOrders(OrdersService())
}

var JwtService = func() *service.Jwt {
	return service.NewJwt()
}

var AuthMiddleware = func() *middleware.AuthMiddleware {
	return middleware.NewAuthMiddleware(JwtService())
}

var HealthcheckController = func() *controller.HealthcheckController {
	return controller.NewHealthCheck()
}

var Config = func() *config.Cfg {
	return config.Config()
}

var Db = func() *database.DB {
	return database.Db(Config().DatabaseURL)
}
