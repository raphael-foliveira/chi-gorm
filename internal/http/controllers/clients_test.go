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
	"github.com/raphael-foliveira/chi-gorm/internal/services"
)

func TestClient(t *testing.T) {
	var ordersStore *mocks.OrdersStore
	var clientsStore *mocks.ClientsStore
	var productsStore *mocks.ProductsStore
	var controller *Clients

	setUp := func() {
		ordersStore = &mocks.OrdersStore{}
		clientsStore = &mocks.ClientsStore{}
		productsStore = &mocks.ProductsStore{}
		clientsService := services.NewClients(clientsStore)
		controller = NewClients(clientsService)
	}

	addClients := func(q int) {
		for i := 0; i < q; i++ {
			var client entities.Client
			var product entities.Product
			faker.FakeData(&client)
			faker.FakeData(&product)
			productsStore.Store = append(productsStore.Store, product)
			for j := 0; j < 10; j++ {
				var order entities.Order
				faker.FakeData(&order)
				order.ClientID = client.ID
				order.ProductID = product.ID
				ordersStore.Store = append(ordersStore.Store, order)
			}
			client.ID = int64(i + 1)
			clientsStore.Store = append(clientsStore.Store, client)
		}
	}

	t.Run("List", func(t *testing.T) {
		setUp()
		addClients(10)
		clientsStore.ShouldError = false
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/", nil)
		err := controller.List(recorder, request)
		if err != nil {
			t.Fatal(err)
		}
		if recorder.Code != 200 {
			t.Errorf("Status code should be 200, got %v", recorder.Code)
		}

		clientsStore.ShouldError = true
		recorder = httptest.NewRecorder()
		request = httptest.NewRequest("GET", "/", nil)
		err = controller.List(recorder, request)
		if err == nil {
			t.Fatal("err should not be nil")
		}
	})

	t.Run("Get", func(t *testing.T) {
		setUp()
		addClients(10)
		clientsStore.ShouldError = false
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

		clientsStore.ShouldError = true
		recorder = httptest.NewRecorder()
		request = httptest.NewRequest("GET", "/99", nil)
		tx = chi.NewRouteContext()
		tx.URLParams.Add("id", "99")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		err = controller.Get(recorder, request)
		if err == nil {
			t.Fatal("err should not be nil")
		}
	})

	t.Run("Create", func(t *testing.T) {
		setUp()
		clientsStore.ShouldError = false
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

		invalidReqBody := `{"foo: 95}`
		recorder = httptest.NewRecorder()
		request = httptest.NewRequest("POST", "/", bytes.NewReader([]byte(invalidReqBody)))
		err = controller.Create(recorder, request)
		apiErr, ok := err.(*exceptions.ApiError)
		if !ok {
			t.Fatal("err should be an ApiError")
		}
		if apiErr.Status != 400 {
			t.Errorf("Status code should be 400, got %v", recorder.Code)
		}

		clientsStore.ShouldError = true
		recorder = httptest.NewRecorder()
		request = httptest.NewRequest("POST", "/", bytes.NewReader(reqBody))
		err = controller.Create(recorder, request)
		if err == nil {
			t.Fatal("err should not be nil")
		}
	})

	t.Run("Update", func(t *testing.T) {
		setUp()
		addClients(10)
		clientsStore.ShouldError = false
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

		invalidReqBody := `{"foo: 95}`
		recorder = httptest.NewRecorder()
		request = httptest.NewRequest("PUT", "/1", bytes.NewReader([]byte(invalidReqBody)))
		tx = chi.NewRouteContext()
		tx.URLParams.Add("id", "1")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		err = controller.Update(recorder, request)
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
		err = controller.Update(recorder, request)
		if err == nil {
			t.Fatal("err should not be nil")
		}

	})

	t.Run("Delete", func(t *testing.T) {
		setUp()
		addClients(10)
		clientsStore.ShouldError = false
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

		recorder = httptest.NewRecorder()
		request = httptest.NewRequest("DELETE", "/99", nil)
		tx = chi.NewRouteContext()
		tx.URLParams.Add("id", "99")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		err = controller.Delete(recorder, request)
		if err == nil {
			t.Fatal("err should not be nil")
		}
	})
}
