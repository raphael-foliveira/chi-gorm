package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJson(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	Json(handler).ServeHTTP(recorder, request)
	if recorder.Header().Get("Content-Type") != "application/json" {
		t.Error("Content-Type header not set")
	}
}
