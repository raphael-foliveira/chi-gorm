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
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/exceptions"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/stretchr/testify/assert"
)

func TestClient_List(t *testing.T) {
	t.Run("should list all clients", testCase(func(t *testing.T, deps *testDependencies) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/", nil)
		ctx := controller.NewContext(recorder, request)
		err := deps.clientsController.List(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, recorder.Code)
	}))

	t.Run("should return an error when store fails", testCase(func(t *testing.T, deps *testDependencies) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/", nil)
		ctx := controller.NewContext(recorder, request)
		err := deps.clientsController.List(ctx)
		assert.Error(t, err)
	}))
}

func TestClient_Get(t *testing.T) {
	t.Run("should get a client", testCase(func(t *testing.T, deps *testDependencies) {
		client := entities.Client{
			Name:   "John Doe",
			Email:  "john@doe.com",
			Orders: []entities.Order{},
		}
		if db := database.DB.Create(&client); db.Error != nil {
			t.Fatal(db.Error)
		}

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", fmt.Sprintf("/%d", client.ID), nil)
		tx := chi.NewRouteContext()
		tx.URLParams.Add("id", fmt.Sprintf("%d", client.ID))
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		ctx := controller.NewContext(recorder, request)
		err := deps.clientsController.Get(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, recorder.Code)
	}))

	t.Run("should return an error when store fails", testCase(func(t *testing.T, deps *testDependencies) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/99", nil)
		tx := chi.NewRouteContext()
		tx.URLParams.Add("id", "99")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		ctx := controller.NewContext(recorder, request)
		err := deps.clientsController.Get(ctx)
		assert.Error(t, err)
	}))
}

func TestClient_Create(t *testing.T) {
	t.Run("should create a client", testCase(func(t *testing.T, deps *testDependencies) {
		recorder := httptest.NewRecorder()
		var newClient schemas.CreateClient
		faker.FakeData(&newClient)
		reqBody, _ := json.Marshal(newClient)
		request := httptest.NewRequest("POST", "/", bytes.NewReader(reqBody))
		ctx := controller.NewContext(recorder, request)
		err := deps.clientsController.Create(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, recorder.Code)
	}))

	t.Run("should return an error when sent invalid data", testCase(func(t *testing.T, deps *testDependencies) {
		invalidReqBody := `{"foo: 95}`
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("POST", "/", bytes.NewReader([]byte(invalidReqBody)))
		ctx := controller.NewContext(recorder, request)
		err := deps.clientsController.Create(ctx)
		apiErr, ok := err.(*exceptions.ApiError)
		assert.True(t, ok, "err should be an ApiError")
		assert.Equal(t, http.StatusUnprocessableEntity, apiErr.Status)
	}))

	t.Run("should return an error when store fails", testCase(func(t *testing.T, deps *testDependencies) {
		recorder := httptest.NewRecorder()
		var newClient schemas.CreateClient
		faker.FakeData(&newClient)
		reqBody, _ := json.Marshal(newClient)
		request := httptest.NewRequest("POST", "/", bytes.NewReader(reqBody))
		ctx := controller.NewContext(recorder, request)
		err := deps.clientsController.Create(ctx)
		assert.Error(t, err)
	}))
}

func TestClient_Update(t *testing.T) {
	t.Run("should update a client", testCase(func(t *testing.T, deps *testDependencies) {
		recorder := httptest.NewRecorder()
		var newClient schemas.UpdateClient
		faker.FakeData(&newClient)
		reqBody, _ := json.Marshal(newClient)
		request := httptest.NewRequest("PUT", "/1", bytes.NewReader(reqBody))
		tx := chi.NewRouteContext()
		tx.URLParams.Add("id", "1")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		ctx := controller.NewContext(recorder, request)
		err := deps.clientsController.Update(ctx)
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
		err := deps.clientsController.Update(ctx)
		apiErr, ok := err.(*exceptions.ApiError)
		assert.True(t, ok, "err should be an ApiError")
		assert.Equal(t, http.StatusUnprocessableEntity, apiErr.Status)
	}))

	t.Run("should return an error when store fails", testCase(func(t *testing.T, deps *testDependencies) {
		recorder := httptest.NewRecorder()
		var newClient schemas.UpdateClient
		faker.FakeData(&newClient)
		reqBody, _ := json.Marshal(newClient)
		request := httptest.NewRequest("PUT", "/99", bytes.NewReader(reqBody))
		tx := chi.NewRouteContext()
		tx.URLParams.Add("id", "99")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		ctx := controller.NewContext(recorder, request)
		err := deps.clientsController.Update(ctx)
		assert.Error(t, err)
	}))
}

func TestClient_Delete(t *testing.T) {
	t.Run("should delete a client", testCase(func(t *testing.T, deps *testDependencies) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("DELETE", "/1", nil)
		tx := chi.NewRouteContext()
		tx.URLParams.Add("id", "1")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		ctx := controller.NewContext(recorder, request)
		err := deps.clientsController.Delete(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, recorder.Code)
	}))

	t.Run("should return an error when store fails", testCase(func(t *testing.T, deps *testDependencies) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("DELETE", "/99", nil)
		tx := chi.NewRouteContext()
		tx.URLParams.Add("id", "99")
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		ctx := controller.NewContext(recorder, request)
		err := deps.clientsController.Delete(ctx)
		assert.Error(t, err)
	}))
}
