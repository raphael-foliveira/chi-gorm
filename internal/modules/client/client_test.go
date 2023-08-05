package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal"
	"github.com/raphael-foliveira/chi-gorm/internal/db"
)

var testDb *db.DB
var testRouter *chi.Mux

func InsertClientsHelper(qt int) {
	for i := 0; i < qt; i++ {
		client := Client{}
		err := faker.FakeData(&client)
		client.ID = 0
		if err != nil {
			panic(err)
		}
		tx := testDb.Create(&client)
		if tx.Error != nil {
			panic(tx.Error)
		}
	}
}

func ClearClientsTable() {
	tx := testDb.Delete(&Client{}, "true")
	if tx.Error != nil {
		tx.Rollback()
		panic(tx.Error)
	}
	tx.Commit()
}

func TestMain(m *testing.M) {
	testDb = db.Connect(internal.TestConfig.DatabaseURL)
	testRouter = chi.NewRouter()
	clientTestRouter, err := NewRouter(testDb)
	if err != nil {
		panic(err)
	}
	testRouter.Handle("/clients", clientTestRouter)
	ClearClientsTable()
	m.Run()

}

func TestList(t *testing.T) {
	t.Run("should return an empty list when there are no clients", func(t *testing.T) {
		ClearClientsTable()
		req, err := http.NewRequest("GET", "/clients", nil)
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

	t.Run("should return a populated list when there are clients", func(t *testing.T) {
		ClearClientsTable()
		InsertClientsHelper(10)
		req, err := http.NewRequest("GET", "/clients", nil)
		if err != nil {
			t.Fatal(err)
		}
		rec := httptest.NewRecorder()
		testRouter.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status code %v, got %v", http.StatusOK, rec.Code)
		}

		if rec.Body.String() == "[]" {
			t.Errorf("Expected body %v, got %v", "[]", rec.Body.String())
		}
	})
}

func TestGet(t *testing.T) {
	t.Run("should return 404 when client does not exist", func(t *testing.T) {
		ClearClientsTable()
		req, err := http.NewRequest("GET", "/clients/1", nil)
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

	t.Run("should return 200 when client exists", func(t *testing.T) {
		ClearClientsTable()
		client := Client{}
		err := faker.FakeData(&client)
		if err != nil {
			t.Fatal(err)
		}
		tx := testDb.Create(&client)
		if tx.Error != nil {
			t.Fatal(tx.Error)
		}
		req, err := http.NewRequest("GET", fmt.Sprintf("/clients/%v", client.ID), nil)
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
		ClearClientsTable()
		client := Client{}
		err := faker.FakeData(&client)
		if err != nil {
			t.Fatal(err)
		}
		tx := testDb.Create(&client)
		if tx.Error != nil {
			t.Fatal(tx.Error)
		}
		req, err := http.NewRequest("GET", "/clients/fff", nil)
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
		ClearClientsTable()
		req, err := http.NewRequest("POST", "/clients", strings.NewReader("invalid body"))
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
		ClearClientsTable()
		client := Client{}
		err := faker.FakeData(&client)
		if err != nil {
			t.Fatal(err)
		}
		buf := new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(client)
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest("POST", "/clients", buf)
		if err != nil {
			t.Fatal(err)
		}
		rec := httptest.NewRecorder()
		testRouter.ServeHTTP(rec, req)

		if rec.Code != http.StatusCreated {
			t.Errorf("Expected status code %v, got %v", http.StatusCreated, rec.Code)
		}

		if !strings.Contains(rec.Body.String(), client.Name) {
			t.Errorf("Expected body %v, got %v", client.Name, rec.Body.String())
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("should return 404 when client does not exist", func(t *testing.T) {
		ClearClientsTable()
		req, err := http.NewRequest("DELETE", "/clients/1", nil)
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

	t.Run("should return 204 when client exists", func(t *testing.T) {
		ClearClientsTable()
		client := Client{}
		err := faker.FakeData(&client)
		if err != nil {
			t.Fatal(err)
		}
		tx := testDb.Create(&client)
		if tx.Error != nil {
			t.Fatal(tx.Error)
		}
		req, err := http.NewRequest("DELETE", fmt.Sprintf("/clients/%v", client.ID), nil)
		if err != nil {
			t.Fatal(err)
		}
		rec := httptest.NewRecorder()
		testRouter.ServeHTTP(rec, req)

		if rec.Code != http.StatusNoContent {
			t.Errorf("Expected status code %v, got %v", http.StatusOK, rec.Code)
		}
	})

	t.Run("should return an error when id is invalid", func(t *testing.T) {
		ClearClientsTable()
		client := Client{}
		err := faker.FakeData(&client)
		if err != nil {
			t.Fatal(err)
		}
		tx := testDb.Create(&client)
		if tx.Error != nil {
			t.Fatal(tx.Error)
		}
		req, err := http.NewRequest("DELETE", "/clients/fff", nil)
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
	t.Run("should return 404 when client does not exist", func(t *testing.T) {
		ClearClientsTable()
		req, err := http.NewRequest("PUT", "/clients/1", nil)
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

	t.Run("should return 200 when client exists", func(t *testing.T) {
		ClearClientsTable()
		InsertClientsHelper(1)
		client := Client{}
		client.Name = "updated"
		testDb.First(&client)
		buf := new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(client)
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest("PUT", fmt.Sprintf("/clients/%v", client.ID), buf)
		if err != nil {
			t.Fatal(err)
		}
		rec := httptest.NewRecorder()
		testRouter.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status code %v, got %v", http.StatusOK, rec.Code)
		}

		if !strings.Contains(rec.Body.String(), fmt.Sprint(client.Name)) {
			t.Errorf("Expected body %v, got %v", "success", rec.Body.String())
		}
	})

	t.Run("should return 400 when body is invalid", func(t *testing.T) {
		ClearClientsTable()
		InsertClientsHelper(1)
		client := Client{}
		testDb.First(&client)
		buf := new(bytes.Buffer)
		buf.Write([]byte("invalid body"))
		req, err := http.NewRequest("PUT", fmt.Sprintf("/clients/%v", client.ID), buf)
		if err != nil {
			t.Fatal(err)
		}
		rec := httptest.NewRecorder()
		testRouter.ServeHTTP(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %v, got %v", http.StatusBadRequest, rec.Code)
		}

		if !strings.Contains(rec.Body.String(), "error") {
			t.Errorf("Expected body %v, got %v", "error", rec.Body.String())
		}
	})
	t.Run("should return 400 when id is invalid", func(t *testing.T) {
		ClearClientsTable()
		InsertClientsHelper(1)
		client := Client{}
		testDb.First(&client)
		buf := new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(client)
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest("PUT", "/clients/fff", buf)
		if err != nil {
			t.Fatal(err)
		}
		rec := httptest.NewRecorder()
		testRouter.ServeHTTP(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %v, got %v", http.StatusBadRequest, rec.Code)
		}

		if !strings.Contains(rec.Body.String(), "error") {
			t.Errorf("Expected body %v, got %v", "error", rec.Body.String())
		}
	})
}
