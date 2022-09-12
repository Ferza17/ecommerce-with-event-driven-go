package cart

import (
	"context"

	"google.golang.org/grpc"

	cartPublisher "github.com/Ferza17/event-driven-cart-service/module/cart/publisher"
	"github.com/Ferza17/event-driven-cart-service/module/cart/repository/cassandradb"
	"github.com/Ferza17/event-driven-cart-service/module/cart/repository/mongodb"
	"github.com/Ferza17/event-driven-cart-service/module/cart/repository/redis"
	"github.com/Ferza17/event-driven-cart-service/saga"

	"github.com/Ferza17/event-driven-cart-service/middleware"
	cartUseCase "github.com/Ferza17/event-driven-cart-service/module/cart/usecase"
	"github.com/Ferza17/event-driven-cart-service/utils"
)

func newCartUseCase(ctx context.Context) cartUseCase.CartUseCaseStore {
	return cartUseCase.NewCartUseCase(
		mongodb.NewCartMongoDBRepository(middleware.GetMongoDBFromContext(ctx)),
		cassandradb.NewCartCassandraDBRepository(middleware.GetCassandraDBFromContext(ctx)),
		redis.NewCartRedisRepository(middleware.GetRedisFromContext(ctx)),
		cartPublisher.NewCartPublisher(middleware.GetRabbitMQAmqpFromContext(ctx)),
		saga.NewSaga(),
	)
}

func RegisterCartUseCaseContext(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, utils.CartUseCaseContextKey, newCartUseCase(ctx))
	return ctx
}

func UnaryRegisterCartUseCaseContext() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(context.WithValue(ctx, utils.CartUseCaseContextKey, newCartUseCase(ctx)), req)
	}
}

func GetCartUseCaseFromContext(c context.Context) cartUseCase.CartUseCaseStore {
	return c.Value(utils.CartUseCaseContextKey).(cartUseCase.CartUseCaseStore)
}
