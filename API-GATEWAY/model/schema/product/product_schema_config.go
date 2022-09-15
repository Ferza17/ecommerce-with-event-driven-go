package product

import "github.com/graphql-go/graphql"

var ProductSchemaConfig = graphql.SchemaConfig{
	Query: productQueryType,
	//Mutation: productMutationType,
}
