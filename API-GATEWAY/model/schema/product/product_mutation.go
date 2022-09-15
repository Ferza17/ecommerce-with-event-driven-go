package product

import (
	"github.com/graphql-go/graphql"
)

var productMutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:   "ProductMutation",
		Fields: graphql.Fields{},
	},
)
