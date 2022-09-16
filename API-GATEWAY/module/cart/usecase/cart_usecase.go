package usecase

import (
	"context"

	errorHandler "github.com/Ferza17/event-driven-api-gateway/helper/error"
	"github.com/Ferza17/event-driven-api-gateway/helper/tracing"
	"github.com/Ferza17/event-driven-api-gateway/model/pb"
)

type cartUseCase struct {
	cartService    pb.CartServiceClient
	productService pb.ProductServiceClient
}

func NewCartUseCase(
	cartService pb.CartServiceClient,
	productService pb.ProductServiceClient,
) CartUseCaseStore {
	return &cartUseCase{
		cartService:    cartService,
		productService: productService,
	}
}

func (u *cartUseCase) FindCartByUserId(ctx context.Context, request *pb.FindCartByUserIdRequest) (response *pb.Cart, err error) {
	response = &pb.Cart{}
	span, ctx := tracing.StartSpanFromContext(ctx, "CartUseCase-FindCartByUserId")
	defer span.Finish()
	response, err = u.cartService.FindCartByUserId(ctx, request)
	if err != nil {
		err = errorHandler.HandlerGrpcError(err)
	}
	return
}

func (u *cartUseCase) FindCartItems(ctx context.Context, request *pb.FindCartItemsRequest) (response *pb.FindCartItemsResponse, err error) {
	response = &pb.FindCartItemsResponse{}
	span, ctx := tracing.StartSpanFromContext(ctx, "CartUseCase-FindCartItems")
	defer span.Finish()
	return
}
