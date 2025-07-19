package api

import (
	"net/http"
)

type HealthCheckController struct{}

func (h *HealthCheckController) healthCheck(ctx *Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func NewHealthCheckController() *HealthCheckController {
	return &HealthCheckController{}
}
