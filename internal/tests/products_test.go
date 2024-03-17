package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/stretchr/testify/assert"
)

func TestProducts_List(t *testing.T) {
	tearDown := setUp()
	defer tearDown()
	products := []entities.Product{}
	db.Find(&products)
	expectedBody := schemas.NewProducts(products)
	response, err := tClient.makeRequest("GET", "/products", nil)
	assert.NoError(t, err)
	defer response.Body.Close()
	responseBody := []schemas.Product{}
	json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, responseBody[0].Name, expectedBody[0].Name)
}

func TestProduct_Get(t *testing.T) {
	tearDown := setUp()
	defer tearDown()
	product := entities.Product{}
	db.First(&product)
	expectedBody := schemas.NewProduct(&product)

	response, err := tClient.makeRequest("GET", "/products/"+fmt.Sprint(product.ID), nil)
	assert.NoError(t, err)
	defer response.Body.Close()
	responseBody := schemas.Product{}
	json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, responseBody.Name, expectedBody.Name)
}

func TestProducts_Create(t *testing.T) {
	tearDown := setUp()
	defer tearDown()
	product := schemas.CreateProduct{}
	faker.FakeData(&product)
	expectedBody := schemas.Product{}
	expectedBody.Name = product.Name
	expectedBody.Price = product.Price
	response, err := tClient.makeRequest("POST", "/products", product)
	assert.NoError(t, err)
	defer response.Body.Close()
	responseBody := entities.Product{}
	json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, http.StatusCreated, response.StatusCode)
	assert.Equal(t, responseBody.Name, expectedBody.Name)
}

func TestProducts_Update(t *testing.T) {
	tearDown := setUp()
	defer tearDown()
	product := entities.Product{}
	db.First(&product)
	update := schemas.UpdateProduct{}
	faker.FakeData(&update)
	expectedBody := schemas.Product{}
	expectedBody.Name = update.Name
	expectedBody.Price = update.Price
	response, err := tClient.makeRequest("PUT", "/products/"+fmt.Sprint(product.ID), update)
	assert.NoError(t, err)
	defer response.Body.Close()
	responseBody := entities.Product{}
	json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, responseBody.Name, expectedBody.Name)
}

func TestProducts_Delete(t *testing.T) {
	tearDown := setUp()
	defer tearDown()
	product := entities.Product{}
	db.First(&product)

	response, err := tClient.makeRequest("DELETE", "/products/"+fmt.Sprint(product.ID), nil)
	assert.NoError(t, err)
	defer response.Body.Close()
	assert.Equal(t, http.StatusNoContent, response.StatusCode)
}
