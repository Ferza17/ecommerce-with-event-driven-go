package usecase

import (
	"context"

	"github.com/Ferza17/event-driven-api-gateway/model/pb"
)

type ProductUseCaseStore interface {
	FindProductById(ctx context.Context, request *pb.FindProductByIdRequest) (response *pb.Product, err error)
	FindProducts(ctx context.Context, request *pb.FindProductsRequest) (response *pb.FindProductsResponse, err error)
}
