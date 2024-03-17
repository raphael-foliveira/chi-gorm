package controller_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/go-faker/faker/v4"
	"github.com/raphael-foliveira/chi-gorm/internal/exceptions"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/mocks"
)

var clientsController *controller.Clients

func TestClient_List(t *testing.T) {
	t.Run("should list all clients", testCase(func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/", nil)
		ctx := controller.NewContext(recorder, request)
		err := clientsController.List(ctx)
		if err != nil {
			t.Fatal(err)
		}
		if recorder.Code != 200 {
			t.Errorf("Status code should be 200, got %v", recorder.Code)
		}
	}))

	t.Run("should return an error when store fails", testCase(func(t *testing.T) {
		mocks.ClientsRepository.ExpectError()
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/", nil)
		ctx := controller.NewContext(recorder, request)
		err := clientsController.List(ctx)
		if err == nil {
			t.Fatal("err should not be nil")
		}
	}))
}

func TestClient_Get(t *testing.T) {
	t.Run("should get a client", testCase(func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/1", nil)
		tx := chi.NewRouteContext()
		tx.URLParams.Add("id", "1")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		ctx := controller.NewContext(recorder, request)
		err := clientsController.Get(ctx)
		if err != nil {
			t.Fatal(err)
		}
		if recorder.Code != 200 {
			t.Errorf("Status code should be 200, got %v", recorder.Code)
		}
	}))

	t.Run("should return an error when store fails", testCase(func(t *testing.T) {
		mocks.ClientsRepository.ExpectError()
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/99", nil)
		tx := chi.NewRouteContext()
		tx.URLParams.Add("id", "99")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		ctx := controller.NewContext(recorder, request)
		err := clientsController.Get(ctx)
		if err == nil {
			t.Fatal("err should not be nil")
		}
	}))
}

func TestClient_Create(t *testing.T) {
	t.Run("should create a client", testCase(func(t *testing.T) {
		recorder := httptest.NewRecorder()
		var newClient schemas.CreateClient
		faker.FakeData(&newClient)
		reqBody, _ := json.Marshal(newClient)
		request := httptest.NewRequest("POST", "/", bytes.NewReader(reqBody))
		ctx := controller.NewContext(recorder, request)
		err := clientsController.Create(ctx)
		if err != nil {
			t.Fatal(err)
		}
		if recorder.Code != 201 {
			t.Errorf("Status code should be 201, got %v", recorder.Code)
		}
	}))

	t.Run("should return an error when sent invalid data", testCase(func(t *testing.T) {
		invalidReqBody := `{"foo: 95}`
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("POST", "/", bytes.NewReader([]byte(invalidReqBody)))
		ctx := controller.NewContext(recorder, request)
		err := clientsController.Create(ctx)
		apiErr, ok := err.(*exceptions.ApiError)
		if !ok {
			t.Fatal("err should be an ApiError")
		}
		if apiErr.Status != http.StatusUnprocessableEntity {
			t.Errorf("Status code should be 422, got %v", recorder.Code)
		}
	}))

	t.Run("should return an error when store fails", testCase(func(t *testing.T) {
		mocks.ClientsRepository.ExpectError()
		recorder := httptest.NewRecorder()
		var newClient schemas.CreateClient
		faker.FakeData(&newClient)
		reqBody, _ := json.Marshal(newClient)
		request := httptest.NewRequest("POST", "/", bytes.NewReader(reqBody))
		ctx := controller.NewContext(recorder, request)
		err := clientsController.Create(ctx)
		if err == nil {
			t.Fatal("err should not be nil")
		}
	}))
}

func TestClient_Update(t *testing.T) {
	t.Run("should update a client", testCase(func(t *testing.T) {
		recorder := httptest.NewRecorder()
		var newClient schemas.UpdateClient
		faker.FakeData(&newClient)
		reqBody, _ := json.Marshal(newClient)
		request := httptest.NewRequest("PUT", "/1", bytes.NewReader(reqBody))
		tx := chi.NewRouteContext()
		tx.URLParams.Add("id", "1")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		ctx := controller.NewContext(recorder, request)
		err := clientsController.Update(ctx)
		if err != nil {
			t.Fatal(err)
		}
		if recorder.Code != 200 {
			t.Errorf("Status code should be 200, got %v", recorder.Code)
		}
	}))

	t.Run("should return an error when sent invalid data", testCase(func(t *testing.T) {
		invalidReqBody := `{"foo: 95}`
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("PUT", "/1", bytes.NewReader([]byte(invalidReqBody)))
		tx := chi.NewRouteContext()
		tx.URLParams.Add("id", "1")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		ctx := controller.NewContext(recorder, request)
		err := clientsController.Update(ctx)
		apiErr, ok := err.(*exceptions.ApiError)
		if !ok {
			t.Fatal("err should be an ApiError")
		}
		if apiErr.Status != 422 {
			t.Errorf("Status code should be 422, got %v", apiErr.Status)
		}
	}))

	t.Run("should return an error when store fails", testCase(func(t *testing.T) {
		mocks.ClientsRepository.ExpectError()
		recorder := httptest.NewRecorder()
		var newClient schemas.UpdateClient
		faker.FakeData(&newClient)
		reqBody, _ := json.Marshal(newClient)
		request := httptest.NewRequest("PUT", "/99", bytes.NewReader(reqBody))
		tx := chi.NewRouteContext()
		tx.URLParams.Add("id", "99")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		ctx := controller.NewContext(recorder, request)
		err := clientsController.Update(ctx)
		if err == nil {
			t.Fatal("err should not be nil")
		}
	}))
}

func TestClient_Delete(t *testing.T) {
	t.Run("should delete a client", testCase(func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("DELETE", "/1", nil)
		tx := chi.NewRouteContext()
		tx.URLParams.Add("id", "1")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		ctx := controller.NewContext(recorder, request)
		err := clientsController.Delete(ctx)
		if err != nil {
			t.Fatal(err)
		}
		if recorder.Code != 204 {
			t.Errorf("Status code should be 204, got %v", recorder.Code)
		}
	}))

	t.Run("should return an error when store fails", testCase(func(t *testing.T) {
		mocks.ClientsRepository.ExpectError()
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("DELETE", "/99", nil)
		tx := chi.NewRouteContext()
		tx.URLParams.Add("id", "99")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		ctx := controller.NewContext(recorder, request)
		err := clientsController.Delete(ctx)
		if err == nil {
			t.Fatal("err should not be nil")
		}
	}))
}
