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
}

func TestClient_Get(t *testing.T) {
	t.Run("should get a client", testCase(func(t *testing.T, deps *testDependencies) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", fmt.Sprintf("/%d", deps.clientsStubs[0].ID), nil)
		tx := chi.NewRouteContext()
		tx.URLParams.Add("id", fmt.Sprintf("%d", deps.clientsStubs[0].ID))
		request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tx))
		ctx := controller.NewContext(recorder, request)
		err := deps.clientsController.Get(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, recorder.Code)
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
}
