package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
)

type Routers struct {
	Products    *chi.Mux
	Orders      *chi.Mux
	Clients     *chi.Mux
	HealthCheck http.HandlerFunc
}

func NewRouters(controllers *controller.Controllers) *Routers {
	return &Routers{
		Products:    Products(controllers.Products),
		Orders:      Orders(controllers.Orders),
		Clients:     Clients(controllers.Clients),
		HealthCheck: wrap(healthCheck),
	}
}
