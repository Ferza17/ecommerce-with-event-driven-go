package usecase

import (
	"context"

	errorHandler "github.com/Ferza17/event-driven-api-gateway/helper/error"
	"github.com/Ferza17/event-driven-api-gateway/helper/tracing"
	"github.com/Ferza17/event-driven-api-gateway/model/pb"
)

type cartUseCase struct {
	cartService pb.CartServiceClient
}

func NewCartUseCase(
	cartService pb.CartServiceClient,
) CartUseCaseStore {
	return &cartUseCase{
		cartService: cartService,
	}
}

func (u *cartUseCase) FindCartByUserId(ctx context.Context, request *pb.FindCartByUserIdRequest) (response *pb.Cart, err error) {
	response = &pb.Cart{}
	span, ctx := tracing.StartSpanFromContext(ctx, "CartUseCase-FindCartByUserId")
	defer span.Finish()
	response, err = u.cartService.FindCartByUserId(ctx, request)
	if err != nil {
		err = errorHandler.HandlerGrpcError(err)
		return
	}
	return
}

func (u *cartUseCase) FindCartItems(ctx context.Context, request *pb.FindCartItemsRequest) (response *pb.FindCartItemsResponse, err error) {
	response = &pb.FindCartItemsResponse{}
	span, ctx := tracing.StartSpanFromContext(ctx, "CartUseCase-FindCartItems")
	defer span.Finish()
	return
}
