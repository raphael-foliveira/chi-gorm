package controller

import (
	"net/http/httptest"
	"testing"
)

func TestUtils(t *testing.T) {
	t.Run("Should return an error when given an invalid id", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/foo", nil)
		_, err := getIdFromPath(request)
		if err == nil {
			t.Error("Should return an error")
		}
	})

	t.Run("Should return an error when given an invalid body", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/foo", nil)
		_, err := parseBody(request, &struct{}{})
		if err == nil {
			t.Error("Should return an error")
		}
	})
}
