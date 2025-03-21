package controller_test

import (
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
	"github.com/raphael-foliveira/chi-gorm/internal/testhelpers"
)

type testDependencies struct {
	clientsController  *controller.Clients
	ordersController   *controller.Orders
	productsController *controller.Products
	clientsStubs       []entities.Client
	productsStubs      []entities.Product
	ordersStubs        []entities.Order
}

func newTestDependencies(t *testing.T) *testDependencies {
	t.Helper()

	clientsRepository := repository.NewClients(database.DB)
	ordersRepository := repository.NewOrders(database.DB)
	productsRepository := repository.NewProducts(database.DB)
	clientsController := controller.NewClients(clientsRepository, ordersRepository)
	ordersController := controller.NewOrders(ordersRepository)
	productsController := controller.NewProducts(productsRepository)

	var (
		clientsStub  []entities.Client
		productsStub []entities.Product
		ordersStub   []entities.Order
	)

	faker.FakeData(&clientsStub)
	faker.FakeData(&productsStub)
	faker.FakeData(&ordersStub)
	for i := range clientsStub {
		clientsStub[i].Orders = []entities.Order{}
	}

	database.DB.Create(&clientsStub)
	database.DB.Create(&productsStub)
	database.DB.Create(&ordersStub)

	t.Cleanup(func() {
	})

	return &testDependencies{
		clientsController:  clientsController,
		ordersController:   ordersController,
		productsController: productsController,
		clientsStubs:       clientsStub,
		productsStubs:      productsStub,
		ordersStubs:        ordersStub,
	}
}

func testCase(testFunc func(*testing.T, *testDependencies)) func(*testing.T) {
	return func(t *testing.T) {
		deps := newTestDependencies(t)
		testFunc(t, deps)
	}
}

func TestMain(m *testing.M) {
	testhelpers.StartDB()
	m.Run()
	database.DB.Exec(`
	DELETE FROM orders;
	DELETE FROM products;
	DELETE FROM clients;
	`)
	database.Close()
}
