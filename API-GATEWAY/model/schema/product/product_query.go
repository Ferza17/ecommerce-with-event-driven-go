package product

import (
	"github.com/graphql-go/graphql"

	productPresenter "github.com/Ferza17/event-driven-api-gateway/module/product/presenter/graphql"
)

var productQueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ProductQuery",
		Fields: graphql.Fields{
			"FindProductById": {
				Type: productType,
				Args: findProductByIdArgsType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return productPresenter.HandleFindProductById(p)
				},
			},
			"FindProducts": {
				Type: findProductsResponse,
				Args: findProductsArgsType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return productPresenter.HandleFindProducts(p)
				},
			},
		},
	},
)
