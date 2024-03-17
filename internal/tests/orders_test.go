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

func TestOrders_List(t *testing.T) {
	tearDown := setUp()
	defer tearDown()
	orders := []entities.Order{}
	db.Find(&orders)
	expectedBody := schemas.NewOrders(orders)
	response, err := tClient.makeRequest("GET", "/orders", nil)
	assert.NoError(t, err)
	defer response.Body.Close()
	responseBody := []schemas.Order{}
	json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, responseBody[0].Quantity, expectedBody[0].Quantity)
}

func TestOrders_Get(t *testing.T) {
	tearDown := setUp()
	defer tearDown()
	order := entities.Order{}
	db.First(&order)
	expectedBody := schemas.NewOrder(&order)
	response, err := tClient.makeRequest("GET", "/orders/"+fmt.Sprint(order.ID), nil)
	assert.NoError(t, err)
	defer response.Body.Close()
	responseBody := schemas.Order{}
	json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, responseBody.Quantity, expectedBody.Quantity)
}

func TestOrders_Create(t *testing.T) {
	tearDown := setUp()
	defer tearDown()
	product := entities.Product{}
	db.First(&product)
	client := entities.Client{}
	db.First(&client)
	order := schemas.CreateOrder{
		ProductID: product.ID,
		ClientID:  client.ID,
	}
	faker.FakeData(&order)
	expectedBody := schemas.Order{}
	expectedBody.Quantity = order.Quantity
	response, err := tClient.makeRequest("POST", "/orders", order)
	assert.NoError(t, err)
	defer response.Body.Close()
	responseBody := entities.Order{}
	json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, http.StatusCreated, response.StatusCode)
	assert.Equal(t, responseBody.Quantity, expectedBody.Quantity)
}

func TestOrders_Update(t *testing.T) {
	tearDown := setUp()
	defer tearDown()
	order := entities.Order{}
	db.First(&order)
	update := schemas.UpdateOrder{}
	faker.FakeData(&update)
	expectedBody := schemas.Order{}
	expectedBody.Quantity = update.Quantity
	response, err := tClient.makeRequest("PUT", "/orders/"+fmt.Sprint(order.ID), update)
	assert.NoError(t, err)
	defer response.Body.Close()
	responseBody := entities.Order{}
	json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, responseBody.Quantity, expectedBody.Quantity)
}

func TestOrders_Delete(t *testing.T) {
	tearDown := setUp()
	defer tearDown()
	order := entities.Order{}
	db.First(&order)
	response, err := tClient.makeRequest("DELETE", "/orders/"+fmt.Sprint(order.ID), nil)
	assert.NoError(t, err)
	defer response.Body.Close()
	assert.Equal(t, http.StatusNoContent, response.StatusCode)
}
