package routes

import (
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/internal/http/res"
)

func healthCheck(w http.ResponseWriter, r *http.Request) error {
	return res.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func HealthCheck() http.HandlerFunc {
	return useHandler(healthCheck)
}
