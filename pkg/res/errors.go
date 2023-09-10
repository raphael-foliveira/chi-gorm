package res

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Error(w http.ResponseWriter, status int, message string, err error) error {
	fmt.Println(err)
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(ApiError{
		Message: message,
		Status:  status,
	})
}
