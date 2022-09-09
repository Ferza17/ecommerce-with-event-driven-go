package rest

import (
	"github.com/go-chi/chi/v5"
)

func routes(r *chi.Mux) {
	r.Route("/api/v1", func(r chi.Router) {
	})
}
