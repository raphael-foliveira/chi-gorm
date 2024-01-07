package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/go-faker/faker/v4"
	"github.com/raphael-foliveira/chi-gorm/internal/exceptions"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/mocks"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

func TestProducts(t *testing.T) {

	controller := NewProducts(service.NewProducts(mocks.ProductsStore))

	t.Run("List", func(t *testing.T) {
		t.Run("should list all products", func(t *testing.T) {
			addProducts(10)
			mocks.ProductsStore.Error = nil
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", "/", nil)
			err := controller.List(recorder, request)
			if err != nil {
				t.Fatal(err)
			}
			if recorder.Code != 200 {
				t.Errorf("Status code should be 200, got %v", recorder.Code)
			}
		})

		t.Run("should return an error when store fails", func(t *testing.T) {
			mocks.ProductsStore.Error = errors.New("")
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", "/", nil)
			err := controller.List(recorder, request)
			if err == nil {
				t.Error("Should return an error")
			}
		})
	})

	t.Run("Get", func(t *testing.T) {
		t.Run("should get a product", func(t *testing.T) {
			addProducts(10)
			mocks.ProductsStore.Error = nil
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
		})

		t.Run("should return an error when store fails", func(t *testing.T) {
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", "/9999", nil)
			tx := chi.NewRouteContext()
			tx.URLParams.Add("id", "9999")
			request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
			err := controller.Get(recorder, request)
			if err == nil {
				t.Error("Should return an error")
			}
		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("should create a product", func(t *testing.T) {
			mocks.ProductsStore.Error = nil
			recorder := httptest.NewRecorder()
			var newProduct schemas.CreateProduct
			faker.FakeData(&newProduct)
			reqBody, _ := json.Marshal(newProduct)
			request := httptest.NewRequest("POST", "/", bytes.NewReader(reqBody))
			err := controller.Create(recorder, request)
			if err != nil {
				t.Fatal(err)
			}
			if recorder.Code != 201 {
				t.Errorf("Status code should be 201, got %v", recorder.Code)
			}
		})

		t.Run("should return an error when sent invalid data", func(t *testing.T) {
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
		})

		t.Run("should return an error when store fails", func(t *testing.T) {
			mocks.ProductsStore.Error = errors.New("")
			var newProduct schemas.CreateProduct
			faker.FakeData(&newProduct)
			reqBody, _ := json.Marshal(newProduct)
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("POST", "/", bytes.NewReader(reqBody))
			err := controller.Create(recorder, request)
			if err == nil {
				t.Error("Should return an error")
			}
		})
	})

	t.Run("Update", func(t *testing.T) {
		t.Run("should update a product", func(t *testing.T) {
			addProducts(10)
			mocks.ProductsStore.Error = nil
			recorder := httptest.NewRecorder()
			var newProduct schemas.UpdateProduct
			faker.FakeData(&newProduct)
			reqBody, _ := json.Marshal(newProduct)
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
		})

		t.Run("should return an error when sent invalid data", func(t *testing.T) {
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
		})

		t.Run("should return an error when store fails", func(t *testing.T) {
			recorder := httptest.NewRecorder()
			var newProduct schemas.UpdateProduct
			faker.FakeData(&newProduct)
			reqBody, _ := json.Marshal(newProduct)
			request := httptest.NewRequest("PUT", "/9999", bytes.NewReader(reqBody))
			tx := chi.NewRouteContext()
			tx.URLParams.Add("id", "9999")
			request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
			err := controller.Update(recorder, request)
			if err == nil {
				t.Error("Should return an error")
			}
		})
	})

	t.Run("Delete", func(t *testing.T) {
		t.Run("should delete a product", func(t *testing.T) {
			addProducts(10)
			mocks.ProductsStore.Error = nil
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
		})

		t.Run("should return an error when store fails", func(t *testing.T) {
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("DELETE", "/9999", nil)
			tx := chi.NewRouteContext()
			tx.URLParams.Add("id", "9999")
			request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
			err := controller.Delete(recorder, request)
			if err == nil {
				t.Error("Should return an error")
			}
		})
	})
}
