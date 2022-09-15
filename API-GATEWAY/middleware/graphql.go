package middleware

import (
	"context"
	"net/http"

	"github.com/graphql-go/graphql"

	"github.com/Ferza17/event-driven-api-gateway/utils"
)

func RegisterSchemaConfigHTTPContext(conn graphql.SchemaConfig, contextKey utils.ContextKey) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			var ctx = r.Context()
			ctx = context.WithValue(ctx, contextKey, conn)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

func GetSchemaConfigFromContext(ctx context.Context, contextKey utils.ContextKey) graphql.SchemaConfig {
	return ctx.Value(contextKey).(graphql.SchemaConfig)
}
