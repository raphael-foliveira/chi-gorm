package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
)

func TestOrders(t *testing.T) {

	testCase(t, "Test list", func(t *testing.T) {
		orders := []entities.Order{}
		db.Find(&orders)
		expectedBody := schemas.NewOrders(orders)

		response, err := tClient.makeRequest("GET", "/orders", nil)
		if err != nil {
			t.Error(err)
		}
		defer response.Body.Close()
		responseBody := []schemas.Order{}
		json.NewDecoder(response.Body).Decode(&responseBody)

		if response.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
		}

		if responseBody[0].Quantity != expectedBody[0].Quantity {
			t.Errorf("Expected quantity %d, got %d", expectedBody[0].Quantity, responseBody[0].Quantity)
		}
	})

	testCase(t, "Test get", func(t *testing.T) {
		order := entities.Order{}
		db.First(&order)
		expectedBody := schemas.NewOrder(&order)

		response, err := tClient.makeRequest("GET", "/orders/"+fmt.Sprint(order.ID), nil)
		if err != nil {
			t.Error(err)
		}
		defer response.Body.Close()
		responseBody := schemas.Order{}
		json.NewDecoder(response.Body).Decode(&responseBody)

		if response.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
		}

		if responseBody.Quantity != expectedBody.Quantity {
			t.Errorf("Expected quantity %d, got %d", expectedBody.Quantity, responseBody.Quantity)
		}
	})

	testCase(t, "Test create", func(t *testing.T) {
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
		if err != nil {
			t.Error(err)
		}
		defer response.Body.Close()

		responseBody := entities.Order{}
		json.NewDecoder(response.Body).Decode(&responseBody)
		if response.StatusCode != http.StatusCreated {
			t.Errorf("Expected status code %d, got %d", http.StatusCreated, response.StatusCode)
		}

		if responseBody.Quantity != expectedBody.Quantity {
			t.Errorf("Expected quantity %d, got %d", expectedBody.Quantity, responseBody.Quantity)
		}
	})

	testCase(t, "Test update", func(t *testing.T) {
		order := entities.Order{}
		db.First(&order)
		update := schemas.UpdateOrder{}
		faker.FakeData(&update)
		expectedBody := schemas.Order{}
		expectedBody.Quantity = update.Quantity

		response, err := tClient.makeRequest("PUT", "/orders/"+fmt.Sprint(order.ID), update)
		if err != nil {
			t.Error(err)
		}
		defer response.Body.Close()

		responseBody := entities.Order{}
		json.NewDecoder(response.Body).Decode(&responseBody)
		if response.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
		}

		if responseBody.Quantity != expectedBody.Quantity {
			t.Errorf("Expected quantity %d, got %d", expectedBody.Quantity, responseBody.Quantity)
		}
	})

	testCase(t, "Test delete", func(t *testing.T) {
		order := entities.Order{}
		db.First(&order)

		response, err := tClient.makeRequest("DELETE", "/orders/"+fmt.Sprint(order.ID), nil)
		if err != nil {
			t.Error(err)
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusNoContent {
			t.Errorf("Expected status code %d, got %d", http.StatusNoContent, response.StatusCode)
		}
	})
}
