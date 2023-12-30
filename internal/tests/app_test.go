package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/raphael-foliveira/chi-gorm/internal/cfg"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/http/server"
)

var testServer *httptest.Server
var testAppServer *server.Server

func TestMain(m *testing.M) {
	err := cfg.LoadEnv("../../.env")
	if err != nil {
		panic(err)
	}
	err = database.InitDb(cfg.TestConfig.DatabaseURL)
	if err != nil {
		panic(err)
	}
	m.Run()
}

func setUp() {
	testApp := testAppServer.CreateApp()
	testServer = httptest.NewServer(testApp)
	populateTables()
}

func clearDatabase() {
	database.Db.Exec("DELETE FROM orders")
	database.Db.Exec("DELETE FROM products")
	database.Db.Exec("DELETE FROM clients")
}

func tearDown() {
	clearDatabase()
}

func makeRequest(method string, endpoint string, body interface{}) (*http.Response, error) {
	hc := &http.Client{}
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		req, err := http.NewRequest(method, testServer.URL+endpoint, bytes.NewReader(bodyBytes))
		if err != nil {
			return nil, err
		}
		return hc.Do(req)
	}
	req, err := http.NewRequest(method, testServer.URL+endpoint, nil)
	if err != nil {
		return nil, err
	}
	return hc.Do(req)
}

func populateTables() {
	clients := [10]entities.Client{}
	products := [10]entities.Product{}
	orders := [10]entities.Order{}
	faker.FakeData(&clients)
	faker.FakeData(&products)
	faker.FakeData(&orders)

	for i := 0; i < 10; i++ {
		clients[i].ID = 0
		products[i].ID = 0
		orders[i].ID = 0
	}
	database.Db.Create(&clients)
	database.Db.Create(&products)

	for i := 0; i < 10; i++ {
		orders[i].ID = 0
		orders[i].ClientID = clients[i].ID
		orders[i].ProductID = products[i].ID
	}
	database.Db.Create(&orders)

}
