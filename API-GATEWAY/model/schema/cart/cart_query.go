package cart

import (
	"github.com/graphql-go/graphql"

	cartPresenter "github.com/Ferza17/event-driven-api-gateway/module/cart/presenter/graphql"
)

var cartQueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CartQuery",
		Fields: graphql.Fields{
			"FindCartByUserId": {
				Type: cartType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return cartPresenter.HandleFindCartByUserId(p)
				},
			},
			"FindCartItems": {
				Type: findCartItemsType,
				Args: findCartItemsArgsType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return nil, nil
				},
			},
		},
	},
)
