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

func TestClient(t *testing.T) {
	repository := &mocks.ClientsRepository{}
	controller := NewClients(repository)

	setUp := func() {
		repository = &mocks.ClientsRepository{}
		controller = NewClients(repository)
	}

	addClients := func(q int) {
		for i := 0; i < q; i++ {
			var client models.Client
			faker.FakeData(&client)
			client.ID = int64(i + 1)
			repository.Store = append(repository.Store, client)
		}
	}

	t.Run("List", func(t *testing.T) {
		setUp()
		addClients(10)
		repository.ShouldError = false
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/", nil)
		err := controller.List(recorder, request)
		if err != nil {
			t.Fatal(err)
		}
		if recorder.Code != 200 {
			t.Errorf("Status code should be 200, got %v", recorder.Code)
		}

		repository.ShouldError = true
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
		addClients(10)
		repository.ShouldError = false
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
		repository.ShouldError = false
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
		if err != nil {
			t.Fatal(err)
		}
		if recorder.Code != 400 {
			t.Errorf("Status code should be 400, got %v", recorder.Code)
		}

		repository.ShouldError = true
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
		addClients(10)
		repository.ShouldError = false
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
		addClients(10)
		repository.ShouldError = false
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
