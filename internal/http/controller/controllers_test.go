package controller_test

import (
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/config"
	"github.com/raphael-foliveira/chi-gorm/internal/mocks"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

func TestMain(m *testing.M) {
	config.LoadCfg("../../../.env.test")
	m.Run()
}

func testCase(testFunc func(*testing.T)) func(*testing.T) {
	return func(t *testing.T) {
		mocks.MountRepositoryStubs()
		service.Initialize(&service.Config{
			ClientsRepository:  mocks.ClientsRepository,
			OrdersRepository:   mocks.OrdersRepository,
			ProductsRepository: mocks.ProductsRepository,
		})
		testFunc(t)
	}
}
