package routes

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
)

func TestHandleApiErr(t *testing.T) {
	t.Run("should handle apiErr when err is ApiError", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		handleApiErr(recorder, &controller.ApiError{
			Message: "test",
			Status:  400,
		})
		if recorder.Code != 400 {
			t.Errorf("Status code should be 400, got %v", recorder.Code)
		}
		var body map[string]interface{}
		json.NewDecoder(recorder.Body).Decode(&body)
		fmt.Println(body)
		message := body["message"]
		if message != "test" {
			t.Errorf("Body should be %v, got %v", "message", message)
		}
	})

	t.Run("should handle notFoundErr when err is NotFoundError", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		handleApiErr(recorder, controller.NotFound("test not found"))
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