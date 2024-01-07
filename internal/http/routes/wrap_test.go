package routes

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/exceptions"
)

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
		if body["message"] != "test" {
			t.Errorf("Body should be %v, got %v", "message", body["message"])
		}
	})

	t.Run("should handle notFoundErr when err is NotFoundError", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		handleApiErr(recorder, exceptions.NewNotFoundError("test not found"))
		if recorder.Code != 404 {
			t.Errorf("Status code should be 404, got %v", recorder.Code)
		}
		var body map[string]interface{}
		json.NewDecoder(recorder.Body).Decode(&body)
		if body["message"] != "test not found" {
			t.Errorf("Body should be %v, got %v", "test not found", body["message"])
		}
	})
}
