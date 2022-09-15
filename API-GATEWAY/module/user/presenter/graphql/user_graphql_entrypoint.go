package graphql

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/graphql-go/graphql"

	errorHandler "github.com/Ferza17/event-driven-api-gateway/helper/error"
	"github.com/Ferza17/event-driven-api-gateway/helper/response"
	"github.com/Ferza17/event-driven-api-gateway/middleware"
	"github.com/Ferza17/event-driven-api-gateway/model/schema"
	"github.com/Ferza17/event-driven-api-gateway/module/user"
	"github.com/Ferza17/event-driven-api-gateway/utils"
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

func userGraphqlEntrypoint(w http.ResponseWriter, r *http.Request) {
	var (
		userSchema = middleware.GetSchemaConfigFromContext(r.Context(), utils.UserSchemaConfigContextKey)
	)
	requestBody, err := schema.ParseBody(r)
	if err != nil {
		response.Nay(w, r, http.StatusBadRequest, utils.ErrBadRequest)
		return
	}
	userSchemaConfig, err := graphql.NewSchema(userSchema)
	if err != nil {
		log.Fatal(err)
	}
	result := graphql.Do(graphql.Params{
		Schema:        userSchemaConfig,
		Context:       user.RegisterUserUseCaseContext(r.Context()),
		RequestString: requestBody.Query,
	})
	if len(result.Errors) > 0 {
		response.Nay(w, r, http.StatusInternalServerError, utils.ErrInternalServerError)
		return
	}
	response.Yay(w, r, http.StatusOK, result)
	return
}

func authGraphqlEntrypoint(w http.ResponseWriter, r *http.Request) {
	var (
		authSchema = middleware.GetSchemaConfigFromContext(r.Context(), utils.AuthSchemaConfigContextKey)
	)
	requestBody, err := schema.ParseBody(r)
	if err != nil {
		response.Nay(w, r, http.StatusBadRequest, utils.ErrBadRequest)
		return
	}
	userSchemaConfig, err := graphql.NewSchema(authSchema)
	if err != nil {
		log.Fatal(err)
	}
	result := graphql.Do(graphql.Params{
		Schema:        userSchemaConfig,
		Context:       user.RegisterUserUseCaseContext(r.Context()),
		RequestString: requestBody.Query,
	})
	if len(result.Errors) > 0 {
		errorCode := errorHandler.HandleGraphQLError(result.Errors)
		response.Nay(w, r, errorCode.StatusCode, errorCode.Error)
		return
	}
	response.Yay(w, r, http.StatusOK, result)
	return
}
