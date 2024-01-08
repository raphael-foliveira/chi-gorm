package controller

import (
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/mocks"
)

func TestMain(m *testing.M) {
	m.Run()
}

func addClients(q int) {
	for i := 0; i < q; i++ {
		var client entities.Client
		var product entities.Product
		faker.FakeData(&client)
		faker.FakeData(&product)
		mocks.ProductsStore.Store = append(mocks.ProductsStore.Store, product)
		client.ID = uint(i + 1)
		for j := 0; j < 10; j++ {
			addOrderToClient(&client)
		}
		mocks.ClientsStore.Store = append(mocks.ClientsStore.Store, client)
	}
}

func addOrderToClient(client *entities.Client) {
	var product entities.Product
	faker.FakeData(&product)
	mocks.ProductsStore.Store = append(mocks.ProductsStore.Store, product)
	var order entities.Order
	faker.FakeData(&order)
	order.ClientID = client.ID
	order.ProductID = product.ID
	mocks.OrdersStore.Store = append(mocks.OrdersStore.Store, order)
}

func setUp() {
	addClients(10)
}

func tearDown() {
	mocks.ClientsStore.Store = []entities.Client{}
	mocks.OrdersStore.Store = []entities.Order{}
	mocks.ProductsStore.Store = []entities.Product{}
}
