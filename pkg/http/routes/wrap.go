package routes

import (
	"fmt"
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/pkg/http/res"
)

func wrap(fn func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := fn(w, r)
		if err != nil {
			apiErr, ok := err.(res.ApiError)
			if ok {
				res.JSON(w, apiErr.Status, apiErr)
				return
			}
			fmt.Println("uncaught error:", err.Error())
			res.Error(w, err, http.StatusInternalServerError, "internal server error")
		}
	}
}
