package order

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/pkg/res"
)

type Controller struct {
	repository iRepository
}

func NewController(r iRepository) *Controller {
	return &Controller{r}
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	var createOrderSchema CreateOrderSchema
	err := json.NewDecoder(r.Body).Decode(&createOrderSchema)
	if err != nil {
		res.Error(w, 400, "invalid request body")
		return
	}
	newOrder := Order{
		ClientID:  createOrderSchema.ClientID,
		ProductID: createOrderSchema.ProductID,
		Quantity:  createOrderSchema.Quantity,
	}
	err = c.repository.Create(&newOrder)
	if err != nil {
		res.Error(w, 500, "internal server error")
		return
	}
	defer r.Body.Close()
	res.JSON(w, http.StatusCreated, &newOrder)
}

func (c *Controller) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		res.Error(w, 400, "invalid id")
		return
	}
	Order, err := c.repository.Get(id)
	if err != nil {
		res.Error(w, 404, "order not found")
		return
	}
	err = json.NewDecoder(r.Body).Decode(&Order)
	if err != nil {
		res.Error(w, 400, "invalid request body")
		return
	}
	defer r.Body.Close()
	err = c.repository.Update(&Order)
	if err != nil {
		res.Error(w, 500, "internal server error")
		return
	}
	res.JSON(w, http.StatusOK, &Order)
}

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		res.Error(w, 400, "invalid id")
		return
	}
	order, err := c.repository.Get(id)
	if err != nil {
		res.Error(w, 404, "order not found")
		return
	}
	err = c.repository.Delete(&order)
	if err != nil {
		res.Error(w, 500, "internal server error")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (c *Controller) List(w http.ResponseWriter, r *http.Request) {
	orders, err := c.repository.List()
	if err != nil {
		res.Error(w, 500, "internal server error")
		return
	}
	res.JSON(w, http.StatusOK, &orders)
}

func (c *Controller) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		res.Error(w, 400, "invalid id")
		return
	}
	order, err := c.repository.Get(id)
	if err != nil {
		res.Error(w, 404, "order not found")
		return
	}
	res.JSON(w, http.StatusOK, &order)
}
