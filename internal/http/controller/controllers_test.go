package controller_test

import (
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/mocks"
)

func TestMain(m *testing.M) {
	mocks.UseMockRepositories()
	m.Run()
}

func setUp() {
	mocks.Populate()
}

func tearDown() {
	mocks.ClearRepositories()
}
