package order

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/go-chi/chi/v5"
)

type MockRepository struct {
	db  []Order
	err bool
}

func (m *MockRepository) List() ([]Order, error) {
	if m.err {
		return nil, errors.New("test")
	}
	return m.db, nil
}

func (m *MockRepository) Get(id uint64) (Order, error) {
	if m.err {
		return Order{}, errors.New("test")
	}
	for _, o := range m.db {
		if uint64(o.ID) == id {
			return o, nil
		}
	}
	return Order{}, errors.New("not found")
}

func (m *MockRepository) Create(c *Order) error {
	if m.err {
		return errors.New("test")
	}
	m.db = append(m.db, *c)
	return nil
}

func (m *MockRepository) Update(c *Order) error {
	if m.err {
		return errors.New("test")
	}
	for i, o := range m.db {
		if o.ID == c.ID {
			m.db[i] = *c
			return nil
		}
	}
	return errors.New("not found")
}

func (m *MockRepository) Delete(c *Order) error {
	if m.err {
		return errors.New("test")
	}
	for i, o := range m.db {
		if o.ID == c.ID {
			m.db = append(m.db[:i], m.db[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}

var testRouter *chi.Mux
var mockRepository *MockRepository

func InsertOrdersHelper(qt int) {
	for i := 0; i < qt; i++ {
		order := Order{}
		err := faker.FakeData(&order)
		if err != nil {
			panic(err)
		}
		mockRepository.db = append(mockRepository.db, order)
	}
}

func ClearOrdersTable() {
	mockRepository.db = []Order{}
}

func TestMain(m *testing.M) {
	testRouter = chi.NewRouter()
	mockRepository = new(MockRepository)
	controller := NewController(mockRepository)
	ordersRouter := NewRouter(controller)

	testRouter.Mount("/orders", ordersRouter)
	ClearOrdersTable()
	code := m.Run()
	os.Exit(code)
}

func TestList(t *testing.T) {
	t.Run("should return an empty list when there are no orders", func(t *testing.T) {
		ClearOrdersTable()
		req, err := http.NewRequest("GET", "/orders", nil)
		if err != nil {
			t.Fatal(err)
		}
		rec := httptest.NewRecorder()
		testRouter.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status code %v, got %v", http.StatusOK, rec.Code)
		}

		if strings.TrimSpace(rec.Body.String()) != "[]" {
			t.Errorf("Expected body %v, got %v", "[]", rec.Body.String())
		}
	})

	t.Run("should return a populated list when there are orders", func(t *testing.T) {
		ClearOrdersTable()
		InsertOrdersHelper(10)
		req, err := http.NewRequest("GET", "/orders", nil)
		if err != nil {
			t.Fatal(err)
		}
		rec := httptest.NewRecorder()
		testRouter.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status code %v, got %v", http.StatusOK, rec.Code)
		}

		if strings.TrimSpace(rec.Body.String()) == "[]" {
			t.Errorf("Expected body %v, got %v", "[]", rec.Body.String())
		}
	})

	t.Run("should return an error when repository fails", func(t *testing.T) {
		ClearOrdersTable()
		mockRepository.err = true
		req, err := http.NewRequest("GET", "/orders", nil)
		if err != nil {
			t.Fatal(err)
		}
		rec := httptest.NewRecorder()
		testRouter.ServeHTTP(rec, req)

		if rec.Code != http.StatusInternalServerError {
			t.Errorf("Expected status code %v, got %v", http.StatusInternalServerError, rec.Code)
		}

		if !strings.Contains(rec.Body.String(), "error") {
			t.Errorf("Expected body %v, got %v", "error", rec.Body.String())
		}
		mockRepository.err = false
	})
}

func TestGet(t *testing.T) {
	t.Run("should return 404 when order does not exist", func(t *testing.T) {
		ClearOrdersTable()
		req, err := http.NewRequest("GET", "/orders/1", nil)
		if err != nil {
			t.Fatal(err)
		}
		rec := httptest.NewRecorder()
		testRouter.ServeHTTP(rec, req)

		if rec.Code != http.StatusNotFound {
			t.Errorf("Expected status code %v, got %v", http.StatusNotFound, rec.Code)
		}
		if !strings.Contains(rec.Body.String(), "error") {
			t.Errorf("Expected body %v, got %v", "error", rec.Body.String())
		}
	})

	t.Run("should return 200 when order exists", func(t *testing.T) {
		ClearOrdersTable()
		InsertOrdersHelper(10)
		order := mockRepository.db[0]
		req, err := http.NewRequest("GET", fmt.Sprintf("/orders/%v", order.ID), nil)
		if err != nil {
			t.Fatal(err)
		}
		rec := httptest.NewRecorder()
		testRouter.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status code %v, got %v", http.StatusOK, rec.Code)
		}
	})

	t.Run("should return 400 when id is invalid", func(t *testing.T) {
		ClearOrdersTable()
		req, err := http.NewRequest("GET", "/orders/invalid", nil)
		if err != nil {
			t.Fatal(err)
		}
		rec := httptest.NewRecorder()
		testRouter.ServeHTTP(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %v, got %v", http.StatusBadRequest, rec.Code)
		}
	})
}

func TestCreate(t *testing.T) {
	t.Run("should return 400 when body is invalid", func(t *testing.T) {
		ClearOrdersTable()
		req, err := http.NewRequest("POST", "/orders", strings.NewReader("invalid body"))
		if err != nil {
			t.Fatal(err)
		}
		rec := httptest.NewRecorder()
		testRouter.ServeHTTP(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %v, got %v", http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("should return 201 when body is valid", func(t *testing.T) {
		ClearOrdersTable()
		InsertOrdersHelper(5)
		order := CreateOrderSchema{}
		err := faker.FakeData(&order)
		if err != nil {
			t.Fatal(err)
		}
		buf := new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(order)
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest("POST", "/orders", buf)
		if err != nil {
			t.Fatal(err)
		}
		rec := httptest.NewRecorder()
		testRouter.ServeHTTP(rec, req)
		newOrder := Order{}
		json.NewDecoder(rec.Body).Decode(&newOrder)

		if rec.Code != http.StatusCreated {
			t.Errorf("Expected status code %v, got %v", http.StatusCreated, rec.Code)
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("should return 404 when order does not exist", func(t *testing.T) {
		ClearOrdersTable()
		req, err := http.NewRequest("DELETE", "/orders/1", nil)
		if err != nil {
			t.Fatal(err)
		}
		rec := httptest.NewRecorder()
		testRouter.ServeHTTP(rec, req)

		if rec.Code != http.StatusNotFound {
			t.Errorf("Expected status code %v, got %v", http.StatusNotFound, rec.Code)
		}
		if !strings.Contains(rec.Body.String(), "error") {
			t.Errorf("Expected body %v, got %v", "error", rec.Body.String())
		}
	})

	t.Run("should return 204 when order exists", func(t *testing.T) {
		ClearOrdersTable()
		InsertOrdersHelper(1)
		order := mockRepository.db[0]
		req, err := http.NewRequest("DELETE", fmt.Sprintf("/orders/%v", order.ID), nil)
		if err != nil {
			t.Fatal(err)
		}
		rec := httptest.NewRecorder()
		testRouter.ServeHTTP(rec, req)

		if rec.Code != http.StatusNoContent {
			t.Errorf("Expected status code %v, got %v", http.StatusOK, rec.Code)
		}
	})

	t.Run("should return 400 when id is invalid", func(t *testing.T) {
		ClearOrdersTable()
		req, err := http.NewRequest("DELETE", "/orders/invalid", nil)
		if err != nil {
			t.Fatal(err)
		}
		rec := httptest.NewRecorder()
		testRouter.ServeHTTP(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %v, got %v", http.StatusBadRequest, rec.Code)
		}
	})
}

func TestUpdate(t *testing.T) {
	t.Run("should return 404 when order does not exist", func(t *testing.T) {
		ClearOrdersTable()
		req, err := http.NewRequest("PUT", "/orders/1", nil)
		if err != nil {
			t.Fatal(err)
		}
		rec := httptest.NewRecorder()
		testRouter.ServeHTTP(rec, req)

		if rec.Code != http.StatusNotFound {
			t.Errorf("Expected status code %v, got %v", http.StatusNotFound, rec.Code)
		}
		if !strings.Contains(rec.Body.String(), "error") {
			t.Errorf("Expected body %v, got %v", "error", rec.Body.String())
		}
	})

	t.Run("should return 200 when order exists", func(t *testing.T) {
		ClearOrdersTable()
		InsertOrdersHelper(1)
		order := mockRepository.db[0]
		order.Quantity = 30
		buf := new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(order)
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest("PUT", fmt.Sprintf("/orders/%v", order.ID), buf)
		if err != nil {
			t.Fatal(err)
		}
		rec := httptest.NewRecorder()
		testRouter.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status code %v, got %v", http.StatusOK, rec.Code)
		}

		if !strings.Contains(rec.Body.String(), fmt.Sprint(order.Quantity)) {
			t.Errorf("Expected body %v, got %v", 30, rec.Body.String())
		}
	})

	t.Run("should return 400 when body is invalid", func(t *testing.T) {
		ClearOrdersTable()
		InsertOrdersHelper(1)
		order := mockRepository.db[0]
		order.Quantity = 30
		buf := new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode("invalid body")
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest("PUT", fmt.Sprintf("/orders/%v", order.ID), buf)
		if err != nil {
			t.Fatal(err)
		}
		rec := httptest.NewRecorder()
		testRouter.ServeHTTP(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %v, got %v", http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("should return 400 when id is invalid", func(t *testing.T) {
		ClearOrdersTable()
		req, err := http.NewRequest("PUT", "/orders/invalid", nil)
		if err != nil {
			t.Fatal(err)
		}
		rec := httptest.NewRecorder()
		testRouter.ServeHTTP(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %v, got %v", http.StatusBadRequest, rec.Code)
		}
	})
}
