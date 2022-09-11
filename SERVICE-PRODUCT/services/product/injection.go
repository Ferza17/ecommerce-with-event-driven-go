package user

import (
	"context"

	"google.golang.org/grpc"

	"github.com/Ferza17/event-driven-product-service/services/product/repository"
	productUseCase "github.com/Ferza17/event-driven-product-service/services/product/usecase"

	"github.com/Ferza17/event-driven-product-service/middleware"
	"github.com/Ferza17/event-driven-product-service/utils"
)

func newProductUseCase(ctx context.Context) productUseCase.ProductUseCaseStore {
	return productUseCase.NewProductUseCase(
		repository.NewProductElasticsearchRepository(middleware.GetElasticsearchFromContext(ctx)),
		repository.NewProductCassandraDBRepository(middleware.GetCassandraDBFromContext(ctx)),
		repository.NewProductCacheRepository(middleware.GetRedisFromContext(ctx)),
	)
}

func RegisterProductUseCaseContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, utils.UserUseCaseContextKey, newProductUseCase(ctx))
}

func UnaryRegisterProductUseCaseContext() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(context.WithValue(ctx, utils.UserUseCaseContextKey, newProductUseCase(ctx)), req)
	}
}

func GetProductUseCaseFromContext(c context.Context) productUseCase.ProductUseCaseStore {
	return c.Value(utils.UserUseCaseContextKey).(productUseCase.ProductUseCaseStore)
}
