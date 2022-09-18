package product

import (
	"context"
	"net/http"

	"github.com/Ferza17/event-driven-api-gateway/middleware"
	productUseCase "github.com/Ferza17/event-driven-api-gateway/module/product/usecase"
	"github.com/Ferza17/event-driven-api-gateway/utils"
)

func newProductUseCase(ctx context.Context) productUseCase.ProductUseCaseStore {
	return productUseCase.NewProductUseCase(
		middleware.GetProductServiceGrpcClientFromContext(ctx),
	)
}

func RegisterProductUseCaseHTTPContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		next.ServeHTTP(w, r.WithContext(context.WithValue(ctx, utils.ProductUseCaseContextKey, newProductUseCase(ctx))))
	})
}

func RegisterProductUseCaseContext(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, utils.ProductUseCaseContextKey, newProductUseCase(ctx))
	return ctx
}

func GetProductUseCaseFromContext(ctx context.Context) productUseCase.ProductUseCaseStore {
	return ctx.Value(utils.ProductUseCaseContextKey).(productUseCase.ProductUseCaseStore)
}
