package controller

import (
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

type Controller interface {
	Create(w http.ResponseWriter, r *http.Request) error
	Update(w http.ResponseWriter, r *http.Request) error
	Delete(w http.ResponseWriter, r *http.Request) error
	List(w http.ResponseWriter, r *http.Request) error
	Get(w http.ResponseWriter, r *http.Request) error
}

type Controllers struct {
	Clients  Controller
	Orders   Controller
	Products Controller
}

func NewControllers(services *service.Services) *Controllers {
	return &Controllers{
		Clients:  NewClients(services.Clients),
		Orders:   NewOrders(services.Orders),
		Products: NewProducts(services.Products),
	}
}
