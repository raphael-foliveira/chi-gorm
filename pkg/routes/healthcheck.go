package routes

import (
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/pkg/res"
)

func healthCheck(w http.ResponseWriter, r *http.Request) error {
	return res.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func HealthCheckRoute() http.HandlerFunc {
	return wrap(healthCheck)
}
