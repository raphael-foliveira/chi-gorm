package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/raphael-foliveira/chi-gorm/internal/config"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/http/server"
)

var (
	testServer *httptest.Server
	tClient    *testClient
	db         *database.DB
)

func TestMain(m *testing.M) {
	config := config.LoadCfg("../../.env.test")
	db = database.Db(config.DatabaseURL)
	m.Run()
	database.Close()
}

func setUp() func() {
	testApp := server.NewApp(db).CreateMainRouter()
	testServer = httptest.NewServer(testApp)
	tClient = newTestClient(testServer)
	populateTables()
	return func() {
		db.Exec("DELETE FROM orders")
		db.Exec("DELETE FROM products")
		db.Exec("DELETE FROM clients")
	}
}

func populateTables() {
	clients := [20]entities.Client{}
	products := [20]entities.Product{}
	orders := [20]entities.Order{}
	faker.FakeData(&clients)
	faker.FakeData(&products)
	faker.FakeData(&orders)

	db.Create(&clients)
	db.Create(&products)

	for i := range orders {
		orders[i].ID = 0
		orders[i].ClientID = clients[i].ID
		orders[i].Client = clients[i]
		orders[i].ProductID = products[i].ID
		orders[i].ProductID = products[i].ID
	}
	db.Create(&orders)
}
