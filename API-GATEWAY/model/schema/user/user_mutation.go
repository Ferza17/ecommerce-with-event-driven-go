package user

import (
	"github.com/graphql-go/graphql"

	"github.com/Ferza17/event-driven-api-gateway/model/schema"
	userPresenter "github.com/Ferza17/event-driven-api-gateway/module/user/presenter/graphql"
)

var authMutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AuthMutation",
		Fields: graphql.Fields{
			"Register": {
				Type: schema.CommandType,
				Args: registerRequestType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return userPresenter.HandleRegister(p)
				},
			},
		},
	},
)
