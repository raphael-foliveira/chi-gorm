package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/raphael-foliveira/chi-gorm/internal/cfg"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/http/server"
)

var testServer *httptest.Server
var tClient *testClient

func TestMain(m *testing.M) {
	err := cfg.LoadCfg("../../.env.test")
	if err != nil {
		panic(err)
	}
	m.Run()
}

func setUp() {
	testApp := server.App().CreateMainRouter()
	testServer = httptest.NewServer(testApp)
	tClient = newTestClient(testServer)
	populateTables()
}

func tearDown() {
	database.Db().Exec("DELETE FROM orders")
	database.Db().Exec("DELETE FROM products")
	database.Db().Exec("DELETE FROM clients")
}

func populateTables() {
	clients := [20]entities.Client{}
	products := [20]entities.Product{}
	orders := [20]entities.Order{}
	faker.FakeData(&clients)
	faker.FakeData(&products)
	faker.FakeData(&orders)

	database.Db().Create(&clients)
	database.Db().Create(&products)

	for i := range orders {
		orders[i].ID = 0
		orders[i].ClientID = clients[i].ID
		orders[i].Client = clients[i]
		orders[i].ProductID = products[i].ID
		orders[i].ProductID = products[i].ID
	}
	database.Db().Create(&orders)
}
