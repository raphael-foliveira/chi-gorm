package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/raphael-foliveira/chi-gorm/pkg/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/pkg/models"
)

func TestOrders(t *testing.T) {

	t.Run("Test list", func(t *testing.T) {
		setUp()
		orders := []models.Order{}
		testDb.Find(&orders)
		expectedBody := schemas.NewOrders(orders)

		response, err := makeRequest("GET", "/orders", nil)
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

		tearDown()
	})

	t.Run("Test get", func(t *testing.T) {
		setUp()
		order := models.Order{}
		testDb.First(&order)
		expectedBody := schemas.NewOrder(&order)

		response, err := makeRequest("GET", "/orders/"+fmt.Sprint(order.ID), nil)
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

		tearDown()
	})

	t.Run("Test create", func(t *testing.T) {
		setUp()
		product := models.Product{}
		testDb.First(&product)
		client := models.Client{}
		testDb.First(&client)
		order := schemas.CreateOrder{
			ProductID: product.ID,
			ClientID:  client.ID,
		}
		faker.FakeData(&order)
		expectedBody := schemas.Order{}
		expectedBody.Quantity = order.Quantity

		response, err := makeRequest("POST", "/orders", order)
		if err != nil {
			t.Error(err)
		}
		defer response.Body.Close()

		responseBody := models.Order{}
		json.NewDecoder(response.Body).Decode(&responseBody)
		if response.StatusCode != http.StatusCreated {
			t.Errorf("Expected status code %d, got %d", http.StatusCreated, response.StatusCode)
		}

		if responseBody.Quantity != expectedBody.Quantity {
			t.Errorf("Expected quantity %d, got %d", expectedBody.Quantity, responseBody.Quantity)
		}

		tearDown()
	})

	t.Run("Test update", func(t *testing.T) {
		setUp()
		order := models.Order{}
		testDb.First(&order)
		update := schemas.UpdateOrder{}
		faker.FakeData(&update)
		expectedBody := schemas.Order{}
		expectedBody.Quantity = update.Quantity

		response, err := makeRequest("PUT", "/orders/"+fmt.Sprint(order.ID), update)
		if err != nil {
			t.Error(err)
		}
		defer response.Body.Close()

		responseBody := models.Order{}
		json.NewDecoder(response.Body).Decode(&responseBody)
		if response.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
		}

		if responseBody.Quantity != expectedBody.Quantity {
			t.Errorf("Expected quantity %d, got %d", expectedBody.Quantity, responseBody.Quantity)
		}

		tearDown()
	})

	t.Run("Test delete", func(t *testing.T) {
		setUp()
		order := models.Order{}
		testDb.First(&order)

		response, err := makeRequest("DELETE", "/orders/"+fmt.Sprint(order.ID), nil)
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
