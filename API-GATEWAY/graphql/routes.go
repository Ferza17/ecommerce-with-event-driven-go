package graphql

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/Ferza17/event-driven-api-gateway/helper/response"
	cartRoutes "github.com/Ferza17/event-driven-api-gateway/module/cart/presenter/graphql"
	userRoutes "github.com/Ferza17/event-driven-api-gateway/module/user/presenter/graphql"
)

func routes(r *chi.Mux) {
	r.Get("/check", func(writer http.ResponseWriter, request *http.Request) {
		response.Yay(writer, request, http.StatusOK, "Good")
		return
	})
	r.Route("/api/v1", func(r chi.Router) {
		r.Mount("/users", userRoutes.Routes())
		r.Mount("/carts", cartRoutes.Routes())
	})
}
