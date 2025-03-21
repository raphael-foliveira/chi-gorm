//go:build integration

package repository_test

import (
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
)

type testDependencies struct {
	clientsRepo  *repository.Clients
	ordersRepo   *repository.Orders
	productsRepo *repository.Products
}

func newTestDependencies(t *testing.T) *testDependencies {
	t.Helper()

	clientsRepo := repository.NewClients(database.DB)
	ordersRepo := repository.NewOrders(database.DB)
	productsRepo := repository.NewProducts(database.DB)
	return &testDependencies{
		clientsRepo:  clientsRepo,
		ordersRepo:   ordersRepo,
		productsRepo: productsRepo,
	}
}

func TestMain(m *testing.M) {
	m.Run()
}
