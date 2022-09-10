package grpc

import (
	"context"

	"github.com/opentracing/opentracing-go"

	errorHandler "github.com/Ferza17/event-driven-account-service/helper/error"
	"github.com/Ferza17/event-driven-account-service/helper/tracing"
	"github.com/Ferza17/event-driven-account-service/middleware"
	"github.com/Ferza17/event-driven-account-service/model/pb"
	"github.com/Ferza17/event-driven-account-service/services/user"
)

type userPresenter struct {
	pb.UnimplementedUserServiceServer
}

func NewUserPresenter() *userPresenter {
	return &userPresenter{}
}

func (p *userPresenter) Login(ctx context.Context, request *pb.LoginRequest) (response *pb.LoginResponse, err error) {
	var (
		userUseCase = user.GetUserUseCaseFromContext(ctx)
		tracer      = middleware.GetTracerFromContext(ctx)
		span        = tracing.StartSpanFromRpc(tracer, "Login")
	)
	response = &pb.LoginResponse{}
	opentracing.SetGlobalTracer(tracer)
	defer span.Finish()
	ctx = opentracing.ContextWithSpan(ctx, span)
	if response, err = userUseCase.FindUserByEmailAndPassword(ctx, request); err != nil {
		err = errorHandler.RpcError(err)
	}
	return
}

func (p *userPresenter) FindUserById(ctx context.Context, request *pb.FindUserByIdRequest) (response *pb.User, err error) {
	var (
		userUseCase = user.GetUserUseCaseFromContext(ctx)
		tracer      = middleware.GetTracerFromContext(ctx)
		span        = tracing.StartSpanFromRpc(tracer, "FindUserByEmail")
	)
	response = &pb.User{}
	opentracing.SetGlobalTracer(tracer)
	defer span.Finish()
	ctx = opentracing.ContextWithSpan(ctx, span)
	if response, err = userUseCase.FindUserById(ctx, request); err != nil {
		err = errorHandler.RpcError(err)
	}
	return
}
