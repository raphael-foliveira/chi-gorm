package client

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
	db  []Client
	err bool
}

func (m *MockRepository) List() ([]Client, error) {
	if m.err {
		return nil, errors.New("test")
	}
	return m.db, nil
}

func (m *MockRepository) Get(id uint64) (Client, error) {
	if m.err {
		return Client{}, errors.New("test")
	}
	for _, o := range m.db {
		if uint64(o.ID) == id {
			return o, nil
		}
	}
	return Client{}, errors.New("not found")
}

func (m *MockRepository) Create(c *Client) error {
	if m.err {
		return errors.New("test")
	}
	m.db = append(m.db, *c)
	return nil
}

func (m *MockRepository) Update(c *Client) error {
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

func (m *MockRepository) Delete(c *Client) error {
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

var repository *MockRepository
var testRouter *chi.Mux

func InsertClientsHelper(qt int) {
	for i := 0; i < qt; i++ {
		client := Client{}
		err := faker.FakeData(&client)
		if err != nil {
			panic(err)
		}
		repository.db = append(repository.db, client)
	}
}

func ClearClientsTable() {
	repository.db = []Client{}
}

func TestMain(m *testing.M) {
	testRouter = chi.NewRouter()
	repository = new(MockRepository)
	clientTestRouter := NewRouter(repository)
	testRouter.Mount("/clients", clientTestRouter)
	ClearClientsTable()
	code := m.Run()
	ClearClientsTable()
	os.Exit(code)
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
		repository.db = append(repository.db, client)

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
		InsertClientsHelper(10)
		client := repository.db[0]

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
		client := repository.db[0]
		client.Name = "updated"
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
		client := repository.db[0]

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
		client := repository.db[0]
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
