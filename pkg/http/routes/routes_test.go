package routes

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestWrap(t *testing.T) {

	t.Run("should handle an uncaught error", func(t *testing.T) {
		cf := func(w http.ResponseWriter, r *http.Request) error {
			return errors.New("uncaught error")
		}
		router := chi.NewRouter()
		router.Get("/", wrap(cf))
		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		router.ServeHTTP(recorder, request)
		if recorder.Code != http.StatusInternalServerError {
			t.Errorf("expected status code %d, got %d", http.StatusInternalServerError, recorder.Code)
		}
	})
}

func TestHealthCheck(t *testing.T) {
	router := chi.NewRouter()
	router.Get("/", HealthCheckRoute())
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, recorder.Code)
	}
}
