package api_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/go-faker/faker/v4"
	"github.com/raphael-foliveira/chi-gorm/internal/http/api"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestOrders_List(t *testing.T) {
	t.Run("should list all orders", testCase(func(t *testing.T, deps *testDependencies) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/", nil)
		ctx := api.NewContext(recorder, request)
		err := deps.ordersController.List(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, recorder.Code)
	}))
}

func TestOrders_Get(t *testing.T) {
	t.Run("should get an order", testCase(func(t *testing.T, deps *testDependencies) {
		orderId := fmt.Sprintf("%v", deps.ordersStubs[0].ID)
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/"+orderId, nil)
		tx := chi.NewRouteContext()
		tx.URLParams.Add("id", orderId)
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		ctx := api.NewContext(recorder, request)
		err := deps.ordersController.Get(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, recorder.Code)
		var requestBody *schemas.Order
		json.NewDecoder(recorder.Body).Decode(&requestBody)
		assert.Equal(t, deps.ordersStubs[0].ID, requestBody.ID)
	}))
}

func TestOrders_Create(t *testing.T) {
	t.Run("should create an order", testCase(func(t *testing.T, deps *testDependencies) {
		recorder := httptest.NewRecorder()
		var newOrder schemas.CreateOrder
		faker.FakeData(&newOrder)
		newOrder.ClientID = deps.clientsStubs[0].ID
		newOrder.ProductID = deps.productsStubs[0].ID
		reqBody, _ := json.Marshal(newOrder)
		request := httptest.NewRequest("POST", "/", bytes.NewReader(reqBody))
		ctx := api.NewContext(recorder, request)
		err := deps.ordersController.Create(ctx)
		var responseBody map[string]any
		json.NewDecoder(recorder.Body).Decode(&responseBody)
		log.Printf("responseBody: %#v", responseBody)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, recorder.Code)
	}))

	t.Run("should return an error when sent invalid data", testCase(func(t *testing.T, deps *testDependencies) {
		invalidReqBody := `{"foo: 95}`
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("POST", "/", bytes.NewReader([]byte(invalidReqBody)))
		ctx := api.NewContext(recorder, request)
		err := deps.ordersController.Create(ctx)
		apiErr, ok := err.(*api.ApiError)
		assert.True(t, ok, "err should be an ApiError")
		assert.Equal(t, http.StatusUnprocessableEntity, apiErr.Status)
	}))
}

func TestOrders_Update(t *testing.T) {
	t.Run("should update an order", testCase(func(t *testing.T, deps *testDependencies) {
		recorder := httptest.NewRecorder()
		order := deps.ordersStubs[0]
		reqBody, _ := json.Marshal(&schemas.CreateOrder{
			Quantity:  10,
			ClientID:  deps.clientsStubs[0].ID,
			ProductID: deps.productsStubs[0].ID,
		})
		request := httptest.NewRequest("PUT", fmt.Sprintf("/%v", order.ID), bytes.NewReader(reqBody))
		tx := chi.NewRouteContext()
		tx.URLParams.Add("id", fmt.Sprintf("%v", order.ID))
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		ctx := api.NewContext(recorder, request)
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
		ctx := api.NewContext(recorder, request)
		err := deps.ordersController.Update(ctx)
		apiErr, ok := err.(*api.ApiError)
		assert.True(t, ok, "err should be an ApiError")
		assert.Equal(t, http.StatusUnprocessableEntity, apiErr.Status)
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
		ctx := api.NewContext(recorder, request)
		err := deps.ordersController.Delete(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, recorder.Code)
	}))
}
