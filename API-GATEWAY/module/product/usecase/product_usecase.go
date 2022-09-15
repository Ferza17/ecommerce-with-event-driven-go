package usecase

import (
	"context"

	errorHandler "github.com/Ferza17/event-driven-api-gateway/helper/error"
	"github.com/Ferza17/event-driven-api-gateway/helper/tracing"
	"github.com/Ferza17/event-driven-api-gateway/model/pb"
)

type productUseCase struct {
	productService pb.ProductServiceClient
}

func NewProductUseCase(
	productService pb.ProductServiceClient,
) ProductUseCaseStore {
	return &productUseCase{
		productService: productService,
	}
}

func (u *productUseCase) FindProductById(ctx context.Context, request *pb.FindProductByIdRequest) (response *pb.Product, err error) {
	response = &pb.Product{}
	span, ctx := tracing.StartSpanFromContext(ctx, "ProductUseCase-FindProductById")
	defer span.Finish()
	response, err = u.productService.FindProductById(ctx, request)
	if err != nil {
		err = errorHandler.HandlerGrpcError(err)
	}
	return
}

func (u *productUseCase) FindProducts(ctx context.Context, request *pb.FindProductsRequest) (response *pb.FindProductsResponse, err error) {
	response = &pb.FindProductsResponse{}
	span, ctx := tracing.StartSpanFromContext(ctx, "ProductUseCase-FindProducts")
	defer span.Finish()
	response, err = u.productService.FindProducts(ctx, request)
	if err != nil {
		err = errorHandler.HandlerGrpcError(err)
	}
	return
}
