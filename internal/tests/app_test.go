package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
	"github.com/raphael-foliveira/chi-gorm/internal/http/server"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
	"github.com/raphael-foliveira/chi-gorm/internal/testhelpers"
)

func TestMain(m *testing.M) {
	m.Run()
}

func initializeDependencies(t *testing.T) {
	t.Helper()

	testhelpers.StartDB()
	app := server.CreateMainRouter()
	clientsRepository := repository.NewClients(database.DB)
	ordersRepository := repository.NewOrders(database.DB)
	productsRepository := repository.NewProducts(database.DB)
	clientsController := controller.NewClients(clientsRepository, ordersRepository)
	ordersController := controller.NewOrders(ordersRepository)
	productsController := controller.NewProducts(productsRepository)

	app.Mount("/clients", clientsController.Routes())
	app.Mount("/orders", ordersController.Routes())
	app.Mount("/products", productsController.Routes())

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
