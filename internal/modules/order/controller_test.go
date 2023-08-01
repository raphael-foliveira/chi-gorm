package order

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"
)

type MockWriter struct {
	mock.Mock
}

func (m *MockWriter) Write(p []byte) (n int, err error) {
	args := m.Called(p)
	return args.Int(0), args.Error(1)
}

func (m *MockWriter) WriteHeader(statusCode int) {
	m.Called(statusCode)
}

func (m *MockWriter) Header() http.Header {
	return http.Header{}
}

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) List() ([]Order, error) {
	args := m.Called()
	return args.Get(0).([]Order), args.Error(1)
}

func (m *MockRepository) Get(id uint64) (Order, error) {
	args := m.Called(id)
	return args.Get(0).(Order), args.Error(1)
}

func (m *MockRepository) Create(c *Order) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockRepository) Update(c *Order) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockRepository) Delete(c *Order) error {
	args := m.Called(c)
	return args.Error(0)
}

func TestControllerList(t *testing.T) {
	t.Run("should return an error when repository returns an error", func(t *testing.T) {
		repository := new(MockRepository)
		controller := NewController(repository)
		repository.On("List").Return([]Order{}, errors.New("test"))
		w := MockWriter{}
		w.On("WriteHeader", 500).Return()
		w.On("Write", mock.Anything).Return(0, nil)
		req, _ := http.NewRequest("GET", "/clients", nil)
		controller.List(&w, req)
	})
}

func TestControllerCreate(t *testing.T) {
	t.Run("should return an error when repository returns an error", func(t *testing.T) {
		repository := new(MockRepository)
		controller := NewController(repository)
		repository.On("Create", mock.Anything).Return(errors.New("test"))
		w := MockWriter{}
		reqB := new(bytes.Buffer)
		json.NewEncoder(reqB).Encode(Order{
			ClientID:  1,
			ProductID: 1,
			Quantity:  1,
		})
		w.On("WriteHeader", 500).Return()
		w.On("Write", mock.Anything).Return(0, nil)
		req, _ := http.NewRequest("POST", "/clients", reqB)
		controller.Create(&w, req)
	})
}

func TestControllerUpdate(t *testing.T) {
	t.Run("should return an error when repository returns an error", func(t *testing.T) {
		repository := new(MockRepository)
		controller := NewController(repository)
		repository.On("Get", mock.Anything).Return(Order{}, nil)
		repository.On("Update", mock.Anything).Return(errors.New("test"))
		w := MockWriter{}
		reqB := new(bytes.Buffer)
		json.NewEncoder(reqB).Encode(Order{
			ClientID:  1,
			ProductID: 1,
			Quantity:  1,
		})
		w.On("WriteHeader", 500).Return()
		w.On("Write", mock.Anything).Return(0, nil)
		req, _ := http.NewRequest("PUT", "/clients/1", reqB)
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
		controller.Update(&w, req)
	})
}

func TestControllerDelete(t *testing.T) {
	t.Run("should return an error when repository returns an error", func(t *testing.T) {
		repository := new(MockRepository)
		controller := NewController(repository)
		repository.On("Get", mock.Anything).Return(Order{}, nil)
		repository.On("Delete", mock.Anything).Return(errors.New("test"))
		w := MockWriter{}
		w.On("WriteHeader", 500).Return()
		w.On("Write", mock.Anything).Return(0, nil)
		req, _ := http.NewRequest("DELETE", "/clients/1", nil)
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
		controller.Delete(&w, req)
	})

	t.Run("should return an error when id is invalid", func(t *testing.T) {
		repository := new(MockRepository)
		controller := NewController(repository)
		w := MockWriter{}
		w.On("WriteHeader", 400).Return()
		w.On("Write", mock.Anything).Return(0, nil)
		req, _ := http.NewRequest("DELETE", "/clients/invalid", nil)
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "invalid")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
		controller.Delete(&w, req)
	})
}

func TestControllerGet(t *testing.T) {
	t.Run("should return an error when id is invalid", func(t *testing.T) {
		repository := new(MockRepository)
		controller := NewController(repository)
		w := MockWriter{}
		w.On("WriteHeader", 400).Return()
		w.On("Write", mock.Anything).Return(0, nil)
		req, _ := http.NewRequest("GET", "/clients/invalid", nil)
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "invalid")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
		controller.Get(&w, req)
	})
}
