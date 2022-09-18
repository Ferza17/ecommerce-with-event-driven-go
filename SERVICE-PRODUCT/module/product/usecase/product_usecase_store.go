package usecase

import (
	"context"

	"github.com/Ferza17/event-driven-product-service/model/pb"
)

type ProductUseCaseCommand interface {
}

type ProductUseCaseQuery interface {
	FindProductById(ctx context.Context, request *pb.FindProductByIdRequest) (response *pb.Product, err error)
	FindProducts(ctx context.Context, request *pb.FindProductsRequest) (response *pb.FindProductsResponse, err error)
	FindProductsByProductIds(ctx context.Context, request *pb.FindProductsByProductIdsRequest) (response *pb.FindProductsByProductIdsResponse, err error)
}

type ProductUseCaseCompensate interface {
}

type ProductUseCaseStore interface {
	ProductUseCaseCompensate
	ProductUseCaseCommand
	ProductUseCaseQuery
}
