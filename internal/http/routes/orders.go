package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controllers"
)

func Orders() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", wrap(controllers.Orders.List))
	router.Post("/", wrap(controllers.Orders.Create))
	router.Get("/{id}", wrap(controllers.Orders.Get))
	router.Delete("/{id}", wrap(controllers.Orders.Delete))
	router.Put("/{id}", wrap(controllers.Orders.Update))
	return router
}
