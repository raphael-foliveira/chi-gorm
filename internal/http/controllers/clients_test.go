package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/exceptions"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/mocks"
)

func TestClient(t *testing.T) {

	addClients := func(q int) {
		for i := 0; i < q; i++ {
			var client entities.Client
			var product entities.Product
			faker.FakeData(&client)
			faker.FakeData(&product)
			mocks.ProductsStore.Store = append(mocks.ProductsStore.Store, product)
			for j := 0; j < 10; j++ {
				var order entities.Order
				faker.FakeData(&order)
				order.ClientID = client.ID
				order.ProductID = product.ID
				mocks.OrdersStore.Store = append(mocks.OrdersStore.Store, order)
			}
			client.ID = uint(i + 1)
			mocks.ClientsStore.Store = append(mocks.ClientsStore.Store, client)
		}
	}

	t.Run("List", func(t *testing.T) {
		addClients(10)
		mocks.ClientsStore.ShouldError = false
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/", nil)
		err := Clients.List(recorder, request)
		if err != nil {
			t.Fatal(err)
		}
		if recorder.Code != 200 {
			t.Errorf("Status code should be 200, got %v", recorder.Code)
		}

		mocks.ClientsStore.ShouldError = true
		recorder = httptest.NewRecorder()
		request = httptest.NewRequest("GET", "/", nil)
		err = Clients.List(recorder, request)
		if err == nil {
			t.Fatal("err should not be nil")
		}
	})

	t.Run("Get", func(t *testing.T) {
		addClients(10)
		mocks.ClientsStore.ShouldError = false
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/1", nil)
		tx := chi.NewRouteContext()
		tx.URLParams.Add("id", "1")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		err := Clients.Get(recorder, request)
		if err != nil {
			t.Fatal(err)
		}
		if recorder.Code != 200 {
			t.Errorf("Status code should be 200, got %v", recorder.Code)
		}

		mocks.ClientsStore.ShouldError = true
		recorder = httptest.NewRecorder()
		request = httptest.NewRequest("GET", "/99", nil)
		tx = chi.NewRouteContext()
		tx.URLParams.Add("id", "99")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		err = Clients.Get(recorder, request)
		if err == nil {
			t.Fatal("err should not be nil")
		}
	})

	t.Run("Create", func(t *testing.T) {
		mocks.ClientsStore.ShouldError = false
		recorder := httptest.NewRecorder()
		var newClient schemas.CreateClient
		faker.FakeData(&newClient)
		reqBody, _ := json.Marshal(newClient)
		request := httptest.NewRequest("POST", "/", bytes.NewReader(reqBody))
		err := Clients.Create(recorder, request)
		if err != nil {
			t.Fatal(err)
		}
		if recorder.Code != 201 {
			t.Errorf("Status code should be 201, got %v", recorder.Code)
		}

		invalidReqBody := `{"foo: 95}`
		recorder = httptest.NewRecorder()
		request = httptest.NewRequest("POST", "/", bytes.NewReader([]byte(invalidReqBody)))
		err = Clients.Create(recorder, request)
		apiErr, ok := err.(*exceptions.ApiError)
		if !ok {
			t.Fatal("err should be an ApiError")
		}
		if apiErr.Status != 400 {
			t.Errorf("Status code should be 400, got %v", recorder.Code)
		}

		mocks.ClientsStore.ShouldError = true
		recorder = httptest.NewRecorder()
		request = httptest.NewRequest("POST", "/", bytes.NewReader(reqBody))
		err = Clients.Create(recorder, request)
		if err == nil {
			t.Fatal("err should not be nil")
		}
	})

	t.Run("Update", func(t *testing.T) {
		addClients(10)
		mocks.ClientsStore.ShouldError = false
		recorder := httptest.NewRecorder()
		var newClient schemas.UpdateClient
		faker.FakeData(&newClient)
		reqBody, _ := json.Marshal(newClient)
		request := httptest.NewRequest("PUT", "/1", bytes.NewReader(reqBody))
		tx := chi.NewRouteContext()
		tx.URLParams.Add("id", "1")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		err := Clients.Update(recorder, request)
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
		err = Clients.Update(recorder, request)
		apiErr, ok := err.(*exceptions.ApiError)
		if !ok {
			t.Fatal("err should be an ApiError")
		}
		if apiErr.Status != 400 {
			t.Errorf("Status code should be 400, got %v", recorder.Code)
		}

		recorder = httptest.NewRecorder()
		request = httptest.NewRequest("PUT", "/99", bytes.NewReader(reqBody))
		tx = chi.NewRouteContext()
		tx.URLParams.Add("id", "99")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		err = Clients.Update(recorder, request)
		if err == nil {
			t.Fatal("err should not be nil")
		}

	})

	t.Run("Delete", func(t *testing.T) {
		addClients(10)
		mocks.ClientsStore.ShouldError = false
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("DELETE", "/1", nil)
		tx := chi.NewRouteContext()
		tx.URLParams.Add("id", "1")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		err := Clients.Delete(recorder, request)
		if err != nil {
			t.Fatal(err)
		}
		if recorder.Code != 204 {
			t.Errorf("Status code should be 204, got %v", recorder.Code)
		}

		recorder = httptest.NewRecorder()
		request = httptest.NewRequest("DELETE", "/99", nil)
		tx = chi.NewRouteContext()
		tx.URLParams.Add("id", "99")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		err = Clients.Delete(recorder, request)
		if err == nil {
			t.Fatal("err should not be nil")
		}
	})
}
