package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
)

func Users() *chi.Mux {
	c := controller.Users()
	router := chi.NewRouter()
	router.Post("/register", useHandler(c.Register))
	return router
}
