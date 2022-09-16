package graphql

import (
	"context"

	"github.com/Ferza17/event-driven-api-gateway/helper/tracing"
	"github.com/Ferza17/event-driven-api-gateway/middleware"
	"github.com/Ferza17/event-driven-api-gateway/model/graph/model"
	"github.com/Ferza17/event-driven-api-gateway/model/pb"
	"github.com/Ferza17/event-driven-api-gateway/module/cart"
)

func HandleFindCartByUserID(ctx context.Context) (response *model.Cart, err error) {
	var (
		cartUseCase = cart.GetCartUseCaseFromContext(ctx)
		identity    = middleware.GetTokenIdentityFromContext(ctx)
	)
	span, ctx := tracing.StartSpanFromContext(ctx, "UserGRPCPresenter-HandleUserLogin")
	defer span.Finish()
	cart, err := cartUseCase.FindCartByUserId(
		ctx,
		&pb.FindCartByUserIdRequest{
			Id: identity.UserId,
		},
	)
	if err != nil {
		return
	}
	response = &model.Cart{
		ID:          cart.GetId(),
		UserID:      cart.GetUserId(),
		TotalPrice:  int(cart.GetTotalPrice()),
		CartItems:   nil,
		CreatedAt:   int(cart.GetCreatedAt()),
		UpdatedAt:   int(cart.GetUpdatedAt()),
		DiscardedAt: int(cart.GetDiscardedAt()),
	}
	return
}
