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

func TestClients(t *testing.T) {
	t.Run("Test list", testCase(func(t *testing.T) {
		clients := []entities.Client{}
		db.Find(&clients)
		expectedBody := schemas.NewClients(clients)

		response, err := tClient.makeRequest("GET", "/clients", nil)
		if err != nil {
			t.Error(err)
		}
		defer response.Body.Close()
		responseBody := []schemas.Client{}
		json.NewDecoder(response.Body).Decode(&responseBody)

		if response.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
		}

		if len(responseBody) != len(expectedBody) {
			t.Errorf("Expected %d clients, got %d", len(expectedBody), len(responseBody))
		}

		if responseBody[0].Name != expectedBody[0].Name {
			t.Errorf("Expected name %s, got %s", expectedBody[0].Name, responseBody[0].Name)
		}
	}))

	t.Run("Test get", testCase(func(t *testing.T) {
		client := entities.Client{}
		db.First(&client)
		expectedBody := schemas.NewClient(&client)

		response, err := tClient.makeRequest("GET", "/clients/"+fmt.Sprint(client.ID), nil)
		if err != nil {
			t.Error(err)
		}
		defer response.Body.Close()
		responseBody := schemas.Client{}
		json.NewDecoder(response.Body).Decode(&responseBody)

		if response.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
		}

		if responseBody.Name != expectedBody.Name {
			t.Errorf("Expected name %s, got %s", expectedBody.Name, responseBody.Name)
		}
	}))

	t.Run("Test create", testCase(func(t *testing.T) {
		client := schemas.CreateClient{}
		faker.FakeData(&client)
		expectedBody := schemas.Client{}
		expectedBody.Name = client.Name
		expectedBody.Email = client.Email

		response, err := tClient.makeRequest("POST", "/clients", client)
		if err != nil {
			t.Error(err)
		}
		defer response.Body.Close()

		responseBody := entities.Client{}
		json.NewDecoder(response.Body).Decode(&responseBody)
		if response.StatusCode != http.StatusCreated {
			t.Errorf("Expected status code %d, got %d", http.StatusCreated, response.StatusCode)
		}

		if responseBody.Name != expectedBody.Name {
			t.Errorf("Expected name %s, got %s", expectedBody.Name, responseBody.Name)
		}
	}))

	t.Run("Test update", testCase(func(t *testing.T) {
		client := entities.Client{}
		db.First(&client)
		update := schemas.UpdateClient{}
		faker.FakeData(&update)
		expectedBody := schemas.Client{}
		expectedBody.Name = update.Name
		expectedBody.Email = update.Email

		response, err := tClient.makeRequest("PUT", "/clients/"+fmt.Sprint(client.ID), update)
		if err != nil {
			t.Error(err)
		}
		defer response.Body.Close()

		responseBody := entities.Client{}
		json.NewDecoder(response.Body).Decode(&responseBody)
		if response.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
		}

		if responseBody.Name != expectedBody.Name {
			t.Errorf("Expected name %s, got %s", expectedBody.Name, responseBody.Name)
		}
	}))

	t.Run("Test delete", testCase(func(t *testing.T) {
		client := entities.Client{}
		db.First(&client)

		response, err := tClient.makeRequest("DELETE", "/clients/"+fmt.Sprint(client.ID), nil)
		if err != nil {
			t.Error(err)
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusNoContent {
			t.Errorf("Expected status code %d, got %d", http.StatusNoContent, response.StatusCode)
		}
	}))
}
