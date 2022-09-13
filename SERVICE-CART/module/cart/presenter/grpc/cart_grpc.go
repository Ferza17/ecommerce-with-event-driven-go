package grpc

import (
	"context"

	"github.com/opentracing/opentracing-go"

	errorHandler "github.com/Ferza17/event-driven-cart-service/helper/error"
	"github.com/Ferza17/event-driven-cart-service/helper/tracing"
	"github.com/Ferza17/event-driven-cart-service/middleware"
	"github.com/Ferza17/event-driven-cart-service/model/pb"
	"github.com/Ferza17/event-driven-cart-service/module/cart"
)

type cartGRPCPresenter struct {
	pb.UnimplementedCartServiceServer
}

func NewCartGRPCPresenter() *cartGRPCPresenter {
	return &cartGRPCPresenter{}
}

func (h *cartGRPCPresenter) FindCartById(ctx context.Context, request *pb.FindCartByCartIdRequest) (response *pb.Cart, err error) {
	var (
		cartUseCase = cart.GetCartUseCaseFromContext(ctx)
		tracer      = middleware.GetTracerFromContext(ctx)
		span        = tracing.StartSpanFromRpc(tracer, "FindCartById")
	)
	response = &pb.Cart{}
	opentracing.SetGlobalTracer(tracer)
	defer span.Finish()
	ctx = opentracing.ContextWithSpan(ctx, span)
	if response, err = cartUseCase.FindCartById(ctx, request); err != nil {
		err = errorHandler.RpcError(err)
	}
	return
}

func (h *cartGRPCPresenter) FindCartItems(ctx context.Context, request *pb.FindCartItemsRequest) (response *pb.FindCartItemsResponse, err error) {
	response = &pb.FindCartItemsResponse{}
	return
}
