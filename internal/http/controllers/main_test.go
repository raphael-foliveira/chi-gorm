package controllers

import (
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/mocks"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
)

func TestMain(m *testing.M) {
	clientsRepository := repository.Clients
	ordersRepository := repository.Orders
	productsRepository := repository.Products
	repository.Clients = mocks.ClientsStore
	repository.Orders = mocks.OrdersStore
	repository.Products = mocks.ProductsStore
	m.Run()
	repository.Clients = clientsRepository
	repository.Orders = ordersRepository
	repository.Products = productsRepository
}
