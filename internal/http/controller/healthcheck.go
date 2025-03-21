package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type HealthCheck struct{}

func (h *HealthCheck) healthCheck(ctx *Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func (c *HealthCheck) Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", useHandler(c.healthCheck))

	return router
}

func NewHealthCheck() *HealthCheck {
	return &HealthCheck{}
}

func (h *HealthCheck) Mount(mux *chi.Mux) {
	router := chi.NewRouter()
	router.Get("/", useHandler(h.healthCheck))

	mux.Mount("/health-check", router)
}
