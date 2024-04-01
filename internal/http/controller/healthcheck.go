package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type HealthcheckController struct {
	*router
}

func healthCheck(ctx *Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func NewHealthCheck() *HealthcheckController {
	router := &router{chi.NewRouter()}
	c := &HealthcheckController{router}
	router.Get("/", healthCheck)
	return c
}
