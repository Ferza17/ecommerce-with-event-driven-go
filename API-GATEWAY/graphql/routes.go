package graphql

import (
	"github.com/go-chi/chi/v5"

	cartRoutes "github.com/Ferza17/event-driven-api-gateway/module/cart/presenter/graphql"
	productRoutes "github.com/Ferza17/event-driven-api-gateway/module/product/presenter/graphql"
	userRoutes "github.com/Ferza17/event-driven-api-gateway/module/user/presenter/graphql"
)

func routes(r *chi.Mux) {
	r.Route("/api/v1", func(r chi.Router) {
		r.Mount("/users", userRoutes.Routes())
		r.Mount("/carts", cartRoutes.Routes())
		r.Mount("/products", productRoutes.Routes())
	})
}
