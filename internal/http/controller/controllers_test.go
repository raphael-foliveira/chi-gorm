package controller_test

import (
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/config"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
	"github.com/raphael-foliveira/chi-gorm/internal/mocks"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

var controllers *controller.Controllers

func TestMain(m *testing.M) {
	services := service.NewServices(mocks.Repositories, &config.Cfg{
		JwtSecret: "",
	})
	config.LoadCfg("../../../.env.test")
	controllers = controller.NewControllers(services)
	clientsController = controllers.ClientsController
	ordersController = controllers.OrdersController
	productsController = controllers.ProductsController
	m.Run()
}

func testCase(testFunc func(*testing.T)) func(*testing.T) {
	return func(t *testing.T) {
		mocks.MountRepositoryStubs()
		testFunc(t)
	}
}
