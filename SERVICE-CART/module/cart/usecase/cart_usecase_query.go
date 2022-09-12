package usecase

import (
	"context"

	"github.com/Ferza17/event-driven-cart-service/model/pb"
)

func (u *cartUseCase) FindCartByCartId(ctx context.Context, request *pb.FindCartByCartIdRequest) (response *pb.Cart, err error) {
	//TODO implement me
	panic("implement me")
}

func (u *cartUseCase) FindCartItems(ctx context.Context, request *pb.FindCartItemsRequest) (response *pb.FindCartItemsResponse, err error) {
	//TODO implement me
	panic("implement me")
}
