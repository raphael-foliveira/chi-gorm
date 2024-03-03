package controller

import (
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

func Clients() *clientsController {
	return NewClientsController(service.Clients())
}

func Orders() *ordersController {
	return NewOrdersController(service.Orders())
}

func Products() *productsController {
	return NewProductsController(service.Products())
}

func Users() *usersController {
	return NewUsersController(service.Users())
}

func Auth() *authController {
	return NewAuthController(service.Auth())
}

type ControllerFunc func(http.ResponseWriter, *http.Request) error
