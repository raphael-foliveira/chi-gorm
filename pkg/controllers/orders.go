package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/pkg/interfaces"
	"github.com/raphael-foliveira/chi-gorm/pkg/models"
	"github.com/raphael-foliveira/chi-gorm/pkg/res"
	"github.com/raphael-foliveira/chi-gorm/pkg/schemas"
)

type OrdersController struct {
	repository interfaces.Repository[models.Order]
}

func NewOrdersController(r interfaces.Repository[models.Order]) *OrdersController {
	return &OrdersController{r}
}

func (c *OrdersController) Create(w http.ResponseWriter, r *http.Request) error {
	var createOrderSchema schemas.CreateOrder
	err := json.NewDecoder(r.Body).Decode(&createOrderSchema)
	if err != nil {
		return res.Error(w, 400, "invalid request body", err)
	}
	newOrder := models.Order{
		ClientID:  createOrderSchema.ClientID,
		ProductID: createOrderSchema.ProductID,
		Quantity:  createOrderSchema.Quantity,
	}
	err = c.repository.Create(&newOrder)
	if err != nil {
		return res.Error(w, 500, "internal server error", err)
	}
	defer r.Body.Close()
	return res.JSON(w, http.StatusCreated, &newOrder)
}

func (c *OrdersController) Update(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return res.Error(w, 400, "invalid id", err)

	}
	Order, err := c.repository.Get(id)
	if err != nil {
		return res.Error(w, 404, "order not found", err)

	}
	err = json.NewDecoder(r.Body).Decode(&Order)
	if err != nil {
		return res.Error(w, 400, "invalid request body", err)

	}
	defer r.Body.Close()
	err = c.repository.Update(&Order)
	if err != nil {
		return res.Error(w, 500, "internal server error", err)

	}
	return res.JSON(w, http.StatusOK, &Order)
}

func (c *OrdersController) Delete(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return res.Error(w, 400, "invalid id", err)

	}
	order, err := c.repository.Get(id)
	if err != nil {
		return res.Error(w, 404, "order not found", err)

	}
	err = c.repository.Delete(&order)
	if err != nil {
		return res.Error(w, 500, "internal server error", err)

	}
	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (c *OrdersController) List(w http.ResponseWriter, r *http.Request) error {
	orders, err := c.repository.List()
	if err != nil {
		return res.Error(w, 500, "internal server error", err)
	}
	return res.JSON(w, http.StatusOK, &orders)
}

func (c *OrdersController) Get(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return res.Error(w, 400, "invalid id", err)
	}
	order, err := c.repository.Get(id)
	if err != nil {
		return res.Error(w, 404, "order not found", err)
	}
	return res.JSON(w, http.StatusOK, &order)
}
