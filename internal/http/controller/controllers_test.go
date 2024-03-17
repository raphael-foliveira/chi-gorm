package controller_test

import (
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/config"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
	"github.com/raphael-foliveira/chi-gorm/internal/mocks"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

var services *service.Services
var controllers *controller.Controllers

func TestMain(m *testing.M) {
	services = service.NewServices(mocks.Repositories, &config.Cfg{
		JwtSecret: "",
	})
	controllers = controller.NewControllers(services)
	m.Run()
}

func testCase(testFunc func(*testing.T)) func(*testing.T) {
	return func(t *testing.T) {
		setUp()
		defer tearDown()
		testFunc(t)
	}
}

func setUp() {
	mocks.Populate()
}

func tearDown() {
	mocks.ClearRepositories()
}
