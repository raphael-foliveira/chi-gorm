//go:build integration

package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/stretchr/testify/assert"
)

func TestProducts_List(t *testing.T) {
	setUp(t)
	products := []entities.Product{}
	database.DB.Find(&products)
	expectedBody := schemas.NewProducts(products)
	response, err := makeRequest("GET", "/products", nil)
	assert.NoError(t, err)
	defer response.Body.Close()
	responseBody := []schemas.Product{}
	json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, responseBody[0].Name, expectedBody[0].Name)
}

func TestProduct_Get(t *testing.T) {
	setUp(t)
	product := entities.Product{}
	database.DB.First(&product)
	expectedBody := schemas.NewProduct(&product)
	response, err := makeRequest("GET", "/products/"+fmt.Sprint(product.ID), nil)
	assert.NoError(t, err)
	defer response.Body.Close()
	responseBody := schemas.Product{}
	json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, responseBody.Name, expectedBody.Name)
}

func TestProducts_Create(t *testing.T) {
	setUp(t)
	product := schemas.CreateProduct{}
	faker.FakeData(&product)
	expectedBody := schemas.Product{}
	expectedBody.Name = product.Name
	response, err := makeRequest("POST", "/products", product)
	assert.NoError(t, err)
	defer response.Body.Close()
	responseBody := entities.Product{}
	json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, http.StatusCreated, response.StatusCode)
	assert.Equal(t, responseBody.Name, expectedBody.Name)
}

func TestProducts_Update(t *testing.T) {
	setUp(t)
	product := entities.Product{}
	database.DB.First(&product)
	update := schemas.UpdateProduct{}
	faker.FakeData(&update)
	expectedBody := schemas.Product{}
	expectedBody.Name = update.Name
	response, err := makeRequest("PUT", "/products/"+fmt.Sprint(product.ID), update)
	assert.NoError(t, err)
	defer response.Body.Close()
	responseBody := entities.Product{}
	json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, responseBody.Name, expectedBody.Name)
}

func TestProducts_Delete(t *testing.T) {
	setUp(t)
	product := entities.Product{}
	database.DB.First(&product)
	response, err := makeRequest("DELETE", "/products/"+fmt.Sprint(product.ID), nil)
	assert.NoError(t, err)
	defer response.Body.Close()
	assert.Equal(t, http.StatusNoContent, response.StatusCode)
}
