package cart

import (
	"log"

	"github.com/graphql-go/graphql"

	"github.com/Ferza17/event-driven-api-gateway/model/pb"
)

var cartSubscriptionType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "cartSubscriptionType",
		Fields: graphql.Fields{
			"SubscribeCartNewStateByCartId": {
				Type: cartType,
				Args: graphql.FieldConfigArgument{
					"id": {
						Type: graphql.String,
					},
				},
				Subscribe: func(p graphql.ResolveParams) (interface{}, error) {
					log.Println(p.Args["id"].(string))
					response := &pb.Cart{
						Id:     "test",
						UserId: "test2",
					}
					return response, nil
				},
			},
		},
	},
)
