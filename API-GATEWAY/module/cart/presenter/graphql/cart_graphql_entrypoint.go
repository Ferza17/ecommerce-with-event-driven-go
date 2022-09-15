package graphql

import (
	"net/http"

	"github.com/graphql-go/graphql"

	"github.com/Ferza17/event-driven-api-gateway/helper/response"
	"github.com/Ferza17/event-driven-api-gateway/middleware"
	"github.com/Ferza17/event-driven-api-gateway/model/schema"
	"github.com/Ferza17/event-driven-api-gateway/module/cart"
	"github.com/Ferza17/event-driven-api-gateway/utils"
)

func cartGraphqlEntrypoint(w http.ResponseWriter, r *http.Request) {
	var (
		cartSchema = middleware.GetSchemaConfigFromContext(r.Context(), utils.CartSchemaConfigContextKey)
	)
	requestBody, err := schema.ParseBody(r)
	if err != nil {
		response.Nay(w, r, http.StatusBadRequest, utils.ErrBadRequest)
		return
	}
	cartSchemaConfig, err := graphql.NewSchema(cartSchema)
	if err != nil {
		response.Nay(w, r, http.StatusInternalServerError, utils.ErrInternalServerError)
		return
	}
	result := graphql.Do(graphql.Params{
		Schema:        cartSchemaConfig,
		Context:       cart.RegisterCartUseCaseContext(r.Context()),
		RequestString: requestBody.Query,
	})
	if len(result.Errors) > 0 {
		response.Nay(w, r, http.StatusInternalServerError, utils.ErrInternalServerError)
		return
	}
	response.Yay(w, r, http.StatusOK, result)
	return
}
