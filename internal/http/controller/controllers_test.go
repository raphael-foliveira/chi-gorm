package controller_test

import (
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/config"
	"github.com/raphael-foliveira/chi-gorm/internal/container"
	"github.com/raphael-foliveira/chi-gorm/internal/mocks"
)

func TestMain(m *testing.M) {
	mocks.UseMockRepositories()
	config.LoadCfg("../../.env.test")
	clientsController = container.ClientsController()
	ordersController = container.OrdersController()
	productsController = container.ProductsController()
	m.Run()
}

func testCase(testFunc func(*testing.T)) func(*testing.T) {
	return func(t *testing.T) {
		mocks.MountRepositoryStubs()
		testFunc(t)
	}
}
