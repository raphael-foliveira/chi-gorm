package controller_test

import (
	"bytes"
	"context"
	"encoding/json"
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
	"github.com/stretchr/testify/assert"
)

var productsController *controller.Products

func TestProducts_List(t *testing.T) {
	t.Run("should list all products", testCase(func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/", nil)
		ctx := controller.NewContext(recorder, request)
		err := productsController.List(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, recorder.Code)
	}))

	t.Run("should return an error when store fails", testCase(func(t *testing.T) {
		mocks.ProductsRepository.ExpectError()
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/", nil)
		ctx := controller.NewContext(recorder, request)
		err := productsController.List(ctx)
		assert.Error(t, err)
	}))
}

func TestProducts_Get(t *testing.T) {
	t.Run("should get a product", testCase(func(t *testing.T) {
		recorder := httptest.NewRecorder()
		productId := fmt.Sprintf("%v", mocks.ProductsStub[0].ID)
		request := httptest.NewRequest("GET", "/"+productId, nil)
		tx := chi.NewRouteContext()
		tx.URLParams.Add("id", productId)
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		ctx := controller.NewContext(recorder, request)
		err := productsController.Get(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, recorder.Code)
	}))

	t.Run("should return an error when store fails", testCase(func(t *testing.T) {
		mocks.ProductsRepository.ExpectError()
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/9999", nil)
		tx := chi.NewRouteContext()
		tx.URLParams.Add("id", "9999")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		ctx := controller.NewContext(recorder, request)
		err := productsController.Get(ctx)
		assert.Error(t, err)
	}))
}

func TestProducts_Create(t *testing.T) {
	t.Run("should create a product", testCase(func(t *testing.T) {
		recorder := httptest.NewRecorder()
		var newProduct schemas.CreateProduct
		faker.FakeData(&newProduct)
		reqBody, _ := json.Marshal(newProduct)
		request := httptest.NewRequest("POST", "/", bytes.NewReader(reqBody))
		ctx := controller.NewContext(recorder, request)
		err := productsController.Create(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, recorder.Code)
	}))

	t.Run("should return an error when sent invalid data", testCase(func(t *testing.T) {
		invalidReqBody := `{"foo: 95}`
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("POST", "/", bytes.NewReader([]byte(invalidReqBody)))
		ctx := controller.NewContext(recorder, request)
		err := productsController.Create(ctx)
		apiErr, ok := err.(*exceptions.ApiError)
		assert.True(t, ok, "err should be an ApiError")
		assert.Equal(t, http.StatusUnprocessableEntity, apiErr.Status)
	}))

	t.Run("should return an error when store fails", testCase(func(t *testing.T) {
		mocks.ProductsRepository.ExpectError()
		var newProduct schemas.CreateProduct
		faker.FakeData(&newProduct)
		reqBody, _ := json.Marshal(newProduct)
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("POST", "/", bytes.NewReader(reqBody))
		ctx := controller.NewContext(recorder, request)
		err := productsController.Create(ctx)
		assert.Error(t, err)
	}))
}

func TestProducts_Update(t *testing.T) {
	t.Run("should update a product", testCase(func(t *testing.T) {
		recorder := httptest.NewRecorder()
		product := mocks.ProductsStub[0]
		productId := fmt.Sprintf("%v", product.ID)
		reqBody, _ := json.Marshal(product)
		request := httptest.NewRequest("PUT", "/"+productId, bytes.NewReader(reqBody))
		tx := chi.NewRouteContext()
		tx.URLParams.Add("id", productId)
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		ctx := controller.NewContext(recorder, request)
		err := productsController.Update(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, recorder.Code)
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
		assert.True(t, ok, "err should be an ApiError")
		assert.Equal(t, http.StatusUnprocessableEntity, apiErr.Status)
	}))

	t.Run("should return an error when store fails", testCase(func(t *testing.T) {
		mocks.ProductsRepository.ExpectError()
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
		assert.Error(t, err)
	}))
}

func TestProducts_Delete(t *testing.T) {
	t.Run("should delete a product", testCase(func(t *testing.T) {
		product := mocks.ProductsStub[0]
		productId := fmt.Sprintf("%v", product.ID)
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("DELETE", "/"+productId, nil)
		tx := chi.NewRouteContext()
		tx.URLParams.Add("id", productId)
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		ctx := controller.NewContext(recorder, request)
		err := productsController.Delete(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, recorder.Code)
	}))

	t.Run("should return an error when store fails", testCase(func(t *testing.T) {
		mocks.ProductsRepository.ExpectError()
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("DELETE", "/9999", nil)
		tx := chi.NewRouteContext()
		tx.URLParams.Add("id", "9999")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		ctx := controller.NewContext(recorder, request)
		err := productsController.Delete(ctx)
		assert.Error(t, err)
	}))

}
