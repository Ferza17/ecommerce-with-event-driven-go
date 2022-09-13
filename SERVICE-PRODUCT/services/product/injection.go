package user

import (
	"context"

	"google.golang.org/grpc"

	"github.com/Ferza17/event-driven-product-service/saga"
	"github.com/Ferza17/event-driven-product-service/services/product/repository/cassandradb"
	"github.com/Ferza17/event-driven-product-service/services/product/repository/elasticsearch"
	"github.com/Ferza17/event-driven-product-service/services/product/repository/postgres"
	"github.com/Ferza17/event-driven-product-service/services/product/repository/redis"
	productUseCase "github.com/Ferza17/event-driven-product-service/services/product/usecase"

	"github.com/Ferza17/event-driven-product-service/middleware"
	"github.com/Ferza17/event-driven-product-service/utils"
)

func newProductUseCase(ctx context.Context) productUseCase.ProductUseCaseStore {
	return productUseCase.NewProductUseCase(
		elasticsearch.NewProductElasticsearchRepository(middleware.GetElasticsearchFromContext(ctx)),
		cassandradb.NewProductCassandraDBRepository(middleware.GetCassandraDBFromContext(ctx)),
		postgres.NewProductPostgresRepository(middleware.GetPostgresSQLFromContext(ctx), middleware.GetPostgresSQLFromContext(ctx)),
		redis.NewProductRedisRepository(middleware.GetRedisFromContext(ctx)),
		saga.NewSaga(),
	)
}

func RegisterProductUseCaseContext(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, utils.UserUseCaseContextKey, newProductUseCase(ctx))
	return ctx
}

func UnaryRegisterProductUseCaseContext() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(context.WithValue(ctx, utils.UserUseCaseContextKey, newProductUseCase(ctx)), req)
	}
}

func GetProductUseCaseFromContext(c context.Context) productUseCase.ProductUseCaseStore {
	return c.Value(utils.UserUseCaseContextKey).(productUseCase.ProductUseCaseStore)
}
