package cart

import "github.com/graphql-go/graphql"

var cartType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Cart",
		Fields: graphql.Fields{
			"id": {
				Type: graphql.String,
			},
			"userId": {
				Type: graphql.String,
			},
			"totalPrice": {
				Type: graphql.Int,
			},
			"cartItems": {
				Type: graphql.NewList(cartItemType),
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

var cartItemType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CartItem",
		Fields: graphql.Fields{
			"id": {
				Type: graphql.String,
			},
			"productId": {
				Type: graphql.String,
			},
			"quantity": {
				Type: graphql.Int,
			},
			"price": {
				Type: graphql.Int,
			},
			"note": {
				Type: graphql.String,
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

var findCartItemsType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "findCartItemsType",
		Fields: graphql.Fields{
			"items": {
				Type: graphql.NewList(cartItemType),
			},
		},
	},
)

var findCartItemsArgsType = graphql.FieldConfigArgument{
	"email": {
		Type: graphql.String,
	},
	"password": {
		Type: graphql.String,
	},
}
