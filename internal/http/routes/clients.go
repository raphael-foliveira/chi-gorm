package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controllers"
)

func Clients() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", wrap(controllers.Clients.List))
	router.Post("/", wrap(controllers.Clients.Create))
	router.Get("/{id}", wrap(controllers.Clients.Get))
	router.Delete("/{id}", wrap(controllers.Clients.Delete))
	router.Put("/{id}", wrap(controllers.Clients.Update))
	return router
}
