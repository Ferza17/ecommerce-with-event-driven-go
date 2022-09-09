package middleware

import (
	"context"
	"net/http"

	"github.com/opentracing/opentracing-go"
)

const (
	tracerConnContextKey = "tracer_connection_key"
)

func RegisterTracerHTTPContext(tracer opentracing.Tracer) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			var ctx = r.Context()
			if tracer != nil {
				ctx = context.WithValue(ctx, tracerConnContextKey, tracer)
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

func GetTracerAccess(ctx context.Context) opentracing.Tracer {
	return ctx.Value(tracerConnContextKey).(opentracing.Tracer)
}
