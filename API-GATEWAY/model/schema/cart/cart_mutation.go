package cart

import (
	"github.com/graphql-go/graphql"
)

var cartMutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:   "CartMutation",
		Fields: graphql.Fields{},
	},
)
