package grpc

import (
	"context"

	"github.com/Ferza17/event-driven-cart-service/model/pb"
)

type cartGRPCPresenter struct {
	pb.UnimplementedCartServiceServer
}

func NewCartGRPCPresenter() *cartGRPCPresenter {
	return &cartGRPCPresenter{}
}

func (receiver cartGRPCPresenter) FindCartByCartId(ctx context.Context, request *pb.FindCartByCartIdRequest) (response *pb.Cart, err error) {
	response = &pb.Cart{}
	return
}

func (receiver cartGRPCPresenter) FindCartItems(ctx context.Context, request *pb.FindCartItemsRequest) (response *pb.FindCartItemsResponse, err error) {
	response = &pb.FindCartItemsResponse{}
	return
}
