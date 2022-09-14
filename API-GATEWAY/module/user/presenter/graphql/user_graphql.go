package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/opentracing/opentracing-go"

	"github.com/Ferza17/event-driven-api-gateway/helper/tracing"
	"github.com/Ferza17/event-driven-api-gateway/middleware"
	"github.com/Ferza17/event-driven-api-gateway/model/pb"
	"github.com/Ferza17/event-driven-api-gateway/module/user"
)

func HandleUserLogin(p graphql.ResolveParams) (response *pb.LoginResponse, err error) {
	var (
		ctx         = p.Context
		userUseCase = user.GetUserUseCaseFromContext(ctx)
		tracer      = middleware.GetTracerFromContext(ctx)
		span        = tracing.StartSpanFromRpc(tracer, "HandleUserLogin")
	)
	response = &pb.LoginResponse{}
	opentracing.SetGlobalTracer(tracer)
	defer span.Finish()
	response, err = userUseCase.Login(
		ctx,
		&pb.LoginRequest{
			Email:    p.Args["email"].(string),
			Password: p.Args["password"].(string),
		},
	)
	return
}

func HandleFindUserById(p graphql.ResolveParams) (response *pb.User, err error) {
	var (
		ctx         = p.Context
		userUseCase = user.GetUserUseCaseFromContext(ctx)
		tracer      = middleware.GetTracerFromContext(ctx)
		span        = tracing.StartSpanFromRpc(tracer, "HandleFindUserById")
	)
	response = &pb.User{}
	ctx = opentracing.ContextWithSpan(ctx, span)
	response, err = userUseCase.FindUserById(
		ctx,
		&pb.FindUserByIdRequest{
			Id: p.Args["id"].(string),
		},
	)
	return
}
