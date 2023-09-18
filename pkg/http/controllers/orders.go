package controllers

import (
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/pkg/http/res"
	"github.com/raphael-foliveira/chi-gorm/pkg/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/pkg/models"
	"github.com/raphael-foliveira/chi-gorm/pkg/persistence/store"
)

type Orders struct {
	ordersStore   store.Orders
	clientsStore  store.Clients
	productsStore store.Products
}

func NewOrders(ordersStore store.Orders, clientsStore store.Clients, productsStore store.Products) *Orders {
	return &Orders{ordersStore, clientsStore, productsStore}
}

func (c *Orders) Create(w http.ResponseWriter, r *http.Request) error {
	var body schemas.CreateOrder
	err := parseBody(r, &body)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest, "bad request")
	}
	newOrder := models.Order{
		ClientID:  body.ClientID,
		ProductID: body.ProductID,
		Quantity:  body.Quantity,
	}
	err = c.ordersStore.Create(&newOrder)
	if err != nil {
		return res.Error(w, err, http.StatusInternalServerError, "internal server error")
	}
	return res.JSON(w, http.StatusCreated, schemas.NewOrder(newOrder))
}

func (c *Orders) Update(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest, err.Error())
	}
	order, err := c.ordersStore.Get(id)
	if err != nil {
		return res.Error(w, err, http.StatusNotFound, "order not found")
	}
	var body schemas.UpdateOrder
	err = parseBody(r, &body)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest, "bad request")
	}
	order.Quantity = body.Quantity
	err = c.ordersStore.Update(order)
	if err != nil {
		return res.Error(w, err, http.StatusInternalServerError, "internal server error")
	}
	return res.JSON(w, http.StatusOK, schemas.NewOrder(*order))
}

func (c *Orders) Delete(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest, err.Error())

	}
	order, err := c.ordersStore.Get(id)
	if err != nil {
		return res.Error(w, err, http.StatusNotFound, "order not found")

	}
	err = c.ordersStore.Delete(order)
	if err != nil {
		return res.Error(w, err, http.StatusInternalServerError, "internal server error")

	}
	return res.SendStatus(w, http.StatusNoContent)
}

func (c *Orders) List(w http.ResponseWriter, r *http.Request) error {
	orders, err := c.ordersStore.List()
	if err != nil {
		return res.Error(w, err, http.StatusInternalServerError, "internal server error")
	}
	return res.JSON(w, http.StatusOK, schemas.NewOrders(orders))
}

func (c *Orders) Get(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest, err.Error())
	}
	order, err := c.ordersStore.Get(id)
	if err != nil {
		return res.Error(w, err, http.StatusNotFound, "order not found")
	}
	client, err := c.clientsStore.Get(order.ClientID)
	if err != nil {
		return res.Error(w, err, http.StatusNotFound, "client not found")
	}
	product, err := c.productsStore.Get(order.ProductID)
	if err != nil {
		return res.Error(w, err, http.StatusNotFound, "product not found")
	}
	return res.JSON(w, http.StatusOK, schemas.NewOrderDetail(*order, *client, *product))
}
