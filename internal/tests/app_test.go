package tests

import (
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
var tClient *testClient

func TestMain(m *testing.M) {
	err := cfg.LoadCfg("../../.env.test")
	if err != nil {
		panic(err)
	}
	config := cfg.GetCfg()
	testDatabase, err = database.GetDb(config.DatabaseURL)
	if err != nil {
		panic(err)
	}
	m.Run()
}

func setUp() {
	testApp := server.NewApp(testDatabase).CreateRouter()
	testServer = httptest.NewServer(testApp)
	tClient = &testClient{testServer}
	populateTables()
}

func tearDown() {
	testDatabase.Exec("DELETE FROM orders")
	testDatabase.Exec("DELETE FROM products")
	testDatabase.Exec("DELETE FROM clients")
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
