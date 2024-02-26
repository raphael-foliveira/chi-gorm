package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
)

func Auth() *chi.Mux {
	c := controller.Auth()
	router := chi.NewRouter()
	router.Post("/login", useHandler(c.Login))
	return router
}
