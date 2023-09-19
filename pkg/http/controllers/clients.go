package controllers

import (
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/pkg/http/res"
	"github.com/raphael-foliveira/chi-gorm/pkg/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/pkg/models"
	"github.com/raphael-foliveira/chi-gorm/pkg/persistence/sqlstore"
)

type Clients struct {
	clientsStore  sqlstore.Clients
	ordersStore   sqlstore.Orders
	productsStore sqlstore.Products
}

func NewClients(clientsStore sqlstore.Clients, ordersStore sqlstore.Orders, productsStore sqlstore.Products) *Clients {
	return &Clients{clientsStore, ordersStore, productsStore}
}

func (c *Clients) Create(w http.ResponseWriter, r *http.Request) error {
	var body schemas.CreateClient
	err := parseBody(r, &body)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest, err.Error())
	}
	newClient := body.ToModel()
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
	productIds := getProductIdsFromOrders(orders)
	products, err := c.productsStore.FindMany(productIds)
	if err != nil {
		return res.Error(w, err, http.StatusInternalServerError, "internal server error")
	}
	clientOrders := []schemas.ClientOrder{}
	for _, o := range orders {
		for _, p := range products {
			if o.ProductID == p.ID {
				clientOrders = append(clientOrders, schemas.NewClientOrder(o, p))
			}
		}
	}
	return res.JSON(w, http.StatusOK, schemas.NewClientDetail(*client, clientOrders))
}

func getProductIdsFromOrders(orders []models.Order) []int64 {
	productIds := []int64{}
	for _, o := range orders {
		productIds = append(productIds, o.ProductID)
	}
	return productIds
}
