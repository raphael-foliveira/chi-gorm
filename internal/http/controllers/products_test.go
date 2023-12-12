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
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/mocks"
)

func TestProducts(t *testing.T) {
	store := &mocks.ProductsStore{}
	controller := NewProducts(store)

	setUp := func() {
		store = &mocks.ProductsStore{}
		controller = NewProducts(store)
	}

	addProducts := func(q int) {
		for i := 0; i < q; i++ {
			var product entities.Product
			faker.FakeData(&product)
			product.ID = int64(i + 1)
			store.Store = append(store.Store, product)
		}
	}

	t.Run("List", func(t *testing.T) {
		setUp()
		addProducts(10)
		store.ShouldError = false
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/", nil)
		err := controller.List(recorder, request)
		if err != nil {
			t.Fatal(err)
		}
		if recorder.Code != 200 {
			t.Errorf("Status code should be 200, got %v", recorder.Code)
		}

		store.ShouldError = true
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
		addProducts(10)
		store.ShouldError = false
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
		controller.Get(recorder, request)
		if recorder.Code != 404 {
			t.Errorf("Status code should be 404, got %v", recorder.Code)
		}
	})

	t.Run("Create", func(t *testing.T) {
		setUp()
		store.ShouldError = false
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

		invalidReqBody := `{"foo: 95}`
		recorder = httptest.NewRecorder()
		request = httptest.NewRequest("POST", "/", bytes.NewReader([]byte(invalidReqBody)))
		controller.Create(recorder, request)
		if recorder.Code != 400 {
			t.Errorf("Status code should be 400, got %v", recorder.Code)
		}

		store.ShouldError = true
		recorder = httptest.NewRecorder()
		request = httptest.NewRequest("POST", "/", bytes.NewReader(reqBody))
		controller.Create(recorder, request)
		if recorder.Code != 500 {
			t.Errorf("Status code should be 500, got %v", recorder.Code)
		}
	})

	t.Run("Update", func(t *testing.T) {
		setUp()
		addProducts(10)
		store.ShouldError = false
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

		invalidReqBody := `{"foo: 95}`
		recorder = httptest.NewRecorder()
		request = httptest.NewRequest("PUT", "/1", bytes.NewReader([]byte(invalidReqBody)))
		tx = chi.NewRouteContext()
		tx.URLParams.Add("id", "1")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		controller.Update(recorder, request)
		if recorder.Code != 400 {
			t.Errorf("Status code should be 400, got %v", recorder.Code)
		}

		recorder = httptest.NewRecorder()
		request = httptest.NewRequest("PUT", "/99", bytes.NewReader(reqBody))
		tx = chi.NewRouteContext()
		tx.URLParams.Add("id", "99")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		controller.Update(recorder, request)
		if recorder.Code != 404 {
			t.Errorf("Status code should be 404, got %v", recorder.Code)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		setUp()
		addProducts(10)
		store.ShouldError = false
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
		controller.Delete(recorder, request)
		if recorder.Code != 404 {
			t.Errorf("Status code should be 404, got %v", recorder.Code)
		}
	})
}
