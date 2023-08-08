package order

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal"
	"github.com/raphael-foliveira/chi-gorm/internal/db"
	"github.com/raphael-foliveira/chi-gorm/internal/modules/client"
	"github.com/raphael-foliveira/chi-gorm/internal/modules/product"
)

var database *db.DB
var router *chi.Mux

func InsertOrdersHelper(qt int) {
	for i := 0; i < qt; i++ {
		c := client.Client{}
		err := faker.FakeData(&c)
		c.ID = 0
		if err != nil {
			panic(err)
		}
		tx := database.Create(&c)
		if tx.Error != nil {
			panic(tx.Error)
		}
	}
	for i := 0; i < qt; i++ {
		p := product.Product{}
		err := faker.FakeData(&p)
		p.ID = 0
		if err != nil {
			panic(err)
		}
		tx := database.Create(&p)
		if tx.Error != nil {
			panic(tx.Error)
		}
	}
	for i := 0; i < qt; i++ {
		c := client.Client{}
		p := product.Product{}

		tx := database.Order("RANDOM()").First(&c)
		if tx.Error != nil {
			panic(tx.Error)
		}
		tx = database.Order("RANDOM()").First(&p)
		if tx.Error != nil {
			panic(tx.Error)
		}
		o := Order{}
		err := faker.FakeData(&o)
		o.ClientID = 0
		o.ProductID = 0
		o.Client = c
		o.Product = p
		fmt.Println(o)
		if err != nil {
			panic(err)
		}
		tx = database.Create(&o)
		if tx.Error != nil {
			panic(tx.Error)
		}
	}
}

func ClearOrdersTable() {
	database.Delete(&Order{}, "1=1")
	database.Delete(&client.Client{}, "1=1")
	database.Delete(&product.Product{}, "1=1")
}

func TestMain(m *testing.M) {
	database = db.Connect(internal.TestConfig.DatabaseURL)
	router = chi.NewRouter()
	ordersRouter, err := NewRouter(database)
	if err != nil {
		panic(err)
	}
	router.Mount("/orders", ordersRouter)
	ClearOrdersTable()
	code := m.Run()
	ClearOrdersTable()
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
		router.ServeHTTP(rec, req)

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
		router.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status code %v, got %v", http.StatusOK, rec.Code)
		}

		if strings.TrimSpace(rec.Body.String()) == "[]" {
			t.Errorf("Expected body %v, got %v", "[]", rec.Body.String())
		}
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
		router.ServeHTTP(rec, req)

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
		order := Order{}
		database.First(&order)
		req, err := http.NewRequest("GET", fmt.Sprintf("/orders/%v", order.ID), nil)
		if err != nil {
			t.Fatal(err)
		}
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

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
		router.ServeHTTP(rec, req)

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
		router.ServeHTTP(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %v, got %v", http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("should return 201 when body is valid", func(t *testing.T) {
		ClearOrdersTable()
		InsertOrdersHelper(5)
		cli := client.Client{}
		pro := product.Product{}

		tx := database.Order("RANDOM()").First(&cli)
		if tx.Error != nil {
			panic(tx.Error)
		}
		tx = database.Order("RANDOM()").First(&pro)
		if tx.Error != nil {
			panic(tx.Error)
		}
		order := CreateOrderSchema{
			ClientID:  cli.ID,
			ProductID: pro.ID,
			Quantity:  10,
		}
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
		router.ServeHTTP(rec, req)

		if rec.Code != http.StatusCreated {
			t.Errorf("Expected status code %v, got %v", http.StatusCreated, rec.Code)
		}

		if !strings.Contains(rec.Body.String(), fmt.Sprint(order.Quantity)) {
			t.Errorf("Expected body %v, got %v", order.Quantity, rec.Body.String())
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
		router.ServeHTTP(rec, req)

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
		order := Order{}
		database.Order("RANDOM()").First(&order)
		req, err := http.NewRequest("DELETE", fmt.Sprintf("/orders/%v", order.ID), nil)
		if err != nil {
			t.Fatal(err)
		}
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

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
		router.ServeHTTP(rec, req)

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
		router.ServeHTTP(rec, req)

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
		order := Order{}
		database.First(&order)
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
		router.ServeHTTP(rec, req)

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
		order := Order{}
		database.First(&order)
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
		router.ServeHTTP(rec, req)

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
		router.ServeHTTP(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %v, got %v", http.StatusBadRequest, rec.Code)
		}
	})
}
