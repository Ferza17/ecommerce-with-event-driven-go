package graphql

import (
	"github.com/go-chi/chi/v5"

	"github.com/Ferza17/event-driven-api-gateway/middleware"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Use(middleware.JwtRequired)
		r.Route("/", func(r chi.Router) {
			r.Post("/", userGraphqlEntrypoint)
		})
	})
	r.Group(func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/", authGraphqlEntrypoint)
		})
	})
	return r
}
