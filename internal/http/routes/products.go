package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
)

func Products(c *controller.Products) *Router {
	router := &Router{chi.NewRouter()}
	router.Get("/", c.List)
	router.Post("/", c.Create)
	router.Get("/{id}", c.Get)
	router.Delete("/{id}", c.Delete)
	router.Put("/{id}", c.Update)
	return router
}
