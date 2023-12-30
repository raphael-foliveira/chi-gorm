package tests

import (
	"bytes"
	"encoding/json"
	"math/rand"
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
	database.Db.Exec("DELETE FROM clients")
	database.Db.Exec("DELETE FROM products")
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
	clients := []entities.Client{}
	products := []entities.Product{}
	orders := []entities.Order{}

	for i := 0; i < 20; i++ {
		var c entities.Client
		faker.FakeData(&c)
		c.ID = 0
		clients = append(clients, c)
	}
	database.Db.Create(&clients)
	for i := 0; i < 20; i++ {
		var p entities.Product
		faker.FakeData(&p)
		p.ID = 0
		products = append(products, p)
	}
	database.Db.Create(&products)
	for i := 0; i < 20; i++ {
		var o entities.Order
		faker.FakeData(&o)
		o.ID = 0
		o.ClientID = clients[rand.Intn(len(clients))].ID
		o.ProductID = products[rand.Intn(len(products))].ID
		orders = append(orders, o)
	}
	database.Db.Create(&orders)

}
