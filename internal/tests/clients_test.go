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

func TestClients_List(t *testing.T) {
	deps := newTestDependencies(t)
	clients := []entities.Client{}
	database.DB.Find(&clients)
	expectedBody := schemas.NewClients(clients)
	response, err := deps.makeRequest("GET", "/clients", nil)
	assert.NoError(t, err)
	defer response.Body.Close()
	responseBody := []schemas.Client{}
	json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, len(expectedBody), len(responseBody))
	assert.Equal(t, expectedBody[0].Name, responseBody[0].Name)
}

func TestClients_Get(t *testing.T) {
	deps := newTestDependencies(t)
	client := entities.Client{}
	database.DB.First(&client)
	expectedBody := schemas.NewClient(&client)
	response, err := deps.makeRequest("GET", "/clients/"+fmt.Sprint(client.ID), nil)
	assert.NoError(t, err)
	defer response.Body.Close()
	responseBody := schemas.Client{}
	json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, expectedBody.Name, responseBody.Name)
}

func TestClients_Create(t *testing.T) {
	deps := newTestDependencies(t)
	client := schemas.CreateClient{}
	faker.FakeData(&client)
	expectedBody := schemas.Client{}
	expectedBody.Name = client.Name
	response, err := deps.makeRequest("POST", "/clients", client)
	assert.NoError(t, err)
	defer response.Body.Close()
	responseBody := entities.Client{}
	json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, http.StatusCreated, response.StatusCode)
	assert.Equal(t, expectedBody.Name, responseBody.Name)
}

func TestClients_Update(t *testing.T) {
	deps := newTestDependencies(t)
	client := entities.Client{}
	database.DB.First(&client)
	update := schemas.UpdateClient{}
	faker.FakeData(&update)
	expectedBody := schemas.Client{}
	expectedBody.Name = update.Name
	response, err := deps.makeRequest("PUT", "/clients/"+fmt.Sprint(client.ID), update)
	assert.NoError(t, err)
	defer response.Body.Close()
	responseBody := entities.Client{}
	json.NewDecoder(response.Body).Decode(&responseBody)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, responseBody.Name, expectedBody.Name)
}

func TestClients_Delete(t *testing.T) {
	deps := newTestDependencies(t)
	client := entities.Client{}
	database.DB.First(&client)
	response, err := deps.makeRequest("DELETE", "/clients/"+fmt.Sprint(client.ID), nil)
	assert.NoError(t, err)
	defer response.Body.Close()
	assert.Equal(t, http.StatusNoContent, response.StatusCode)
}
