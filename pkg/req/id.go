package req

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func ParseUrlIntParam(r *http.Request, name string) (int, error) {
	param := chi.URLParam(r, name)
	return strconv.Atoi(param)
}
