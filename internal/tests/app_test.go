package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/http/server"
	"github.com/raphael-foliveira/chi-gorm/internal/models"
	"gorm.io/gorm"
)

var testServer *httptest.Server
var testDb *gorm.DB

func TestMain(m *testing.M) {
	testDb = database.GetDb()
	database.InitMemory()
	m.Run()
}

func setUp() {
	clearDatabase()
	testApp := server.CreateApp()
	testServer = httptest.NewServer(testApp)
	populateTables()
}

func clearDatabase() {
	database.CloseDb()
	database.InitMemory()
	testDb = database.GetDb()
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

	testDb.Create(&clients)
	testDb.Create(&products)
	testDb.Create(&orders)
}
