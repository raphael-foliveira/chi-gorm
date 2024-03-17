package controller_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
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

func TestProducts(t *testing.T) {
	productsController := controllers.ProductsController

	t.Run("List", func(t *testing.T) {
		t.Run("should list all products", testCase(func(t *testing.T) {
			mocks.ProductsRepository.Error = nil
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", "/", nil)
			ctx := controller.NewContext(recorder, request)
			err := productsController.List(ctx)
			if err != nil {
				t.Fatal(err)
			}
			if recorder.Code != 200 {
				t.Errorf("Status code should be 200, got %v", recorder.Code)
			}
		}))

		t.Run("should return an error when store fails", testCase(func(t *testing.T) {
			mocks.ProductsRepository.Error = errors.New("")
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", "/", nil)
			ctx := controller.NewContext(recorder, request)
			err := productsController.List(ctx)
			if err == nil {
				t.Error("Should return an error")
			}
		}))
	})

	t.Run("Get", func(t *testing.T) {
		t.Run("should get a product", testCase(func(t *testing.T) {
			mocks.ProductsRepository.Error = nil
			recorder := httptest.NewRecorder()
			productId := fmt.Sprintf("%v", mocks.ProductsRepository.Store[0].ID)
			request := httptest.NewRequest("GET", "/"+productId, nil)
			tx := chi.NewRouteContext()
			tx.URLParams.Add("id", productId)
			request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
			ctx := controller.NewContext(recorder, request)
			err := productsController.Get(ctx)
			if err != nil {
				t.Fatal(err)
			}
			if recorder.Code != 200 {
				t.Errorf("Status code should be 200, got %v", recorder.Code)
			}
		}))

		t.Run("should return an error when store fails", testCase(func(t *testing.T) {
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", "/9999", nil)
			tx := chi.NewRouteContext()
			tx.URLParams.Add("id", "9999")
			request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
			ctx := controller.NewContext(recorder, request)
			err := productsController.Get(ctx)
			if err == nil {
				t.Error("Should return an error")
			}
		}))
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("should create a product", testCase(func(t *testing.T) {
			mocks.ProductsRepository.Error = nil
			recorder := httptest.NewRecorder()
			var newProduct schemas.CreateProduct
			faker.FakeData(&newProduct)
			reqBody, _ := json.Marshal(newProduct)
			request := httptest.NewRequest("POST", "/", bytes.NewReader(reqBody))
			ctx := controller.NewContext(recorder, request)
			err := productsController.Create(ctx)
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
			err := productsController.Create(ctx)
			apiErr, ok := err.(*exceptions.ApiError)
			if !ok {
				t.Fatal("err should be an ApiError")
			}
			if apiErr.Status != http.StatusUnprocessableEntity {
				t.Errorf("Status code should be 422, got %v", recorder.Code)
			}
		}))

		t.Run("should return an error when store fails", testCase(func(t *testing.T) {
			mocks.ProductsRepository.Error = errors.New("")
			var newProduct schemas.CreateProduct
			faker.FakeData(&newProduct)
			reqBody, _ := json.Marshal(newProduct)
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("POST", "/", bytes.NewReader(reqBody))
			ctx := controller.NewContext(recorder, request)
			err := productsController.Create(ctx)
			if err == nil {
				t.Error("Should return an error")
			}
		}))
	})

	t.Run("Update", func(t *testing.T) {
		t.Run("should update a product", testCase(func(t *testing.T) {
			mocks.ProductsRepository.Error = nil
			recorder := httptest.NewRecorder()
			product := mocks.ProductsRepository.Store[0]
			productId := fmt.Sprintf("%v", product.ID)
			reqBody, _ := json.Marshal(product)
			request := httptest.NewRequest("PUT", "/"+productId, bytes.NewReader(reqBody))
			tx := chi.NewRouteContext()
			tx.URLParams.Add("id", productId)
			request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
			ctx := controller.NewContext(recorder, request)
			err := productsController.Update(ctx)
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
			err := productsController.Update(ctx)
			apiErr, ok := err.(*exceptions.ApiError)
			if !ok {
				t.Fatal("err should be an ApiError")
			}
			if apiErr.Status != http.StatusUnprocessableEntity {
				t.Errorf("Status code should be 422, got %v", recorder.Code)
			}
		}))

		t.Run("should return an error when store fails", testCase(func(t *testing.T) {
			recorder := httptest.NewRecorder()
			var newProduct schemas.UpdateProduct
			faker.FakeData(&newProduct)
			reqBody, _ := json.Marshal(newProduct)
			request := httptest.NewRequest("PUT", "/9999", bytes.NewReader(reqBody))
			tx := chi.NewRouteContext()
			tx.URLParams.Add("id", "9999")
			request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
			ctx := controller.NewContext(recorder, request)
			err := productsController.Update(ctx)
			if err == nil {
				t.Error("Should return an error")
			}
		}))
	})

	t.Run("Delete", func(t *testing.T) {
		t.Run("should delete a product", testCase(func(t *testing.T) {
			mocks.ProductsRepository.Error = nil
			product := mocks.ProductsRepository.Store[0]
			productId := fmt.Sprintf("%v", product.ID)
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("DELETE", "/"+productId, nil)
			tx := chi.NewRouteContext()
			tx.URLParams.Add("id", productId)
			request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
			ctx := controller.NewContext(recorder, request)
			err := productsController.Delete(ctx)
			if err != nil {
				t.Fatal(err)
			}
			if recorder.Code != 204 {
				t.Errorf("Status code should be 204, got %v", recorder.Code)
			}
		}))

		t.Run("should return an error when store fails", testCase(func(t *testing.T) {
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("DELETE", "/9999", nil)
			tx := chi.NewRouteContext()
			tx.URLParams.Add("id", "9999")
			request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
			ctx := controller.NewContext(recorder, request)
			err := productsController.Delete(ctx)
			if err == nil {
				t.Error("Should return an error")
			}
		}))
	})
}
