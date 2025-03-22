package tests

import (
	"net/http"
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

type testDependencies struct {
	clientsController  *controller.Clients
	ordersController   *controller.Orders
	productsController *controller.Products
	clientsStubs       []entities.Client
	productsStubs      []entities.Product
	ordersStubs        []entities.Order
	testServer         *httptest.Server
	makeRequest        func(method, path string, body interface{}) (*http.Response, error)
}

func newTestDependencies(t *testing.T) *testDependencies {
	t.Helper()

	testhelpers.StartDB()
	app := server.New()
	clientsRepository := repository.NewClients(database.DB)
	ordersRepository := repository.NewOrders(database.DB)
	productsRepository := repository.NewProducts(database.DB)
	clientsController := controller.NewClients(clientsRepository, ordersRepository)
	ordersController := controller.NewOrders(ordersRepository)
	productsController := controller.NewProducts(productsRepository)

	app.Mount(clientsController, ordersController, productsController)

	clients := []entities.Client{}
	products := []entities.Product{}
	orders := []entities.Order{}
	faker.FakeData(&clients)
	faker.FakeData(&products)
	faker.FakeData(&orders)

	database.DB.Create(&clients)
	database.DB.Create(&products)

	for i := range orders {
		if i >= len(clients) || i >= len(products) {
			break
		}
		orders[i].ID = 0
		orders[i].ClientID = clients[i].ID
		orders[i].Client = clients[i]
		orders[i].ProductID = products[i].ID
		orders[i].ProductID = products[i].ID
	}
	database.DB.Create(&orders)

	t.Cleanup(func() {
	})

	testServer := httptest.NewServer(app)
	return &testDependencies{
		clientsController:  clientsController,
		ordersController:   ordersController,
		productsController: productsController,
		clientsStubs:       clients,
		productsStubs:      products,
		ordersStubs:        orders,
		testServer:         testServer,
		makeRequest:        makeRequest(testServer.URL),
	}
}
