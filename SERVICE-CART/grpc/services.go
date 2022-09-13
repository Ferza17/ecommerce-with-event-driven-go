package grpc

import (
	"github.com/Ferza17/event-driven-cart-service/model/pb"
	cartModule "github.com/Ferza17/event-driven-cart-service/module/cart/presenter/grpc"
)

func (srv *Server) RegisterService() {
	// CreateUser Service, Service can be multiple
	pb.RegisterCartServiceServer(srv.grpcServer, cartModule.NewCartGRPCPresenter())
}
