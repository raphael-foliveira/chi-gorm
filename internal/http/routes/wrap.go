package routes

import (
	"fmt"
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/internal/http/res"
)

func wrap(fn func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleApiErr(w, fn(w, r))
	}
}

func handleApiErr(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}
	fmt.Println(err.Error())
	apiErr, ok := err.(res.ApiError)
	if ok {
		res.JSON(w, apiErr.Status, apiErr)
		return
	}
	res.JSON(w, http.StatusInternalServerError, res.ApiError{
		Message: "internal server error",
		Status:  http.StatusInternalServerError,
	})
}
