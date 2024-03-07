package routes

import (
	"github.com/go-chi/chi/v5"
)

type router struct {
	*chi.Mux
}

func (r *router) Get(path string, fn ControllerFunc) {
	r.Mux.Get(path, useHandler(fn))
}

func (r *router) Post(path string, fn ControllerFunc) {
	r.Mux.Post(path, useHandler(fn))
}

func (r *router) Put(path string, fn ControllerFunc) {
	r.Mux.Put(path, useHandler(fn))
}

func (r *router) Patch(path string, fn ControllerFunc) {
	r.Mux.Patch(path, useHandler(fn))
}

func (r *router) Delete(path string, fn ControllerFunc) {
	r.Mux.Delete(path, useHandler(fn))
}
