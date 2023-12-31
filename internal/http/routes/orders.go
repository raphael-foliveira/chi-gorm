package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
)

func Orders() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", wrap(controller.Orders.List))
	router.Post("/", wrap(controller.Orders.Create))
	router.Get("/{id}", wrap(controller.Orders.Get))
	router.Delete("/{id}", wrap(controller.Orders.Delete))
	router.Put("/{id}", wrap(controller.Orders.Update))
	return router
}
