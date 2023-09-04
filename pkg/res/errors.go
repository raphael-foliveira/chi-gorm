package res

import (
	"errors"
	"fmt"
	"net/http"
)

func Error(w http.ResponseWriter, status int, message string, err error) error {
	fmt.Println(err)
	w.WriteHeader(status)
	return errors.New(message)
}
