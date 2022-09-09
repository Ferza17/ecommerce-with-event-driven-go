package middleware

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

const (
	tracerConnContextKey = "tracer_connection_key"
)

func UnaryRegisterTracerContext(tracer opentracing.Tracer) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		newCtx := context.WithValue(ctx, tracerConnContextKey, tracer)
		resp, err = handler(newCtx, req)
		return resp, err
	}
}

func GetTracerAccess(ctx context.Context) opentracing.Tracer {
	return ctx.Value(tracerConnContextKey).(opentracing.Tracer)
}
