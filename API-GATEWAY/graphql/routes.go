package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"

	"github.com/Ferza17/event-driven-api-gateway/middleware"
)

func routes(r *chi.Mux, gqlServer *handler.Server) {
	r.Handle("/", playground.Handler("GraphQL playground", "/api/v1/query"))
	r.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.RegisterTokenHTTPContext)
		r.Handle("/query", gqlServer)
	})
}
