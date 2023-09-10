package res

import "net/http"

func NotFound(w http.ResponseWriter, r *http.Request) error {
	return New(w).Status(http.StatusNotFound).JSON(ApiError{
		Message: "Not Found",
		Status:  http.StatusNotFound,
	})
}
