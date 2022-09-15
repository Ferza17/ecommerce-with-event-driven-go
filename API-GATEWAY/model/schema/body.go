package schema

import (
	"encoding/json"
	"net/http"

	"github.com/RoseRocket/xerrs"
	"github.com/graphql-go/graphql"

	"github.com/Ferza17/event-driven-api-gateway/utils"
)

type (
	GraphqlReqBody struct {
		Query string `json:"query"`
	}
	CommandResponse struct {
		Message commandMessageStatus `json:"message"`
	}

	commandMessageStatus string
)

var CommandType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "commandType",
		Fields: graphql.Fields{
			"message": {
				Type: graphql.String,
			},
		},
	},
)

const (
	CommandSuccess commandMessageStatus = "SUCCESS"
	CommandFailed  commandMessageStatus = "FAILED"
)

func ParseBody(r *http.Request) (body GraphqlReqBody, err error) {
	if r.Body == nil {
		err = xerrs.Mask(err, utils.ErrBadRequest)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		err = xerrs.Mask(err, utils.ErrBadRequest)
	}
	return
}
