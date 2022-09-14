package middleware

import (
	"context"
	"net/http"

	"github.com/opentracing/opentracing-go"

	"github.com/Ferza17/event-driven-api-gateway/utils"
)

func RegisterTracerHTTPContext(tracer opentracing.Tracer) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			var ctx = r.Context()
			if tracer != nil {
				ctx = context.WithValue(ctx, utils.TracerContextKey, tracer)
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

func GetTracerFromContext(ctx context.Context) opentracing.Tracer {
	return ctx.Value(utils.TracerContextKey).(opentracing.Tracer)
}
