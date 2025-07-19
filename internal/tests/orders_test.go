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

func TestOrders_List(t *testing.T) {
	deps := newTestDependencies(t)
	orders := []entities.Order{}
	database.DB.Find(&orders)
	expectedBody := schemas.NewOrders(orders)
	response, err := deps.makeRequest("GET", "/orders", nil)
	assert.NoError(t, err)
	defer response.Body.Close()
	responseBody := []schemas.Order{}
	json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, responseBody[0].Quantity, expectedBody[0].Quantity)
}

func TestOrders_Get(t *testing.T) {
	deps := newTestDependencies(t)
	order := entities.Order{}
	database.DB.First(&order)
	expectedBody := schemas.NewOrder(&order)
	response, err := deps.makeRequest("GET", fmt.Sprintf("/orders/%v", order.ID), nil)
	assert.NoError(t, err)
	defer response.Body.Close()
	responseBody := schemas.Order{}
	json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, responseBody.Quantity, expectedBody.Quantity)
}

func TestOrders_Create(t *testing.T) {
	deps := newTestDependencies(t)
	product := entities.Product{}
	database.DB.First(&product)
	client := entities.Client{}
	database.DB.First(&client)
	order := schemas.CreateOrder{
		ProductID: product.ID,
		ClientID:  client.ID,
	}
	faker.FakeData(&order)
	expectedBody := schemas.Order{}
	expectedBody.Quantity = order.Quantity
	response, err := deps.makeRequest("POST", "/orders", order)
	assert.NoError(t, err)
	defer response.Body.Close()
	var responseBody entities.Order
	json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, http.StatusCreated, response.StatusCode)
	assert.Equal(t, responseBody.Quantity, expectedBody.Quantity)
}

func TestOrders_Update(t *testing.T) {
	deps := newTestDependencies(t)
	var order entities.Order
	database.DB.First(&order)
	update := schemas.UpdateOrder{}
	faker.FakeData(&update)
	expectedQuantity := update.Quantity
	var client entities.Client
	database.DB.First(&client)
	update.ClientID = client.ID
	var product entities.Product
	database.DB.First(&product)
	update.ProductID = product.ID
	response, err := deps.makeRequest("PUT", fmt.Sprintf("/orders/%v", order.ID), update)
	assert.NoError(t, err)
	defer response.Body.Close()
	var responseBody schemas.Order
	json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, responseBody.Quantity, expectedQuantity)
}

func TestOrders_Delete(t *testing.T) {
	deps := newTestDependencies(t)
	order := entities.Order{}
	database.DB.First(&order)
	response, err := deps.makeRequest("DELETE", fmt.Sprintf("/orders/%v", order.ID), nil)
	assert.NoError(t, err)
	defer response.Body.Close()
	assert.Equal(t, http.StatusNoContent, response.StatusCode)
}
