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

func TestOrders_List(t *testing.T) {
	t.Run("should list all orders", testCase(func(t *testing.T, deps *testDependencies) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/", nil)
		ctx := controller.NewContext(recorder, request)
		err := deps.ordersController.List(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, recorder.Code)
	}))

	t.Run("should return an error when store fails", testCase(func(t *testing.T, deps *testDependencies) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/", nil)
		ctx := controller.NewContext(recorder, request)
		err := deps.ordersController.List(ctx)
		assert.Error(t, err)
	}))
}

func TestOrders_Get(t *testing.T) {
	t.Run("should get an order", testCase(func(t *testing.T, deps *testDependencies) {
		orderId := fmt.Sprintf("%v", mocks.OrdersStub[0].ID)
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/"+orderId, nil)
		tx := chi.NewRouteContext()
		tx.URLParams.Add("id", orderId)
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		ctx := controller.NewContext(recorder, request)
		err := deps.ordersController.Get(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, recorder.Code)
		var requestBody *schemas.Order
		json.NewDecoder(recorder.Body).Decode(&requestBody)
		assert.Equal(t, mocks.OrdersStub[0].ID, requestBody.ID)
	}))

	t.Run("should return an error when store fails", testCase(func(t *testing.T, deps *testDependencies) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/9999", nil)
		tx := chi.NewRouteContext()
		tx.URLParams.Add("id", "9999")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		ctx := controller.NewContext(recorder, request)
		err := deps.ordersController.Get(ctx)
		assert.Error(t, err)
	}))
}

func TestOrders_Create(t *testing.T) {
	t.Run("should create an order", testCase(func(t *testing.T, deps *testDependencies) {
		recorder := httptest.NewRecorder()
		var newOrder schemas.CreateOrder
		faker.FakeData(&newOrder)
		reqBody, _ := json.Marshal(newOrder)
		request := httptest.NewRequest("POST", "/", bytes.NewReader(reqBody))
		ctx := controller.NewContext(recorder, request)
		err := deps.ordersController.Create(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, recorder.Code)
	}))

	t.Run("should return an error when sent invalid data", testCase(func(t *testing.T, deps *testDependencies) {
		invalidReqBody := `{"foo: 95}`
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("POST", "/", bytes.NewReader([]byte(invalidReqBody)))
		ctx := controller.NewContext(recorder, request)
		err := deps.ordersController.Create(ctx)
		apiErr, ok := err.(*exceptions.ApiError)
		assert.True(t, ok, "err should be an ApiError")
		assert.Equal(t, http.StatusUnprocessableEntity, apiErr.Status)
	}))

	t.Run("should return an error when store fails", testCase(func(t *testing.T, deps *testDependencies) {
		recorder := httptest.NewRecorder()
		var newOrder schemas.CreateOrder
		faker.FakeData(&newOrder)
		reqBody, _ := json.Marshal(newOrder)
		request := httptest.NewRequest("POST", "/", bytes.NewReader(reqBody))
		ctx := controller.NewContext(recorder, request)
		err := deps.ordersController.Create(ctx)
		assert.Error(t, err)
	}))
}

func TestOrders_Update(t *testing.T) {
	t.Run("should update an order", testCase(func(t *testing.T, deps *testDependencies) {
		recorder := httptest.NewRecorder()
		order := mocks.OrdersStub[0]
		reqBody, _ := json.Marshal(order)
		request := httptest.NewRequest("PUT", fmt.Sprintf("/%v", order.ID), bytes.NewReader(reqBody))
		tx := chi.NewRouteContext()
		tx.URLParams.Add("id", fmt.Sprintf("%v", order.ID))
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		ctx := controller.NewContext(recorder, request)
		err := deps.ordersController.Update(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, recorder.Code)
	}))

	t.Run("should return an error when sent invalid data", testCase(func(t *testing.T, deps *testDependencies) {
		invalidReqBody := `{"foo: 95}`
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("PUT", "/1", bytes.NewReader([]byte(invalidReqBody)))
		tx := chi.NewRouteContext()
		tx.URLParams.Add("id", "1")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		ctx := controller.NewContext(recorder, request)
		err := deps.ordersController.Update(ctx)
		apiErr, ok := err.(*exceptions.ApiError)
		assert.True(t, ok, "err should be an ApiError")
		assert.Equal(t, http.StatusUnprocessableEntity, apiErr.Status)
	}))

	t.Run("should return an error when store fails", testCase(func(t *testing.T, deps *testDependencies) {
		recorder := httptest.NewRecorder()
		var newOrder schemas.UpdateOrder
		faker.FakeData(&newOrder)
		reqBody, _ := json.Marshal(newOrder)
		request := httptest.NewRequest("PUT", "/9999", bytes.NewReader(reqBody))
		tx := chi.NewRouteContext()
		tx.URLParams.Add("id", "9999")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		ctx := controller.NewContext(recorder, request)
		err := deps.ordersController.Update(ctx)
		assert.Error(t, err)
	}))
}

func TestOrders_Delete(t *testing.T) {
	t.Run("should delete an order", testCase(func(t *testing.T, deps *testDependencies) {
		order := mocks.OrdersStub[0]
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("DELETE", fmt.Sprintf("/%v", order.ID), nil)
		tx := chi.NewRouteContext()
		tx.URLParams.Add("id", fmt.Sprintf("%v", order.ID))
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		ctx := controller.NewContext(recorder, request)
		err := deps.ordersController.Delete(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, recorder.Code)
	}))

	t.Run("should return an error when store fails", testCase(func(t *testing.T, deps *testDependencies) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("DELETE", "/9999", nil)
		tx := chi.NewRouteContext()
		tx.URLParams.Add("id", "9999")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		ctx := controller.NewContext(recorder, request)
		err := deps.ordersController.Delete(ctx)
		assert.Error(t, err)
	}))
}
