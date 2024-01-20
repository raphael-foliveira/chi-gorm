package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
)

func Products(c *controller.Products) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", useHandler(c.List))
	router.Post("/", useHandler(c.Create))
	router.Get("/{id}", useHandler(c.Get))
	router.Delete("/{id}", useHandler(c.Delete))
	router.Put("/{id}", useHandler(c.Update))
	return router
}
