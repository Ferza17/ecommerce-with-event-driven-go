package middleware

import (
	"context"
	"net/http"

	"github.com/Ferza17/event-driven-api-gateway/model/pb"
	"github.com/Ferza17/event-driven-api-gateway/utils"
)

func RegisterUserServiceGrpcClientHttpContext(c pb.UserServiceClient) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			var ctx = r.Context()
			if c != nil {
				ctx = context.WithValue(ctx, utils.UserServiceGrpcClientContextKey, c)
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

func RegisterCartServiceGrpcClientHttpContext(c pb.CartServiceClient) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			var ctx = r.Context()
			if c != nil {
				ctx = context.WithValue(ctx, utils.CartServiceGrpcClientContextKey, c)
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

func RegisterProductServiceGrpcClientHttpContext(c pb.ProductServiceClient) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			var ctx = r.Context()
			if c != nil {
				ctx = context.WithValue(ctx, utils.ProductServiceGrpcClientContextKey, c)
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

func GetUserServiceGrpcClientFromContext(ctx context.Context) pb.UserServiceClient {
	return ctx.Value(utils.UserServiceGrpcClientContextKey).(pb.UserServiceClient)
}

func GetCartServiceGrpcClientFromContext(ctx context.Context) pb.CartServiceClient {
	return ctx.Value(utils.CartServiceGrpcClientContextKey).(pb.CartServiceClient)
}

func GetProductServiceGrpcClientFromContext(ctx context.Context) pb.ProductServiceClient {
	return ctx.Value(utils.ProductServiceGrpcClientContextKey).(pb.ProductServiceClient)
}
