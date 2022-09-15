package graphql

import (
	"github.com/go-chi/chi/v5"

	"github.com/Ferza17/event-driven-api-gateway/middleware"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.JwtRequired)
	r.Post("/", cartGraphqlEntrypoint)
	return r
}
