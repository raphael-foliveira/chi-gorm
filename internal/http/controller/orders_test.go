package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/exceptions"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/mocks"
)

func TestOrders(t *testing.T) {
	t.Run("List", func(t *testing.T) {
		addOrders(10)
		mocks.OrdersStore.ShouldError = false
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/", nil)
		err := Orders.List(recorder, request)
		if err != nil {
			t.Fatal(err)
		}
		if recorder.Code != 200 {
			t.Errorf("Status code should be 200, got %v", recorder.Code)
		}

		mocks.OrdersStore.ShouldError = true
		recorder = httptest.NewRecorder()
		request = httptest.NewRequest("GET", "/", nil)
		err = Orders.List(recorder, request)
		if err == nil {
			t.Fatal("err should not be nil")
		}
	})

	t.Run("Get", func(t *testing.T) {
		addOrders(10)
		mocks.OrdersStore.ShouldError = false
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/1", nil)
		tx := chi.NewRouteContext()
		tx.URLParams.Add("id", "1")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		err := Orders.Get(recorder, request)
		if err != nil {
			t.Fatal(err)
		}
		if recorder.Code != 200 {
			t.Errorf("Status code should be 200, got %v", recorder.Code)
		}

		recorder = httptest.NewRecorder()
		request = httptest.NewRequest("GET", "/9999", nil)
		tx.URLParams.Add("id", "9999")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		err = Orders.Get(recorder, request)
		if err == nil {
			t.Fatal("err should not be nil")
		}
	})

	t.Run("Create", func(t *testing.T) {
		mocks.OrdersStore.ShouldError = false
		recorder := httptest.NewRecorder()
		var newOrder schemas.CreateOrder
		faker.FakeData(&newOrder)
		reqBody, _ := json.Marshal(newOrder)
		request := httptest.NewRequest("POST", "/", bytes.NewReader(reqBody))
		err := Orders.Create(recorder, request)
		if err != nil {
			t.Fatal(err)
		}
		if recorder.Code != 201 {
			t.Errorf("Status code should be 201, got %v", recorder.Code)
		}

		invalidReqBody := `{"foo: 95}`
		recorder = httptest.NewRecorder()
		request = httptest.NewRequest("POST", "/", bytes.NewReader([]byte(invalidReqBody)))
		err = Orders.Create(recorder, request)
		apiErr, ok := err.(*exceptions.ApiError)
		if !ok {
			t.Fatal("err should be an ApiError")
		}
		if apiErr.Status != 400 {
			t.Errorf("Status code should be 400, got %v", recorder.Code)
		}

		mocks.OrdersStore.ShouldError = true
		recorder = httptest.NewRecorder()
		request = httptest.NewRequest("POST", "/", bytes.NewReader(reqBody))
		err = Orders.Create(recorder, request)
		if err == nil {
			t.Error("Should return an error")
		}
	})

	t.Run("Update", func(t *testing.T) {
		addOrders(10)
		mocks.OrdersStore.ShouldError = false
		recorder := httptest.NewRecorder()
		var newOrder schemas.UpdateOrder
		faker.FakeData(&newOrder)
		reqBody, _ := json.Marshal(newOrder)
		request := httptest.NewRequest("PUT", "/1", bytes.NewReader(reqBody))
		tx := chi.NewRouteContext()
		tx.URLParams.Add("id", "1")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		err := Orders.Update(recorder, request)
		if err != nil {
			t.Fatal(err)
		}
		if recorder.Code != 200 {
			t.Errorf("Status code should be 200, got %v", recorder.Code)
		}

		invalidReqBody := `{"foo: 95}`
		recorder = httptest.NewRecorder()
		request = httptest.NewRequest("PUT", "/1", bytes.NewReader([]byte(invalidReqBody)))
		tx = chi.NewRouteContext()
		tx.URLParams.Add("id", "1")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		err = Orders.Update(recorder, request)
		apiErr, ok := err.(*exceptions.ApiError)
		if !ok {
			t.Fatal("err should be an ApiError")
		}
		if apiErr.Status != 400 {
			t.Errorf("Status code should be 400, got %v", recorder.Code)
		}

		recorder = httptest.NewRecorder()
		request = httptest.NewRequest("PUT", "/9999", bytes.NewReader(reqBody))
		tx = chi.NewRouteContext()
		tx.URLParams.Add("id", "9999")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		err = Orders.Update(recorder, request)
		if err == nil {
			t.Error("Should return an error")
		}
	})

	t.Run("Delete", func(t *testing.T) {
		addOrders(10)
		mocks.OrdersStore.ShouldError = false
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("DELETE", "/1", nil)
		tx := chi.NewRouteContext()
		tx.URLParams.Add("id", "1")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		err := Orders.Delete(recorder, request)
		if err != nil {
			t.Fatal(err)
		}
		if recorder.Code != 204 {
			t.Errorf("Status code should be 204, got %v", recorder.Code)
		}

		recorder = httptest.NewRecorder()
		request = httptest.NewRequest("DELETE", "/9999", nil)
		tx = chi.NewRouteContext()
		tx.URLParams.Add("id", "9999")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		err = Orders.Delete(recorder, request)
		if err == nil {
			t.Error("Should return an error")
		}
	})

}
