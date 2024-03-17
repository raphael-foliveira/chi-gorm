package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/exceptions"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

type Context struct {
	Response http.ResponseWriter
	Request  *http.Request
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Response: w,
		Request:  r,
	}
}

func (c *Context) SendStatus(status int) error {
	c.Response.WriteHeader(status)
	return nil
}

func (c *Context) JSON(status int, data interface{}) error {
	c.Response.Header().Set("Content-Type", "application/json")
	c.Response.WriteHeader(status)
	return json.NewEncoder(c.Response).Encode(data)
}

func (c *Context) GetUintPathParam(paramName string) (uint, error) {
	id, err := strconv.ParseUint(chi.URLParam(c.Request, paramName), 10, 64)
	if err != nil {
		return 0, exceptions.BadRequest(fmt.Sprintf("invalid %v", paramName))
	}
	return uint(id), nil
}

func (c *Context) ParseBody(v schemas.Validatable) error {
	err := json.NewDecoder(c.Request.Body).Decode(v)
	if err != nil {
		return exceptions.UnprocessableEntity("invalid body")
	}
	if err := v.Validate(); err != nil {
		return err
	}
	return nil
}

type ControllerFunc func(*Context) error

type Controllers struct {
	ClientsController  *Clients
	OrdersController   *Orders
	ProductsController *Products
}

func NewControllers(services *service.Services) *Controllers {
	return &Controllers{
		ClientsController:  NewClients(services.ClientsService),
		OrdersController:   NewOrders(services.OrdersService),
		ProductsController: NewProducts(services.ProductsService),
	}
}
