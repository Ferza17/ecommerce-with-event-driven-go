package cart

import "github.com/graphql-go/graphql"

var cartType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Cart",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"userId": &graphql.Field{
				Type: graphql.String,
			},
			"totalPrice": &graphql.Field{
				Type: graphql.Int,
			},
			"cartItems": &graphql.Field{
				Type: graphql.NewList(cartItemType),
			},
			"createdAt": &graphql.Field{
				Type: graphql.Int,
			},
			"updatedAt": &graphql.Field{
				Type: graphql.Int,
			},
			"discardedAt": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var cartItemType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CartItem",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"productId": &graphql.Field{
				Type: graphql.String,
			},
			"quantity": &graphql.Field{
				Type: graphql.Int,
			},
			"price": &graphql.Field{
				Type: graphql.Int,
			},
			"note": &graphql.Field{
				Type: graphql.String,
			},
			"createdAt": &graphql.Field{
				Type: graphql.Int,
			},
			"updatedAt": &graphql.Field{
				Type: graphql.Int,
			},
			"discardedAt": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var findCartItemsType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "findCartItemsType",
		Fields: graphql.Fields{
			"items": &graphql.Field{
				Type: graphql.NewList(cartItemType),
			},
		},
	},
)

var findCartItemsArgsType = graphql.FieldConfigArgument{
	"email": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"password": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}
