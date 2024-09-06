package controller_test

import (
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/config"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
	"github.com/raphael-foliveira/chi-gorm/internal/mocks"
)

func TestMain(m *testing.M) {
	config.LoadCfg("../../../.env.test")
	clientsController = controller.NewClients()
	productsController = controller.NewProducts()
	ordersController = controller.NewOrders()
	m.Run()
}

func testCase(testFunc func(*testing.T)) func(*testing.T) {
	return func(t *testing.T) {
		mocks.MountRepositoryStubs()
		testFunc(t)
	}
}
