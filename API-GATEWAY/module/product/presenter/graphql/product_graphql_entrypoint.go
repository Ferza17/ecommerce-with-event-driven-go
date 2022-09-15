package graphql

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/graphql-go/graphql"

	"github.com/Ferza17/event-driven-api-gateway/helper/response"
	"github.com/Ferza17/event-driven-api-gateway/middleware"
	"github.com/Ferza17/event-driven-api-gateway/model/schema"
	"github.com/Ferza17/event-driven-api-gateway/module/product"
	"github.com/Ferza17/event-driven-api-gateway/utils"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.JwtRequired)
	r.Post("/", productGraphqlEntrypoint)
	return r
}

func productGraphqlEntrypoint(w http.ResponseWriter, r *http.Request) {
	var (
		productSchema = middleware.GetSchemaConfigFromContext(r.Context(), utils.ProductSchemaConfigContextKey)
	)
	requestBody, err := schema.ParseBody(r)
	if err != nil {
		response.Nay(w, r, http.StatusBadRequest, utils.ErrBadRequest)
		return
	}
	productSchemaConfig, err := graphql.NewSchema(productSchema)
	if err != nil {
		response.Nay(w, r, http.StatusInternalServerError, utils.ErrInternalServerError)
		return
	}
	result := graphql.Do(graphql.Params{
		Schema:        productSchemaConfig,
		Context:       product.RegisterProductUseCaseContext(r.Context()),
		RequestString: requestBody.Query,
	})
	if len(result.Errors) > 0 {
		response.Nay(w, r, http.StatusInternalServerError, utils.ErrInternalServerError)
		return
	}
	response.Yay(w, r, http.StatusOK, result)
	return
}
