package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/raphael-foliveira/chi-gorm/internal/config"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
	"github.com/raphael-foliveira/chi-gorm/internal/http/server"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

func TestMain(m *testing.M) {
	m.Run()
}

func initializeDependencies() {
	config.Initialize("../../.env.test")
	database.Initialize(config.DatabaseURL)
	repository.Initialize()
	service.Initialize()
	controller.Initialize()

	app := server.CreateMainRouter()

	controller.Mount(app)

	testServer = httptest.NewServer(app)
}

func setUp(t *testing.T) {
	initializeDependencies()
	populateTables()
	t.Cleanup(func() {
		database.DB.Exec("DELETE FROM orders")
		database.DB.Exec("DELETE FROM products")
		database.DB.Exec("DELETE FROM clients")
		database.Close()
	})
}

func populateTables() {
	clients := [20]entities.Client{}
	products := [20]entities.Product{}
	orders := [20]entities.Order{}
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
