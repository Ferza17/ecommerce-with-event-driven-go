package usecase

import (
	"context"

	"github.com/Ferza17/event-driven-cart-service/helper/tracing"
	"github.com/Ferza17/event-driven-cart-service/model/pb"
)

func (u *cartUseCase) FindCartById(ctx context.Context, request *pb.FindCartByCartIdRequest) (response *pb.Cart, err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "CartUseCase-FindCartById")
	defer span.Finish()
	response, err = u.cartMongoDBRepository.FindCartById(ctx, request.GetId())
	return
}

func (u *cartUseCase) FindCartItems(ctx context.Context, request *pb.FindCartItemsRequest) (response *pb.FindCartItemsResponse, err error) {
	//TODO implement me
	panic("implement me")
}
