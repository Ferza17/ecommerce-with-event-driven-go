package product

import "github.com/graphql-go/graphql"

var productType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			"id": {
				Type: graphql.String,
			},
			"name": {
				Type: graphql.String,
			},
			"description": {
				Type: graphql.String,
			},
			"uom": {
				Type: graphql.String,
			},
			"image": {
				Type: graphql.String,
			},
			"price": {
				Type: graphql.Int,
			},
			"stock": {
				Type: graphql.Int,
			},
			"createdAt": {
				Type: graphql.Int,
			},
			"updatedAt": {
				Type: graphql.Int,
			},
			"discardedAt": {
				Type: graphql.Int,
			},
		},
	},
)

var findProductsArgsType = graphql.FieldConfigArgument{
	"productIds": {
		Type: graphql.NewList(graphql.String),
	},
	"productName": {
		Type: graphql.NewList(graphql.String),
	},
	"limit": {
		Type: graphql.Int,
	},
	"page": {
		Type: graphql.Int,
	},
}

var findProductsResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "findProductsResponse",
		Fields: graphql.Fields{
			"products": {
				Type: graphql.NewList(productType),
			},
			"total": {
				Type: graphql.Int,
			},
		},
	},
)

var findProductByIdArgsType = graphql.FieldConfigArgument{
	"id": {
		Type: graphql.String,
	},
}
