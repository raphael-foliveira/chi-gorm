package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type healthCheck struct{}

func (h *healthCheck) healthCheck(ctx *Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func NewHealthCheck() *healthCheck {
	return &healthCheck{}
}

func (h *healthCheck) Mount(mux *chi.Mux) {
	router := chi.NewRouter()
	router.Get("/", useHandler(h.healthCheck))

	mux.Mount("/health-check", router)
}
