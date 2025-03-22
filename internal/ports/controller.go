package ports

import "github.com/go-chi/chi/v5"

type Controller interface {
	Mount(mux *chi.Mux)
}
