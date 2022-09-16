package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"

	"github.com/Ferza17/event-driven-api-gateway/middleware"
	"github.com/Ferza17/event-driven-api-gateway/model/graph/generated"
	"github.com/Ferza17/event-driven-api-gateway/model/graph/resolver"
)

func routes(r *chi.Mux) {
	c := generated.Config{
		Resolvers: &resolver.Resolver{},
		Directives: generated.DirectiveRoot{
			Jwt: middleware.DirectiveJwtRequired,
		},
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))
	r.Handle("/", playground.Handler("GraphQL playground", "/api/v1/query"))

	r.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.RegisterTokenHTTPContext)
		r.Handle("/query", srv)
	})
}
