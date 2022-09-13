package usecase

import (
	"context"

	"github.com/Ferza17/event-driven-cart-service/model/pb"
)

type CartUseCaseCommand interface {
	CreateCart(ctx context.Context, request *pb.CreateCartRequest) (response *pb.CreateCartResponse, err error)
}

type CartUseCaseQuery interface {
	FindCartById(ctx context.Context, request *pb.FindCartByCartIdRequest) (response *pb.Cart, err error)
	FindCartItems(ctx context.Context, request *pb.FindCartItemsRequest) (response *pb.FindCartItemsResponse, err error)
}

type CartUseCaseCompensate interface {
	RollbackNewUserSAGA(ctx context.Context, transactionId string) (err error)
}

type CartUseCaseStore interface {
	CartUseCaseCommand
	CartUseCaseQuery
	CartUseCaseCompensate
}
