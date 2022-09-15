package graphql

import (
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
	"github.com/opentracing/opentracing-go"

	"github.com/Ferza17/event-driven-api-gateway/helper/tracing"
	"github.com/Ferza17/event-driven-api-gateway/middleware"
	"github.com/Ferza17/event-driven-api-gateway/model/pb"
	"github.com/Ferza17/event-driven-api-gateway/model/schema"
	"github.com/Ferza17/event-driven-api-gateway/module/user"
)

func HandleUserLogin(p graphql.ResolveParams) (response *pb.LoginResponse, err error) {
	var (
		ctx         = p.Context
		userUseCase = user.GetUserUseCaseFromContext(ctx)
		tracer      = middleware.GetTracerFromContext(ctx)
		span        = tracing.StartSpanFromRpc(tracer, "HandleUserLogin")
	)
	opentracing.SetGlobalTracer(tracer)
	defer span.Finish()
	ctx = opentracing.ContextWithSpan(ctx, span)
	response, err = userUseCase.FindUserByEmailAndPassword(
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
		identity    = middleware.GetTokenIdentityFromContext(ctx)
	)
	opentracing.SetGlobalTracer(tracer)
	defer span.Finish()
	ctx = opentracing.ContextWithSpan(ctx, span)
	response, err = userUseCase.FindUserById(
		ctx,
		&pb.FindUserByIdRequest{
			Id: identity.UserId,
		},
	)
	return
}

func HandleRegister(p graphql.ResolveParams) (response schema.CommandResponse, err error) {
	var (
		ctx         = p.Context
		userUseCase = user.GetUserUseCaseFromContext(ctx)
		tracer      = middleware.GetTracerFromContext(ctx)
		span        = tracing.StartSpanFromRpc(tracer, "HandleRegister")
	)
	opentracing.SetGlobalTracer(tracer)
	defer span.Finish()
	ctx = opentracing.ContextWithSpan(ctx, span)
	response, err = userUseCase.CreateUser(
		ctx,
		&pb.RegisterRequest{
			Username:      p.Args["username"].(string),
			Email:         p.Args["email"].(string),
			Password:      p.Args["password"].(string),
			TransactionId: uuid.NewString(),
		},
	)
	return
}
