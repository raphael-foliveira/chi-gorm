package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/raphael-foliveira/chi-gorm/pkg/http/server"
	"github.com/raphael-foliveira/chi-gorm/pkg/models"
	"github.com/raphael-foliveira/chi-gorm/pkg/persistence/db"
)

var testServer *httptest.Server

func TestMain(m *testing.M) {
	db.InitMemory()
	m.Run()
}

func setUp() {
	clearDatabase()
	testApp := server.CreateApp()
	testServer = httptest.NewServer(testApp)
	populateTables()
}

func clearDatabase() {
	sqlDb, _ := db.Db.DB()
	sqlDb.Close()
	db.Db = nil
	db.InitMemory()
}

func tearDown() {
	testServer.Close()
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
	clients := []models.Client{}
	products := []models.Product{}
	orders := []models.Order{}

	for i := 0; i < 20; i++ {
		var c models.Client
		faker.FakeData(&c)
		clients = append(clients, c)
	}
	for i := 0; i < 20; i++ {
		var p models.Product
		faker.FakeData(&p)
		products = append(products, p)
	}
	for i := 0; i < 20; i++ {
		var o models.Order
		faker.FakeData(&o)
		o.ClientID = int64(i + 1)
		o.ProductID = int64(i + 1)
		orders = append(orders, o)
	}

	db.Db.Create(&clients)
	db.Db.Create(&products)
	db.Db.Create(&orders)
}
