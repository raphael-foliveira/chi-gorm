package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/pkg/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/pkg/mocks"
	"github.com/raphael-foliveira/chi-gorm/pkg/models"
)

func TestOrders(t *testing.T) {
	var ordersStore *mocks.OrdersStore
	var clientsStore *mocks.ClientsStore
	var productsStore *mocks.ProductsStore
	var controller *Orders

	setUp := func() {
		ordersStore = &mocks.OrdersStore{}
		clientsStore = &mocks.ClientsStore{}
		productsStore = &mocks.ProductsStore{}
		controller = NewOrders(ordersStore, clientsStore, productsStore)
	}

	addOrders := func(q int) {
		for i := 0; i < q; i++ {
			var order *models.Order
			var client *models.Client
			var product *models.Product
			faker.FakeData(&order)
			faker.FakeData(&client)
			faker.FakeData(&product)
			order.ID = int64(i + 1)
			client.ID = int64(i + 1)
			product.ID = int64(i + 1)
			order.ClientID = client.ID
			order.ProductID = product.ID
			ordersStore.Store = append(ordersStore.Store, *order)
			clientsStore.Store = append(clientsStore.Store, *client)
			productsStore.Store = append(productsStore.Store, *product)
		}
	}

	t.Run("List", func(t *testing.T) {
		setUp()
		addOrders(10)
		ordersStore.ShouldError = false
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/", nil)
		err := controller.List(recorder, request)
		if err != nil {
			t.Fatal(err)
		}
		if recorder.Code != 200 {
			t.Errorf("Status code should be 200, got %v", recorder.Code)
		}

		ordersStore.ShouldError = true
		recorder = httptest.NewRecorder()
		request = httptest.NewRequest("GET", "/", nil)
		err = controller.List(recorder, request)
		if err != nil {
			t.Fatal(err)
		}
		if recorder.Code != 500 {
			t.Errorf("Status code should be 500, got %v", recorder.Code)
		}
	})

	t.Run("Get", func(t *testing.T) {
		setUp()
		addOrders(10)
		ordersStore.ShouldError = false
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

		recorder = httptest.NewRecorder()
		request = httptest.NewRequest("GET", "/99", nil)
		tx.URLParams.Add("id", "99")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		err = controller.Get(recorder, request)
		if err != nil {
			t.Fatal(err)
		}
		if recorder.Code != 404 {
			t.Errorf("Status code should be 404, got %v", recorder.Code)
		}
	})

	t.Run("Create", func(t *testing.T) {
		setUp()
		ordersStore.ShouldError = false
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

		invalidReqBody := `{"foo: 95}`
		recorder = httptest.NewRecorder()
		request = httptest.NewRequest("POST", "/", bytes.NewReader([]byte(invalidReqBody)))
		err = controller.Create(recorder, request)
		if err != nil {
			t.Fatal(err)
		}
		if recorder.Code != 400 {
			t.Errorf("Status code should be 400, got %v", recorder.Code)
		}

		ordersStore.ShouldError = true
		recorder = httptest.NewRecorder()
		request = httptest.NewRequest("POST", "/", bytes.NewReader(reqBody))
		err = controller.Create(recorder, request)
		if err != nil {
			t.Fatal(err)
		}
		if recorder.Code != 500 {
			t.Errorf("Status code should be 500, got %v", recorder.Code)
		}
	})

	t.Run("Update", func(t *testing.T) {
		setUp()
		addOrders(10)
		ordersStore.ShouldError = false
		recorder := httptest.NewRecorder()
		var newOrder schemas.UpdateOrder
		faker.FakeData(&newOrder)
		reqBody, _ := json.Marshal(newOrder)
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
		if err != nil {
			t.Fatal(err)
		}
		if recorder.Code != 400 {
			t.Errorf("Status code should be 400, got %v", recorder.Code)
		}

		recorder = httptest.NewRecorder()
		request = httptest.NewRequest("PUT", "/99", bytes.NewReader(reqBody))
		tx = chi.NewRouteContext()
		tx.URLParams.Add("id", "99")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		err = controller.Update(recorder, request)
		if err != nil {
			t.Fatal(err)
		}
		if recorder.Code != 404 {
			t.Errorf("Status code should be 404, got %v", recorder.Code)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		setUp()
		addOrders(10)
		ordersStore.ShouldError = false
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
		if err != nil {
			t.Fatal(err)
		}
		if recorder.Code != 404 {
			t.Errorf("Status code should be 404, got %v", recorder.Code)
		}
	})

}
