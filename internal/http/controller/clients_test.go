package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/go-faker/faker/v4"
	"github.com/raphael-foliveira/chi-gorm/internal/exceptions"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/mocks"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

func TestClient(t *testing.T) {

	controller := NewClients(service.NewClients(mocks.ClientsStore, mocks.OrdersStore))

	t.Run("List", func(t *testing.T) {
		t.Run("should list all clients", func(t *testing.T) {
			setUp()
			mocks.ClientsStore.Error = nil
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", "/", nil)
			err := controller.List(recorder, request)
			if err != nil {
				t.Fatal(err)
			}
			if recorder.Code != 200 {
				t.Errorf("Status code should be 200, got %v", recorder.Code)
			}
			tearDown()
		})

		t.Run("should return an error when store fails", func(t *testing.T) {
			setUp()
			mocks.ClientsStore.Error = errors.New("")
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", "/", nil)
			err := controller.List(recorder, request)
			if err == nil {
				t.Fatal("err should not be nil")
			}
			tearDown()
		})
	})

	t.Run("Get", func(t *testing.T) {
		t.Run("should get a client", func(t *testing.T) {
			setUp()
			mocks.ClientsStore.Error = nil
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", "/1", nil)
			tx := chi.NewRouteContext()
			tx.URLParams.Add("id", "1")
			request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
			err := controller.Get(recorder, request)
			if err != nil {
				t.Fatal(err)
			}
			if recorder.Code != 200 {
				t.Errorf("Status code should be 200, got %v", recorder.Code)
			}
			tearDown()
		})

		t.Run("should return an error when store fails", func(t *testing.T) {
			setUp()
			mocks.ClientsStore.Error = errors.New("")
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", "/99", nil)
			tx := chi.NewRouteContext()
			tx.URLParams.Add("id", "99")
			request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
			err := controller.Get(recorder, request)
			if err == nil {
				t.Fatal("err should not be nil")
			}
			tearDown()
		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("should create a client", func(t *testing.T) {
			setUp()
			mocks.ClientsStore.Error = nil
			recorder := httptest.NewRecorder()
			var newClient schemas.CreateClient
			faker.FakeData(&newClient)
			reqBody, _ := json.Marshal(newClient)
			request := httptest.NewRequest("POST", "/", bytes.NewReader(reqBody))
			err := controller.Create(recorder, request)
			if err != nil {
				t.Fatal(err)
			}
			if recorder.Code != 201 {
				t.Errorf("Status code should be 201, got %v", recorder.Code)
			}
			tearDown()
		})

		t.Run("should return an error when sent invalid data", func(t *testing.T) {
			setUp()
			invalidReqBody := `{"foo: 95}`
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("POST", "/", bytes.NewReader([]byte(invalidReqBody)))
			err := controller.Create(recorder, request)
			apiErr, ok := err.(*exceptions.ApiError)
			if !ok {
				t.Fatal("err should be an ApiError")
			}
			if apiErr.Status != 400 {
				t.Errorf("Status code should be 400, got %v", recorder.Code)
			}
			tearDown()
		})

		t.Run("should return an error when store fails", func(t *testing.T) {
			setUp()
			mocks.ClientsStore.Error = errors.New("")
			recorder := httptest.NewRecorder()
			var newClient schemas.CreateClient
			faker.FakeData(&newClient)
			reqBody, _ := json.Marshal(newClient)
			request := httptest.NewRequest("POST", "/", bytes.NewReader(reqBody))
			err := controller.Create(recorder, request)
			if err == nil {
				t.Fatal("err should not be nil")
			}
			tearDown()
		})
	})

	t.Run("Update", func(t *testing.T) {
		t.Run("should update a client", func(t *testing.T) {
			setUp()
			mocks.ClientsStore.Error = nil
			recorder := httptest.NewRecorder()
			var newClient schemas.UpdateClient
			faker.FakeData(&newClient)
			reqBody, _ := json.Marshal(newClient)
			request := httptest.NewRequest("PUT", "/1", bytes.NewReader(reqBody))
			tx := chi.NewRouteContext()
			tx.URLParams.Add("id", "1")
			request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
			err := controller.Update(recorder, request)
			if err != nil {
				t.Fatal(err)
			}
			if recorder.Code != 200 {
				t.Errorf("Status code should be 200, got %v", recorder.Code)
			}
			tearDown()
		})

		t.Run("should return an error when sent invalid data", func(t *testing.T) {
			setUp()
			invalidReqBody := `{"foo: 95}`
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("PUT", "/1", bytes.NewReader([]byte(invalidReqBody)))
			tx := chi.NewRouteContext()
			tx.URLParams.Add("id", "1")
			request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
			err := controller.Update(recorder, request)
			apiErr, ok := err.(*exceptions.ApiError)
			if !ok {
				t.Fatal("err should be an ApiError")
			}
			if apiErr.Status != 400 {
				t.Errorf("Status code should be 400, got %v", recorder.Code)
			}
			tearDown()
		})

		t.Run("should return an error when store fails", func(t *testing.T) {
			setUp()
			recorder := httptest.NewRecorder()
			var newClient schemas.UpdateClient
			faker.FakeData(&newClient)
			reqBody, _ := json.Marshal(newClient)
			request := httptest.NewRequest("PUT", "/99", bytes.NewReader(reqBody))
			tx := chi.NewRouteContext()
			tx.URLParams.Add("id", "99")
			request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
			err := controller.Update(recorder, request)
			if err == nil {
				t.Fatal("err should not be nil")
			}
			tearDown()
		})
	})

	t.Run("Delete", func(t *testing.T) {
		t.Run("should delete a client", func(t *testing.T) {
			setUp()
			mocks.ClientsStore.Error = nil
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("DELETE", "/1", nil)
			tx := chi.NewRouteContext()
			tx.URLParams.Add("id", "1")
			request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
			err := controller.Delete(recorder, request)
			if err != nil {
				t.Fatal(err)
			}
			if recorder.Code != 204 {
				t.Errorf("Status code should be 204, got %v", recorder.Code)
			}
			tearDown()
		})

		t.Run("should return an error when store fails", func(t *testing.T) {
			setUp()
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("DELETE", "/99", nil)
			tx := chi.NewRouteContext()
			tx.URLParams.Add("id", "99")
			request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
			err := controller.Delete(recorder, request)
			if err == nil {
				t.Fatal("err should not be nil")
			}
			tearDown()
		})

		t.Run("GET products", func(t *testing.T) {
			setUp()
			mocks.ClientsStore.Error = nil
			client := mocks.ClientsStore.Store[0]
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", fmt.Sprintf("/%v/products", client.ID), nil)
			tx := chi.NewRouteContext()
			tx.URLParams.Add("id", fmt.Sprintf("%v", client.GetId()))
			request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
			err := controller.GetProducts(recorder, request)
			if err != nil {
				t.Fatal(err)
			}
			tearDown()
		})
	})

}
