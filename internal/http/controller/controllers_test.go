package controller_test

import (
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/mocks"
)

func TestMain(m *testing.M) {
	mocks.UseMockRepositories()
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
