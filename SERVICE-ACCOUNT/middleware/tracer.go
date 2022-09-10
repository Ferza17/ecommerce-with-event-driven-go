package middleware

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	"github.com/Ferza17/event-driven-account-service/utils"
)

func UnaryRegisterTracerContext(tracer opentracing.Tracer) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		newCtx := context.WithValue(ctx, utils.TracerContextKey, tracer)
		resp, err = handler(newCtx, req)
		return resp, err
	}
}

func RegisterTracerContext(tracer opentracing.Tracer, parentContext context.Context) (ctx context.Context) {
	ctx = context.WithValue(parentContext, utils.TracerContextKey, tracer)
	return
}

func GetTracerFromContext(ctx context.Context) opentracing.Tracer {
	return ctx.Value(utils.TracerContextKey).(opentracing.Tracer)
}
