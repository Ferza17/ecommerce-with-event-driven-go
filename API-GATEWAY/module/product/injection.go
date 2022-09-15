package product

import (
	"context"

	"github.com/Ferza17/event-driven-api-gateway/middleware"
	productUseCase "github.com/Ferza17/event-driven-api-gateway/module/product/usecase"
	"github.com/Ferza17/event-driven-api-gateway/utils"
)

func newProductUseCase(ctx context.Context) productUseCase.ProductUseCaseStore {
	return productUseCase.NewProductUseCase(
		middleware.GetProductServiceGrpcClientFromContext(ctx),
	)
}

func RegisterProductUseCaseContext(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, utils.ProductUseCaseContextKey, newProductUseCase(ctx))
	return ctx
}

func GetProductUseCaseFromContext(ctx context.Context) productUseCase.ProductUseCaseStore {
	return ctx.Value(utils.ProductUseCaseContextKey).(productUseCase.ProductUseCaseStore)
}
