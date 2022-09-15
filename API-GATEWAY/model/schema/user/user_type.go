package user

import "github.com/graphql-go/graphql"

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": {
				Type: graphql.String,
			},
			"username": {
				Type: graphql.String,
			},
			"email": {
				Type: graphql.String,
			},
			"password": {
				Type: graphql.String,
			},
			"devices": {
				Type: graphql.NewList(deviceType),
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

var deviceType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Device",
		Fields: graphql.Fields{
			"id": {
				Type: graphql.String,
			},
			"deviceId": {
				Type: graphql.String,
			},
			"accessToken": {
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

// loginResponseType Section
var loginResponseType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "LoginResponse",
		Fields: graphql.Fields{
			"userId": {
				Type: graphql.String,
			},
			"token": {
				Type: graphql.String,
			},
		},
	},
)

var loginRequestArgsType = graphql.FieldConfigArgument{
	"email": {
		Type: graphql.String,
	},
	"password": {
		Type: graphql.String,
	},
}

var registerRequestType = graphql.FieldConfigArgument{
	"username": {
		Type: graphql.String,
	},
	"email": {
		Type: graphql.String,
	},
	"password": {
		Type: graphql.String,
	},
}
