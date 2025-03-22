package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type HealthCheck struct{}

func (h *HealthCheck) healthCheck(ctx *Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func NewHealthCheck() *HealthCheck {
	return &HealthCheck{}
}

func (h *HealthCheck) Mount(mux chi.Router) {
	router := NewRouter()
	router.Get("/", h.healthCheck)

	mux.Mount("/health-check", router)
}
