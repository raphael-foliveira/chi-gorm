package routes

import (
	"github.com/go-chi/chi/v5"
)

type router struct {
	*chi.Mux
}

func (rr *router) Get(path string, fn ControllerFunc) {
	rr.Mux.Get(path, useHandler(fn))
}

func (rr *router) Post(path string, fn ControllerFunc) {
	rr.Mux.Post(path, useHandler(fn))
}

func (rr *router) Put(path string, fn ControllerFunc) {
	rr.Mux.Put(path, useHandler(fn))
}

func (rr *router) Patch(path string, fn ControllerFunc) {
	rr.Mux.Patch(path, useHandler(fn))
}

func (rr *router) Delete(path string, fn ControllerFunc) {
	rr.Mux.Delete(path, useHandler(fn))
}
