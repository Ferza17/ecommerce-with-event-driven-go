package usecase

import (
	"context"

	"github.com/Ferza17/event-driven-api-gateway/model/pb"
)

type CartUseCaseStore interface {
	FindCartByUserId(ctx context.Context, request *pb.FindCartByUserIdRequest) (response *pb.Cart, err error)
	FindCartItems(ctx context.Context, request *pb.FindCartItemsRequest) (response *pb.FindCartItemsResponse, err error)
}
