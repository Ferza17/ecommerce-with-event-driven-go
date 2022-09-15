package usecase

import (
	"context"

	"github.com/Ferza17/event-driven-cart-service/helper/tracing"
	"github.com/Ferza17/event-driven-cart-service/model/pb"
)

func (u *cartUseCase) FindCartByUserId(ctx context.Context, request *pb.FindCartByUserIdRequest) (response *pb.Cart, err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "CartUseCase-FindCartByUserId")
	defer span.Finish()
	response, err = u.cartMongoDBRepository.FindCartByUserId(ctx, request.GetId())
	return
}

func (u *cartUseCase) FindCartItems(ctx context.Context, request *pb.FindCartItemsRequest) (response *pb.FindCartItemsResponse, err error) {
	return
}
