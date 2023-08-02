package product

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/server/srverr"
	"github.com/raphael-foliveira/chi-gorm/pkg/res"
)

type Controller struct {
	repository iRepository
}

func NewController(r iRepository) *Controller {
	return &Controller{r}
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	var newProduct Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		srverr.Error(w, 400, "bad request")
		return
	}
	err = c.repository.Create(&newProduct)
	if err != nil {
		srverr.Error(w, 500, "internal server error")
		return
	}
	defer r.Body.Close()
	res.JSON(w, http.StatusCreated, &newProduct)
}

func (c *Controller) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		srverr.Error(w, 400, "bad request")
		return
	}
	product, err := c.repository.Get(id)
	if err != nil {
		srverr.Error(w, 404, "product not found")
		return
	}
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		srverr.Error(w, 400, "bad request")
		return
	}
	defer r.Body.Close()
	err = c.repository.Update(&product)
	if err != nil {
		srverr.Error(w, 500, "internal server error")
		return
	}
	res.JSON(w, http.StatusOK, &product)
}

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		srverr.Error(w, 400, "bad request")
		return
	}
	product, err := c.repository.Get(id)
	if err != nil {
		srverr.Error(w, 404, "product not found")
		return
	}
	err = c.repository.Delete(&product)
	if err != nil {
		srverr.Error(w, 500, "internal server error")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (c *Controller) List(w http.ResponseWriter, r *http.Request) {
	products, err := c.repository.List()
	if err != nil {
		srverr.Error(w, 500, "internal server error")
		return
	}
	res.JSON(w, http.StatusOK, &products)
}

func (c *Controller) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		srverr.Error(w, 400, "bad request")
		return
	}
	product, err := c.repository.Get(id)
	if err != nil {
		srverr.Error(w, 404, "product not found")
		return
	}
	res.JSON(w, http.StatusOK, &product)
}
