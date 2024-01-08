package controller

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
)

func TestUtils(t *testing.T) {
	t.Run("Should return an error when given an invalid id", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/foo", nil)
		_, err := getUintPathParam(request, "id")
		if err == nil {
			t.Error("Should return an error")
		}
	})

	t.Run("Should return an error when given an invalid body", func(t *testing.T) {
		b := bytes.NewBuffer([]byte(`{"message": "some invalid json", "name": ""}`))
		request := httptest.NewRequest("GET", "/foo", b)
		schema := schemas.UpdateClient{}
		_, err := parseBody(request, &schema)
		if err == nil {
			t.Error("Should return an error")
		}
	})
}
