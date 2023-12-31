package controller

import (
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/mocks"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
)

func TestMain(m *testing.M) {
	clientsRepository := repository.Clients
	ordersRepository := repository.Orders
	productsRepository := repository.Products
	repository.Clients = mocks.ClientsStore
	repository.Orders = mocks.OrdersStore
	repository.Products = mocks.ProductsStore
	m.Run()
	repository.Clients = clientsRepository
	repository.Orders = ordersRepository
	repository.Products = productsRepository
}

func addClients(q int) {
	for i := 0; i < q; i++ {
		var client entities.Client
		var product entities.Product
		faker.FakeData(&client)
		faker.FakeData(&product)
		mocks.ProductsStore.Store = append(mocks.ProductsStore.Store, product)
		for j := 0; j < 10; j++ {
			var order entities.Order
			faker.FakeData(&order)
			order.ClientID = client.ID
			order.ProductID = product.ID
			mocks.OrdersStore.Store = append(mocks.OrdersStore.Store, order)
		}
		client.ID = uint(i + 1)
		mocks.ClientsStore.Store = append(mocks.ClientsStore.Store, client)
	}
}

func addOrders(q int) {
	for i := 0; i < q; i++ {
		var order *entities.Order
		var client *entities.Client
		var product *entities.Product
		faker.FakeData(&order)
		faker.FakeData(&client)
		faker.FakeData(&product)
		order.ID = uint(i + 1)
		client.ID = uint(i + 1)
		product.ID = uint(i + 1)
		order.ClientID = client.ID
		order.ProductID = product.ID
		mocks.OrdersStore.Store = append(mocks.OrdersStore.Store, *order)
		mocks.ClientsStore.Store = append(mocks.ClientsStore.Store, *client)
		mocks.ProductsStore.Store = append(mocks.ProductsStore.Store, *product)
	}
}

func addProducts(q int) {
	for i := 0; i < q; i++ {
		var product entities.Product
		faker.FakeData(&product)
		product.ID = uint(i + 1)
		mocks.ProductsStore.Store = append(mocks.ProductsStore.Store, product)
	}
}
