package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controllers"
)

func Products() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", wrap(controllers.Products.List))
	router.Post("/", wrap(controllers.Products.Create))
	router.Get("/{id}", wrap(controllers.Products.Get))
	router.Delete("/{id}", wrap(controllers.Products.Delete))
	router.Put("/{id}", wrap(controllers.Products.Update))
	return router
}
