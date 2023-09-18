package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/raphael-foliveira/chi-gorm/pkg/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/pkg/models"
	"github.com/raphael-foliveira/chi-gorm/pkg/persistence/db"
)

func TestProducts(t *testing.T) {

	t.Run("Test list", func(t *testing.T) {
		setUp()
		products := []models.Product{}
		db.Db.Find(&products)
		expectedBody := schemas.NewProducts(products)

		response, err := makeRequest("GET", "/products", nil)
		if err != nil {
			t.Error(err)
		}
		defer response.Body.Close()
		responseBody := []schemas.Product{}
		json.NewDecoder(response.Body).Decode(&responseBody)

		if response.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
		}

		if responseBody[0].Name != expectedBody[0].Name {
			t.Errorf("Expected name %s, got %s", expectedBody[0].Name, responseBody[0].Name)
		}

		tearDown()
	})

	t.Run("Test get", func(t *testing.T) {
		setUp()
		product := models.Product{}
		db.Db.First(&product)
		expectedBody := schemas.NewProduct(product)

		response, err := makeRequest("GET", "/products/"+fmt.Sprint(product.ID), nil)
		if err != nil {
			t.Error(err)
		}
		defer response.Body.Close()
		responseBody := schemas.Product{}
		json.NewDecoder(response.Body).Decode(&responseBody)

		if response.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
		}

		if responseBody.Name != expectedBody.Name {
			t.Errorf("Expected name %s, got %s", expectedBody.Name, responseBody.Name)
		}

		tearDown()
	})

	t.Run("Test create", func(t *testing.T) {
		setUp()
		product := schemas.CreateProduct{}
		faker.FakeData(&product)
		expectedBody := schemas.Product{}
		expectedBody.Name = product.Name
		expectedBody.Price = product.Price

		response, err := makeRequest("POST", "/products", product)
		if err != nil {
			t.Error(err)
		}
		defer response.Body.Close()

		responseBody := models.Product{}
		json.NewDecoder(response.Body).Decode(&responseBody)
		if response.StatusCode != http.StatusCreated {
			t.Errorf("Expected status code %d, got %d", http.StatusCreated, response.StatusCode)
		}

		if responseBody.Name != expectedBody.Name {
			t.Errorf("Expected name %s, got %s", expectedBody.Name, responseBody.Name)
		}

		tearDown()
	})

	t.Run("Test update", func(t *testing.T) {
		setUp()
		product := models.Product{}
		db.Db.First(&product)
		update := schemas.UpdateProduct{}
		faker.FakeData(&update)
		expectedBody := schemas.Product{}
		expectedBody.Name = update.Name
		expectedBody.Price = update.Price

		response, err := makeRequest("PUT", "/products/"+fmt.Sprint(product.ID), update)
		if err != nil {
			t.Error(err)
		}
		defer response.Body.Close()

		responseBody := models.Product{}
		json.NewDecoder(response.Body).Decode(&responseBody)
		if response.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
		}

		if responseBody.Name != expectedBody.Name {
			t.Errorf("Expected name %s, got %s", expectedBody.Name, responseBody.Name)
		}

		tearDown()
	})

	t.Run("Test delete", func(t *testing.T) {
		setUp()
		product := models.Product{}
		db.Db.First(&product)

		response, err := makeRequest("DELETE", "/products/"+fmt.Sprint(product.ID), nil)
		if err != nil {
			t.Error(err)
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusNoContent {
			t.Errorf("Expected status code %d, got %d", http.StatusNoContent, response.StatusCode)
		}

		tearDown()
	})
}
