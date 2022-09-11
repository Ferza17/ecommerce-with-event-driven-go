package cart

import (
	"context"

	"google.golang.org/grpc"

	"github.com/Ferza17/event-driven-cart-service/services/cart/repository"

	"github.com/Ferza17/event-driven-cart-service/middleware"
	cartUseCase "github.com/Ferza17/event-driven-cart-service/services/cart/usecase"
	"github.com/Ferza17/event-driven-cart-service/utils"
)

func newCartUseCase(ctx context.Context) cartUseCase.CartUseCaseStore {
	return cartUseCase.NewCartUseCase(
		repository.NewCartMongoDBRepository(middleware.GetMongoDBFromContext(ctx)),
		repository.NewCartCassandraDBRepository(middleware.GetCassandraDBFromContext(ctx)),
		repository.NewCartRedisRepository(middleware.GetRedisFromContext(ctx)),
	)
}

func RegisterCartUseCaseContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, utils.CartUseCaseContextKey, newCartUseCase(ctx))
}

func UnaryRegisterCartUseCaseContext() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(context.WithValue(ctx, utils.CartUseCaseContextKey, newCartUseCase(ctx)), req)
	}
}

func GetUserUseCaseFromContext(c context.Context) cartUseCase.CartUseCaseStore {
	return c.Value(utils.CartUseCaseContextKey).(cartUseCase.CartUseCaseStore)
}
