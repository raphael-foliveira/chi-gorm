package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/pkg/controllers"
)

func Products(controller *controllers.Products) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", wrap(controller.List))
	router.Post("/", wrap(controller.Create))
	router.Get("/{id}", wrap(controller.Get))
	router.Delete("/{id}", wrap(controller.Delete))
	router.Put("/{id}", wrap(controller.Update))
	return router
}
