package cart

import (
	"context"
	"net/http"

	"github.com/Ferza17/event-driven-api-gateway/middleware"
	cartUseCase "github.com/Ferza17/event-driven-api-gateway/module/cart/usecase"
	"github.com/Ferza17/event-driven-api-gateway/utils"
)

func newCartUseCase(ctx context.Context) cartUseCase.CartUseCaseStore {
	return cartUseCase.NewCartUseCase(
		middleware.GetCartServiceGrpcClientFromContext(ctx),
		middleware.GetProductServiceGrpcClientFromContext(ctx),
	)
}

func RegisterCartUseCaseHTTPContext() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			var ctx = r.Context()
			next.ServeHTTP(w, r.WithContext(context.WithValue(ctx, utils.CartUseCaseContextKey, newCartUseCase(ctx))))
		}
		return http.HandlerFunc(fn)
	}
}

func RegisterCartUseCaseContext(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, utils.CartUseCaseContextKey, newCartUseCase(ctx))
	return ctx
}

func GetCartUseCaseFromContext(ctx context.Context) cartUseCase.CartUseCaseStore {
	return ctx.Value(utils.CartUseCaseContextKey).(cartUseCase.CartUseCaseStore)
}
