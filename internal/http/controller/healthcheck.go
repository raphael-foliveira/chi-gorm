package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func healthCheck(ctx *Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func NewHealthCheck() *Router {
	router := &Router{chi.NewRouter()}
	router.Get("/", healthCheck)
	return router
}
