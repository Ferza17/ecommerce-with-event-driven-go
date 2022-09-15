package cart

import "github.com/graphql-go/graphql"

var CartSchemaConfig = graphql.SchemaConfig{
	Query: cartQueryType,
	//Mutation: cartMutationType,
}
