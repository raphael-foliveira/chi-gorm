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
	"gorm.io/gorm"
)

var (
	testServer *httptest.Server
	tClient    *testClient
	db         *gorm.DB
)

func TestMain(m *testing.M) {
	config := config.LoadCfg("../../.env.test")
	db = database.Initialize(config.DatabaseURL)
	m.Run()
	database.Close()
}

func initializeDependencies() {
	clientsRepository := repository.NewClients(db)
	productsRepository := repository.NewProducts(db)
	ordersRepository := repository.NewOrders(db)
	clientsService := service.NewClients(clientsRepository, ordersRepository)
	productsService := service.NewProducts(productsRepository)
	ordersService := service.NewOrders(ordersRepository)
	clientsController := controller.NewClients(clientsService)
	productsController := controller.NewProducts(productsService)
	ordersController := controller.NewOrders(ordersService)

	app := server.CreateMainRouter()

	clientsController.Mount(app)
	productsController.Mount(app)
	ordersController.Mount(app)
	testServer = httptest.NewServer(app)
}

func setUp(t *testing.T) {
	initializeDependencies()
	tClient = newTestClient(testServer)
	populateTables()
	t.Cleanup(func() {
		db.Exec("DELETE FROM orders")
		db.Exec("DELETE FROM products")
		db.Exec("DELETE FROM clients")
	})
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
