package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
)

func Products() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", wrap(controller.Products.List))
	router.Post("/", wrap(controller.Products.Create))
	router.Get("/{id}", wrap(controller.Products.Get))
	router.Delete("/{id}", wrap(controller.Products.Delete))
	router.Put("/{id}", wrap(controller.Products.Update))
	return router
}
