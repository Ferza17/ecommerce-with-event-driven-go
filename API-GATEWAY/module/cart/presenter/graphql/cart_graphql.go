package graphql

import (
	"github.com/graphql-go/graphql"

	"github.com/Ferza17/event-driven-api-gateway/helper/tracing"
	"github.com/Ferza17/event-driven-api-gateway/middleware"
	"github.com/Ferza17/event-driven-api-gateway/model/pb"
	"github.com/Ferza17/event-driven-api-gateway/module/cart"
)

func HandleFindCartByUserId(p graphql.ResolveParams) (response *pb.Cart, err error) {
	var (
		ctx         = p.Context
		cartUseCase = cart.GetCartUseCaseFromContext(ctx)
		identity    = middleware.GetTokenIdentityFromContext(ctx)
	)
	span, ctx := tracing.StartSpanFromContext(ctx, "UserGRPCPresenter-HandleUserLogin")
	defer span.Finish()
	response, err = cartUseCase.FindCartByUserId(
		ctx,
		&pb.FindCartByUserIdRequest{
			Id: identity.UserId,
		},
	)
	return
}
