package grpc

import (
	"context"

	errorHandler "github.com/Ferza17/event-driven-cart-service/helper/error"
	"github.com/Ferza17/event-driven-cart-service/helper/tracing"
	"github.com/Ferza17/event-driven-cart-service/model/pb"
	"github.com/Ferza17/event-driven-cart-service/module/cart"
)

type cartGRPCPresenter struct {
	pb.UnimplementedCartServiceServer
}

func NewCartGRPCPresenter() *cartGRPCPresenter {
	return &cartGRPCPresenter{}
}

func (h *cartGRPCPresenter) FindCartByUserId(ctx context.Context, request *pb.FindCartByUserIdRequest) (response *pb.Cart, err error) {
	var (
		cartUseCase = cart.GetCartUseCaseFromContext(ctx)
	)
	response = &pb.Cart{}
	span, ctx := tracing.StartSpanFromContext(ctx, "CartGRPCPresenter-FindCartByUserId")
	defer span.Finish()
	if response, err = cartUseCase.FindCartByUserId(ctx, request); err != nil {
		err = errorHandler.RpcErrorHandler(err)
	}
	return
}

func (h *cartGRPCPresenter) FindCartItems(ctx context.Context, request *pb.FindCartItemsRequest) (response *pb.FindCartItemsResponse, err error) {
	response = &pb.FindCartItemsResponse{}
	return
}
