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
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/mocks"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

func TestOrders(t *testing.T) {

	controller := NewOrders(service.NewOrders(mocks.OrdersStore))
	t.Run("List", func(t *testing.T) {
		t.Run("should list all orders", func(t *testing.T) {
			setUp()
			mocks.OrdersStore.Error = nil
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
			mocks.OrdersStore.Error = errors.New("")
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", "/", nil)
			err := controller.List(recorder, request)
			if err == nil {
				t.Fatal("err should not be nil")
			}
		})
	})

	t.Run("Get", func(t *testing.T) {
		t.Run("should get an order", func(t *testing.T) {
			setUp()
			mocks.OrdersStore.Error = nil
			orderId := fmt.Sprintf("%v", mocks.OrdersStore.Store[0].ID)
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", "/"+orderId, nil)
			tx := chi.NewRouteContext()
			tx.URLParams.Add("id", orderId)
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
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", "/9999", nil)
			tx := chi.NewRouteContext()
			tx.URLParams.Add("id", "9999")
			request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
			err := controller.Get(recorder, request)
			if err == nil {
				t.Fatal("err should not be nil")
			}
		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("should create an order", func(t *testing.T) {
			mocks.OrdersStore.Error = nil
			recorder := httptest.NewRecorder()
			var newOrder schemas.CreateOrder
			faker.FakeData(&newOrder)
			reqBody, _ := json.Marshal(newOrder)
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
			apiErr, ok := err.(*ApiError)
			if !ok {
				t.Fatal("err should be an ApiError")
			}
			if apiErr.Status != 400 {
				t.Errorf("Status code should be 400, got %v", recorder.Code)
			}
		})

		t.Run("should return an error when store fails", func(t *testing.T) {
			mocks.OrdersStore.Error = errors.New("")
			recorder := httptest.NewRecorder()
			var newOrder schemas.CreateOrder
			faker.FakeData(&newOrder)
			reqBody, _ := json.Marshal(newOrder)
			request := httptest.NewRequest("POST", "/", bytes.NewReader(reqBody))
			err := controller.Create(recorder, request)
			if err == nil {
				t.Error("Should return an error")
			}
		})
	})

	t.Run("Update", func(t *testing.T) {
		t.Run("should update an order", func(t *testing.T) {
			setUp()
			mocks.OrdersStore.Error = nil
			recorder := httptest.NewRecorder()
			order := mocks.OrdersStore.Store[0]
			reqBody, _ := json.Marshal(order)
			request := httptest.NewRequest("PUT", fmt.Sprintf("/%v", order.ID), bytes.NewReader(reqBody))
			tx := chi.NewRouteContext()
			tx.URLParams.Add("id", fmt.Sprintf("%v", order.ID))
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
			invalidReqBody := `{"foo: 95}`
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("PUT", "/1", bytes.NewReader([]byte(invalidReqBody)))
			tx := chi.NewRouteContext()
			tx.URLParams.Add("id", "1")
			request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
			err := controller.Update(recorder, request)
			apiErr, ok := err.(*ApiError)
			if !ok {
				t.Fatal("err should be an ApiError")
			}
			if apiErr.Status != 400 {
				t.Errorf("Status code should be 400, got %v", recorder.Code)
			}
		})

		t.Run("should return an error when store fails", func(t *testing.T) {
			recorder := httptest.NewRecorder()
			var newOrder schemas.UpdateOrder
			faker.FakeData(&newOrder)
			reqBody, _ := json.Marshal(newOrder)
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
		t.Run("should delete an order", func(t *testing.T) {
			setUp()
			mocks.OrdersStore.Error = nil
			order := mocks.OrdersStore.Store[0]
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("DELETE", fmt.Sprintf("/%v", order.ID), nil)
			tx := chi.NewRouteContext()
			tx.URLParams.Add("id", fmt.Sprintf("%v", order.ID))
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
