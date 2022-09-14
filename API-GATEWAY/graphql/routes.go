package graphql

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/graphql-go/graphql"

	"github.com/Ferza17/event-driven-api-gateway/helper/response"
	"github.com/Ferza17/event-driven-api-gateway/model/schema"
	userSchema "github.com/Ferza17/event-driven-api-gateway/model/schema/user"
	"github.com/Ferza17/event-driven-api-gateway/module/user"
	"github.com/Ferza17/event-driven-api-gateway/utils"
)

func routes(r *chi.Mux) {
	r.Get("/check", func(writer http.ResponseWriter, request *http.Request) {
		response.Yay(writer, request, http.StatusOK, "Good")
		return
	})
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/users", func(w http.ResponseWriter, r *http.Request) {
			var rBody schema.GraphqlReqBody
			if r.Body == nil {
				response.Nay(w, r, http.StatusBadRequest, utils.ErrBadRequest)
				return
			}
			err := json.NewDecoder(r.Body).Decode(&rBody)
			if err != nil {
				response.Nay(w, r, http.StatusBadRequest, utils.ErrBadRequest)
				return
			}
			userSchemaConfig, err := graphql.NewSchema(userSchema.UserSchemaConfig)
			if err != nil {
				log.Fatal(err)
			}
			result := graphql.Do(graphql.Params{
				Schema:        userSchemaConfig,
				Context:       user.RegisterUserUseCaseContext(r.Context()),
				RequestString: rBody.Query,
			})
			if len(result.Errors) > 0 {
				response.Nay(w, r, http.StatusInternalServerError, utils.ErrInternalServerError)
				return
			}
			response.Yay(w, r, http.StatusOK, result)
			return
		})
	})
}
