package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/raphael-foliveira/chi-gorm/internal/config"
	"github.com/raphael-foliveira/chi-gorm/internal/container"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"gorm.io/gorm"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	m.Run()
	database.Close()
}

func initializeDependencies(t *testing.T) {
	t.Helper()
	config := config.LoadCfg("../../.env.test")
	app := container.InitializeDependencies(config)
	db = database.DB
	testServer = httptest.NewServer(app)
}

func setUp(t *testing.T) {
	t.Helper()
	initializeDependencies(t)
	populateTables(t)
	t.Cleanup(func() {
		db.Exec("DELETE FROM orders")
		db.Exec("DELETE FROM products")
		db.Exec("DELETE FROM clients")
		database.Close()
	})
}

func populateTables(t *testing.T) {
	t.Helper()
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
