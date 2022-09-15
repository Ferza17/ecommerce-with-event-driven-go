package grpc

import (
	"github.com/Ferza17/event-driven-product-service/model/pb"
	productModule "github.com/Ferza17/event-driven-product-service/module/product/presenter/grpc"
)

func (srv *Server) RegisterService() {
	// CreateUser Service, Service can be multiple
	pb.RegisterProductServiceServer(srv.grpcServer, productModule.NewProductGRPCPresenter())
}
