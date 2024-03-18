package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
)

func healthCheck(ctx *controller.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func HealthCheck() *Router {
	router := &Router{chi.NewRouter()}
	router.Get("/", healthCheck)
	return router
}
