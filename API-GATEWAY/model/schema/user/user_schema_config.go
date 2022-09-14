package user

import (
	"github.com/graphql-go/graphql"
)

var UserSchemaConfig = graphql.SchemaConfig{
	Query: userQueryType,
}
