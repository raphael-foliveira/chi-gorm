package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type HealthCheck struct{}

func healthCheck(ctx *Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func NewHealthCheck() *HealthCheck {
	return &HealthCheck{}
}

func (h *HealthCheck) Mount(mux *chi.Mux) {
	router := chi.NewRouter()
	router.Get("/", useHandler(healthCheck))

	mux.Mount("/health-check", router)
}