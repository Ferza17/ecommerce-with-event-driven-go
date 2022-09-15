package middleware

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	"github.com/Ferza17/event-driven-cart-service/utils"
)

func UnaryRegisterTracerContext(tracer opentracing.Tracer) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		opentracing.SetGlobalTracer(tracer)
		return handler(context.WithValue(ctx, utils.TracerContextKey, tracer), req)
	}
}

func RegisterTracerContext(tracer opentracing.Tracer, ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, utils.TracerContextKey, tracer)
	return ctx
}

func GetTracerFromContext(ctx context.Context) opentracing.Tracer {
	return ctx.Value(utils.TracerContextKey).(opentracing.Tracer)
}
