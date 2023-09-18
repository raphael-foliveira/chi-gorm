package controllers

import (
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/pkg/http/res"
	"github.com/raphael-foliveira/chi-gorm/pkg/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/pkg/models"
	"github.com/raphael-foliveira/chi-gorm/pkg/persistence/store"
)

type Clients struct {
	clientsStore  store.Clients
	ordersStore   store.Orders
	productsStore store.Products
}

func NewClients(clientsStore store.Clients, ordersStore store.Orders, productsStore store.Products) *Clients {
	return &Clients{clientsStore, ordersStore, productsStore}
}

func (c *Clients) Create(w http.ResponseWriter, r *http.Request) error {
	var body schemas.CreateClient
	err := parseBody(r, &body)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest, err.Error())
	}
	newClient := models.Client{
		Name:  body.Name,
		Email: body.Email,
	}
	err = c.clientsStore.Create(&newClient)
	if err != nil {
		return res.Error(w, err, http.StatusInternalServerError, "internal server error")
	}
	return res.JSON(w, http.StatusCreated, &newClient)
}

func (c *Clients) Update(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest, err.Error())
	}
	client, err := c.clientsStore.Get(id)
	if err != nil {
		return res.Error(w, err, http.StatusNotFound, "client not found")
	}
	var body schemas.UpdateClient
	err = parseBody(r, &body)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest, "bad request")
	}
	client.Name = body.Name
	client.Email = body.Email
	err = c.clientsStore.Update(client)
	if err != nil {
		return res.Error(w, err, http.StatusInternalServerError, "internal server error")
	}
	return res.JSON(w, http.StatusOK, &client)
}

func (c *Clients) Delete(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest, err.Error())
	}
	client, err := c.clientsStore.Get(id)
	if err != nil {
		return res.Error(w, err, http.StatusNotFound, "client not found")
	}
	err = c.clientsStore.Delete(client)
	if err != nil {
		return res.Error(w, err, http.StatusInternalServerError, "internal server error")
	}
	return res.SendStatus(w, http.StatusNoContent)
}

func (c *Clients) List(w http.ResponseWriter, r *http.Request) error {
	clients, err := c.clientsStore.List()
	if err != nil {
		return res.Error(w, err, http.StatusInternalServerError, "internal server error")
	}
	return res.JSON(w, http.StatusOK, schemas.NewClients(clients))
}

func (c *Clients) Get(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest, "bad request")
	}
	client, err := c.clientsStore.Get(id)
	if err != nil {
		return res.Error(w, err, http.StatusNotFound, "client not found")
	}
	orders, err := c.ordersStore.GetByClientId(client.ID)
	if err != nil {
		return res.Error(w, err, http.StatusInternalServerError, "internal server error")
	}
	clientOrders := []schemas.ClientOrder{}
	for _, o := range orders {
		product, err := c.productsStore.Get(o.ProductID)
		if err != nil {
			return res.Error(w, err, http.StatusInternalServerError, "internal server error")
		}
		clientOrders = append(clientOrders, schemas.NewClientOrder(o, *product))
	}
	return res.JSON(w, http.StatusOK, schemas.NewClientDetail(*client, clientOrders))
}
