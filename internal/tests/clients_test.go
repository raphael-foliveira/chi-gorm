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
	t.Run("Test list", func(t *testing.T) {
		setUp()
		clients := []entities.Client{}
		testDatabase.Find(&clients)
		expectedBody := schemas.NewClients(clients)

		response, err := makeRequest("GET", "/clients", nil)
		if err != nil {
			t.Error(err)
		}
		defer response.Body.Close()
		responseBody := []schemas.Client{}
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
		client := entities.Client{}
		testDatabase.First(&client)
		expectedBody := schemas.NewClient(&client)

		response, err := makeRequest("GET", "/clients/"+fmt.Sprint(client.ID), nil)
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

		tearDown()
	})

	t.Run("Test create", func(t *testing.T) {
		setUp()
		client := schemas.CreateClient{}
		faker.FakeData(&client)
		expectedBody := schemas.Client{}
		expectedBody.Name = client.Name
		expectedBody.Email = client.Email

		response, err := makeRequest("POST", "/clients", client)
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

		tearDown()
	})

	t.Run("Test update", func(t *testing.T) {
		setUp()
		client := entities.Client{}
		testDatabase.First(&client)
		update := schemas.UpdateClient{}
		faker.FakeData(&update)
		expectedBody := schemas.Client{}
		expectedBody.Name = update.Name
		expectedBody.Email = update.Email

		response, err := makeRequest("PUT", "/clients/"+fmt.Sprint(client.ID), update)
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

		tearDown()
	})

	t.Run("Test delete", func(t *testing.T) {
		setUp()
		client := entities.Client{}
		testDatabase.First(&client)

		response, err := makeRequest("DELETE", "/clients/"+fmt.Sprint(client.ID), nil)
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
