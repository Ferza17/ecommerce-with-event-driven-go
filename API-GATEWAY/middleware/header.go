package middleware

import (
	"context"
	"net/http"
	"regexp"

	"github.com/Ferza17/event-driven-api-gateway/utils"
)

type RequestHeader struct {
	ClientId      string
	ClientVersion string
	DataVersion   string
	DeviceId      string
	ForwardedHost string
	ForwardedIP   string
}

var regexClientIdPattern = regexp.MustCompile("\\w+\\.\\w+")

func Header() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			var ctx = context.WithValue(r.Context(), utils.HeadersContextKey, &RequestHeader{
				ClientId:      r.Header.Get("X-Client-Id"),
				ClientVersion: r.Header.Get("X-Client-Version"),
				DataVersion:   r.Header.Get("X-Data-Version"),
				DeviceId:      r.Header.Get("X-Device-Id"),
				ForwardedHost: r.Header.Get("X-Forwarded-Host"),
				ForwardedIP:   r.Header.Get("X-Forwarded-For"),
			})
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

func RegisterHeaderHTTPContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.WithValue(r.Context(), utils.HeadersContextKey, &RequestHeader{
			ClientId:      r.Header.Get("X-Client-Id"),
			ClientVersion: r.Header.Get("X-Client-Version"),
			DataVersion:   r.Header.Get("X-Data-Version"),
			DeviceId:      r.Header.Get("X-Device-Id"),
			ForwardedHost: r.Header.Get("X-Forwarded-Host"),
			ForwardedIP:   r.Header.Get("X-Forwarded-For"),
		})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetRequestHeader(r *http.Request) *RequestHeader {
	return r.Context().Value(utils.HeadersContextKey).(*RequestHeader)
}
