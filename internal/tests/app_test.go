//go:build integration

package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/domain"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
	"github.com/raphael-foliveira/chi-gorm/internal/http/server"
	"github.com/raphael-foliveira/chi-gorm/internal/testhelpers"
)

func TestMain(m *testing.M) {
	m.Run()
}

func initializeDependencies(t *testing.T) {
	t.Helper()

	testhelpers.StartDB(t)
	app := server.CreateMainRouter()
	controller.Mount(app)
	testServer = httptest.NewServer(app)
}

func setUp(t *testing.T) {
	initializeDependencies(t)
	populateTables(t)
	t.Cleanup(func() {
		database.DB.Exec("DELETE FROM orders")
		database.DB.Exec("DELETE FROM products")
		database.DB.Exec("DELETE FROM clients")
		database.Close()
	})
}

func populateTables(t *testing.T) {
	t.Helper()
	clients := [20]domain.Client{}
	products := [20]domain.Product{}
	orders := [20]domain.Order{}
	faker.FakeData(&clients)
	faker.FakeData(&products)
	faker.FakeData(&orders)

	database.DB.Create(&clients)
	database.DB.Create(&products)

	for i := range orders {
		orders[i].ID = 0
		orders[i].ClientID = clients[i].ID
		orders[i].Client = clients[i]
		orders[i].ProductID = products[i].ID
		orders[i].ProductID = products[i].ID
	}
	database.DB.Create(&orders)
}
