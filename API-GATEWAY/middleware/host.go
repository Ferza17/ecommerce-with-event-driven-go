package middleware

import (
	"context"
	"net/http"
)

// HostInfo contains information specific to this services
type HostInfo struct {
	CodeName string
}

const hostInfoKey = "host_info"

func Host(codename string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			var ctx = r.Context()
			ctx = context.WithValue(ctx, hostInfoKey, HostInfo{codename})
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}
