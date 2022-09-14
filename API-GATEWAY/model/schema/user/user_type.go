package user

import "github.com/graphql-go/graphql"

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"username": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"password": &graphql.Field{
				Type: graphql.String,
			},
			"devices": &graphql.Field{
				Type: graphql.NewList(deviceType),
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

var deviceType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Device",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"deviceId": &graphql.Field{
				Type: graphql.String,
			},
			"accessToken": &graphql.Field{
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

// loginResponseType Section
var loginResponseType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "LoginResponse",
		Fields: graphql.Fields{
			"userId": &graphql.Field{
				Type: graphql.String,
			},
			"token": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var loginRequestArgsType = graphql.FieldConfigArgument{
	"email": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"password": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}

// findUserByIdRequestArgsType Section
var findUserByIdRequestArgsType = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}
