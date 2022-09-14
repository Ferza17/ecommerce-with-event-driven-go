package grpc

import (
	"context"

	"github.com/opentracing/opentracing-go"

	errorHandler "github.com/Ferza17/event-driven-product-service/helper/error"
	"github.com/Ferza17/event-driven-product-service/helper/tracing"
	"github.com/Ferza17/event-driven-product-service/middleware"
	"github.com/Ferza17/event-driven-product-service/model/pb"
	"github.com/Ferza17/event-driven-product-service/services/product"
)

type productGRPCPresenter struct {
	pb.UnimplementedProductServiceServer
}

func NewProductGRPCPresenter() *productGRPCPresenter {
	return &productGRPCPresenter{}
}

func (h *productGRPCPresenter) FindProductById(ctx context.Context, request *pb.FindProductByIdRequest) (response *pb.Product, err error) {
	var (
		productUseCase = product.GetProductUseCaseFromContext(ctx)
		tracer         = middleware.GetTracerFromContext(ctx)
		span           = tracing.StartSpanFromRpc(tracer, "FindProductById")
	)
	response = &pb.Product{}
	opentracing.SetGlobalTracer(tracer)
	defer span.Finish()
	ctx = opentracing.ContextWithSpan(ctx, span)
	if response, err = productUseCase.FindProductById(ctx, request); err != nil {
		err = errorHandler.RpcErrorHandler(err)
	}
	return
}

func (h *productGRPCPresenter) FindProducts(ctx context.Context, request *pb.FindProductsRequest) (response *pb.FindProductsResponse, err error) {
	var (
		productUseCase = product.GetProductUseCaseFromContext(ctx)
		tracer         = middleware.GetTracerFromContext(ctx)
		span           = tracing.StartSpanFromRpc(tracer, "FindProducts")
	)
	response = &pb.FindProductsResponse{}
	opentracing.SetGlobalTracer(tracer)
	defer span.Finish()
	ctx = opentracing.ContextWithSpan(ctx, span)
	if response, err = productUseCase.FindProducts(ctx, request); err != nil {
		err = errorHandler.RpcErrorHandler(err)
	}
	return
}
