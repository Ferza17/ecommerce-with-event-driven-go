package user

import (
	"github.com/graphql-go/graphql"

	userPresenter "github.com/Ferza17/event-driven-api-gateway/module/user/presenter/graphql"
)

var userQueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "UserQuery",
		Fields: graphql.Fields{
			"FindUserById": {
				Type: userType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return userPresenter.HandleFindUserById(p)
				},
			},
		},
	},
)

var authQueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AuthQuery",
		Fields: graphql.Fields{
			"Login": {
				Type: loginResponseType,
				Args: loginRequestArgsType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return userPresenter.HandleUserLogin(p)
				},
			},
		},
	},
)
