package middleware

import (
	"context"
	"net/http"

	"github.com/Ferza17/event-driven-api-gateway/utils"
)

// HostInfo contains information specific to this module
type HostInfo struct {
	CodeName string
}

func Host(codename string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			var ctx = r.Context()
			ctx = context.WithValue(ctx, utils.HostInfoKey, HostInfo{codename})
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}
