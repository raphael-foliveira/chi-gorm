package controller_test

import (
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/config"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
	"github.com/raphael-foliveira/chi-gorm/internal/mocks"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

func testCase(testFunc func(*testing.T)) func(*testing.T) {
	return func(t *testing.T) {
		config.Initialize("../../../.env.test")
		mocks.MountRepositories()
		service.Initialize()
		controller.Initialize()
		mux := chi.NewMux()
		controller.Mount(mux)
		testFunc(t)
	}
}
