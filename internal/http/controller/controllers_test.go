package controller_test

import (
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/config"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
	"github.com/raphael-foliveira/chi-gorm/internal/mocks"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

func TestMain(m *testing.M) {
	config.LoadCfg("../../../.env.test")
	clientsController = controller.NewClients(
		service.NewClients(
			mocks.ClientsRepository,
			mocks.OrdersRepository,
		))
	productsController = controller.NewProducts(
		service.NewProducts(
			mocks.ProductsRepository,
		))
	ordersController = controller.NewOrders(
		service.NewOrders(
			mocks.OrdersRepository,
		))
	m.Run()
}

func testCase(testFunc func(*testing.T)) func(*testing.T) {
	return func(t *testing.T) {
		mocks.MountRepositoryStubs()
		testFunc(t)
	}
}
