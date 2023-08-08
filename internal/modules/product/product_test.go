package product

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
)

var database *db.DB
var router *chi.Mux

func InsertProductsHelper(qt int) {
	for i := 0; i < qt; i++ {
		product := Product{}
		err := faker.FakeData(&product)
		product.ID = 0
		if err != nil {
			panic(err)
		}
		tx := database.Create(&product)
		if tx.Error != nil {
			panic(tx.Error)
		}
	}
}

func ClearProductsTable() {
	database.Raw("delete from orders")
	tx := database.Delete(&Product{}, "1=1")
	if tx.Error != nil {
		panic(tx.Error)
	}
}

func TestMain(m *testing.M) {
	database = db.Connect(internal.TestConfig.DatabaseURL)
	router = chi.NewRouter()
	productsRouter, err := NewRouter(database)
	if err != nil {
		panic(err)
	}
	router.Mount("/products", productsRouter)
	ClearProductsTable()
	code := m.Run()
	ClearProductsTable()
	os.Exit(code)
}

func TestList(t *testing.T) {
	t.Run("should return an empty list when there are no products", func(t *testing.T) {
		ClearProductsTable()
		req, err := http.NewRequest("GET", "/products", nil)
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

	t.Run("should return a populated list when there are products", func(t *testing.T) {
		ClearProductsTable()
		InsertProductsHelper(10)
		req, err := http.NewRequest("GET", "/products", nil)
		if err != nil {
			t.Fatal(err)
		}
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status code %v, got %v", http.StatusOK, rec.Code)
		}

		if rec.Body.String() == "[]" {
			t.Errorf("Expected body %v, got %v", "[]", rec.Body.String())
		}
	})
}

func TestGet(t *testing.T) {
	t.Run("should return 404 when product does not exist", func(t *testing.T) {
		ClearProductsTable()
		req, err := http.NewRequest("GET", "/products/1", nil)
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

	t.Run("should return 200 when product exists", func(t *testing.T) {
		ClearProductsTable()
		InsertProductsHelper(1)
		product := Product{}
		tx := database.First(&product)
		if tx.Error != nil {
			t.Fatal(tx.Error)
		}
		req, err := http.NewRequest("GET", fmt.Sprintf("/products/%v", product.ID), nil)
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
		ClearProductsTable()
		req, err := http.NewRequest("GET", "/products/invalid", nil)
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
		ClearProductsTable()
		req, err := http.NewRequest("POST", "/products", strings.NewReader("invalid body"))
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
		ClearProductsTable()
		product := Product{}
		err := faker.FakeData(&product)
		if err != nil {
			t.Fatal(err)
		}
		buf := new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(product)
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest("POST", "/products", buf)
		if err != nil {
			t.Fatal(err)
		}
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		if rec.Code != http.StatusCreated {
			t.Errorf("Expected status code %v, got %v", http.StatusCreated, rec.Code)
		}

		if !strings.Contains(rec.Body.String(), product.Name) {
			t.Errorf("Expected body %v, got %v", product.Name, rec.Body.String())
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("should return 404 when product does not exist", func(t *testing.T) {
		ClearProductsTable()
		req, err := http.NewRequest("DELETE", "/products/1", nil)
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

	t.Run("should return 204 when product exists", func(t *testing.T) {
		ClearProductsTable()
		product := Product{}
		err := faker.FakeData(&product)
		if err != nil {
			t.Fatal(err)
		}
		tx := database.Create(&product)
		if tx.Error != nil {
			t.Fatal(tx.Error)
		}
		req, err := http.NewRequest("DELETE", fmt.Sprintf("/products/%v", product.ID), nil)
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
		ClearProductsTable()
		req, err := http.NewRequest("DELETE", "/products/invalid", nil)
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
	t.Run("should return 404 when product does not exist", func(t *testing.T) {
		ClearProductsTable()
		req, err := http.NewRequest("PUT", "/products/1", nil)
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

	t.Run("should return 200 when product exists", func(t *testing.T) {
		ClearProductsTable()
		product := Product{}
		err := faker.FakeData(&product)
		if err != nil {
			t.Fatal(err)
		}
		tx := database.Create(&product)
		if tx.Error != nil {
			t.Fatal(tx.Error)
		}
		buf := new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(product)
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest("PUT", fmt.Sprintf("/products/%v", product.ID), buf)
		if err != nil {
			t.Fatal(err)
		}
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status code %v, got %v", http.StatusOK, rec.Code)
		}

		if !strings.Contains(rec.Body.String(), fmt.Sprint(product.Name)) {
			t.Errorf("Expected body %v, got %v", "success", rec.Body.String())
		}
	})

	t.Run("should return 400 when id is invalid", func(t *testing.T) {
		ClearProductsTable()
		req, err := http.NewRequest("PUT", "/products/invalid", nil)
		if err != nil {
			t.Fatal(err)
		}
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %v, got %v", http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("should return 400 when body is invalid", func(t *testing.T) {
		ClearProductsTable()
		product := Product{}
		err := faker.FakeData(&product)
		if err != nil {
			t.Fatal(err)
		}
		tx := database.Create(&product)
		if tx.Error != nil {
			t.Fatal(tx.Error)
		}
		req, err := http.NewRequest("PUT", fmt.Sprintf("/products/%v", product.ID), strings.NewReader("invalid body"))
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
