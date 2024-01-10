package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/raphael-foliveira/chi-gorm/internal/cfg"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/http/server"
	"gorm.io/gorm"
)

var testServer *httptest.Server
var testDatabase *gorm.DB

func TestMain(m *testing.M) {
	err := cfg.LoadCfg("../../.env.test")
	if err != nil {
		panic(err)
	}
	testDatabase, err = database.GetDb(cfg.DatabaseURL)
	if err != nil {
		panic(err)
	}
	m.Run()
}

func setUp() {
	testApp := server.NewServer(testDatabase).CreateApp()
	testServer = httptest.NewServer(testApp)
	populateTables()
}

func tearDown() {
	testDatabase.Exec("DELETE FROM orders")
	testDatabase.Exec("DELETE FROM products")
	testDatabase.Exec("DELETE FROM clients")
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
	testDatabase.Create(&clients)
	testDatabase.Create(&products)

	for i := 0; i < 10; i++ {
		orders[i].ID = 0
		orders[i].ClientID = clients[i].ID
		orders[i].ProductID = products[i].ID
	}
	testDatabase.Create(&orders)

}
