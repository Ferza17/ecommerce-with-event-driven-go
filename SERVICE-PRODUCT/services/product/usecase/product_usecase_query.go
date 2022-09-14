package usecase

import (
	"context"

	"github.com/Ferza17/event-driven-product-service/helper/tracing"
	"github.com/Ferza17/event-driven-product-service/model/pb"
)

func (u *productUseCase) FindProductById(ctx context.Context, request *pb.FindProductByIdRequest) (response *pb.Product, err error) {
	response = &pb.Product{}
	span, ctx := tracing.StartSpanFromContext(ctx, "ProductUseCase-FindProductById")
	defer span.Finish()
	response, err = u.productPostgresRepository.FindProductById(ctx, request.GetId())
	return
}

func (u *productUseCase) FindProducts(ctx context.Context, request *pb.FindProductsRequest) (response *pb.FindProductsResponse, err error) {
	response = &pb.FindProductsResponse{}
	span, ctx := tracing.StartSpanFromContext(ctx, "ProductUseCase-FindProducts")
	defer span.Finish()
	response, err = u.productPostgresRepository.FindProducts(ctx, request)
	return
}
