package cart

import (
	"context"

	"github.com/Ferza17/event-driven-api-gateway/middleware"
	cartUseCase "github.com/Ferza17/event-driven-api-gateway/module/cart/usecase"
	"github.com/Ferza17/event-driven-api-gateway/utils"
)

func newCartUseCase(ctx context.Context) cartUseCase.CartUseCaseStore {
	return cartUseCase.NewCartUseCase(
		middleware.GetCartServiceGrpcClientFromContext(ctx),
	)
}

func RegisterCartUseCaseContext(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, utils.CartUseCaseContextKey, newCartUseCase(ctx))
	return ctx
}

func GetCartUseCaseFromContext(ctx context.Context) cartUseCase.CartUseCaseStore {
	return ctx.Value(utils.CartUseCaseContextKey).(cartUseCase.CartUseCaseStore)
}
