package user

import (
	"github.com/graphql-go/graphql"
)

var UserSchemaConfig = graphql.SchemaConfig{
	Query: userQueryType,
}

var AuthSchemaConfig = graphql.SchemaConfig{
	Query:    authQueryType,
	Mutation: authMutationType,
}
