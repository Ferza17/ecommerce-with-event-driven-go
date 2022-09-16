package grpc

import (
	"context"

	errorHandler "github.com/Ferza17/event-driven-account-service/helper/error"
	"github.com/Ferza17/event-driven-account-service/helper/tracing"
	"github.com/Ferza17/event-driven-account-service/model/pb"
	"github.com/Ferza17/event-driven-account-service/module/user"
)

type userGRPCPresenter struct {
	pb.UnimplementedUserServiceServer
}

func NewUserGRPCPresenter() *userGRPCPresenter {
	return &userGRPCPresenter{}
}

func (p *userGRPCPresenter) Login(ctx context.Context, request *pb.LoginRequest) (response *pb.LoginResponse, err error) {
	var (
		userUseCase = user.GetUserUseCaseFromContext(ctx)
	)
	response = &pb.LoginResponse{}
	span, ctx := tracing.StartSpanFromContext(ctx, "UserGRPCPresenter-Login")
	defer span.Finish()
	if response, err = userUseCase.FindUserByEmailAndPassword(ctx, request); err != nil {
		err = errorHandler.RpcErrorHandler(err)
	}
	return
}

func (p *userGRPCPresenter) FindUserById(ctx context.Context, request *pb.FindUserByIdRequest) (response *pb.User, err error) {
	var (
		userUseCase = user.GetUserUseCaseFromContext(ctx)
	)
	response = &pb.User{}
	span, ctx := tracing.StartSpanFromContext(ctx, "UserGRPCPresenter-FindUserById")
	defer span.Finish()
	if response, err = userUseCase.FindUserById(ctx, request); err != nil {
		err = errorHandler.RpcErrorHandler(err)
	}
	return
}

func (p *userGRPCPresenter) FindUserByEmail(ctx context.Context, request *pb.FindUserByEmailRequest) (response *pb.User, err error) {
	var (
		userUseCase = user.GetUserUseCaseFromContext(ctx)
	)
	response = &pb.User{}
	span, ctx := tracing.StartSpanFromContext(ctx, "UserGRPCPresenter-FindUserByEmail")
	defer span.Finish()
	if response, err = userUseCase.FindUserByEmail(ctx, request); err != nil {
		err = errorHandler.RpcErrorHandler(err)
	}
	return
}
