package controller

import (
	"net/http"
)

type HealthCheck struct{}

func (h *HealthCheck) healthCheck(ctx *Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func NewHealthCheck() *HealthCheck {
	return &HealthCheck{}
}

