package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/exceptions"
)

func TestWrap(t *testing.T) {
	t.Run("should handle an uncaught error", func(t *testing.T) {
		cf := func(w http.ResponseWriter, r *http.Request) error {
			return errors.New("uncaught error")
		}
		router := chi.NewRouter()
		router.Get("/", useHandler(cf))
		recorder := httptest.NewRecorder()
		request, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Error(err)
		}
		router.ServeHTTP(recorder, request)
		if recorder.Code != http.StatusInternalServerError {
			t.Errorf("expected status code %d, got %d", http.StatusInternalServerError, recorder.Code)
		}
	})
}

func TestHealthCheck(t *testing.T) {
	router := chi.NewRouter()
	router.Mount("/", HealthCheck())
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Error(err)
	}
	router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, recorder.Code)
	}
}

func TestHandleApiErr(t *testing.T) {
	t.Run("should handle apiErr when err is ApiError", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		handleApiErr(recorder, &exceptions.ApiError{
			Message: "test",
			Status:  400,
		})
		if recorder.Code != 400 {
			t.Errorf("Status code should be 400, got %v", recorder.Code)
		}
		var body map[string]interface{}
		json.NewDecoder(recorder.Body).Decode(&body)
		fmt.Println(body)
		message := body["error"]
		if message != "test" {
			t.Errorf("Body should be %v, got %v", "test", message)
		}
	})

	t.Run("should handle notFoundErr when err is NotFoundError", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		handleApiErr(recorder, exceptions.NotFound("test not found"))
		if recorder.Code != 404 {
			t.Errorf("Status code should be 404, got %v", recorder.Code)
		}
		var body map[string]interface{}
		json.NewDecoder(recorder.Body).Decode(&body)
		message := body["error"]
		if message != "test not found" {
			t.Errorf("Body should be %v, got %v", "test not found", message)
		}
	})
}
