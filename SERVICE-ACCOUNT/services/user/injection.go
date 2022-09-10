package user

import (
	"context"

	"google.golang.org/grpc"

	"github.com/Ferza17/event-driven-account-service/middleware"
	"github.com/Ferza17/event-driven-account-service/services/user/repository"
	userUseCase "github.com/Ferza17/event-driven-account-service/services/user/usecase"
	"github.com/Ferza17/event-driven-account-service/utils"
)

func newUserUseCase(ctx context.Context) userUseCase.UserUseCaseStore {
	return userUseCase.NewUserUseCase(
		repository.NewUserMongoDBRepository(middleware.GetMongoDBFromContext(ctx)),
		repository.NewUserCassandraDBRepository(middleware.GetCassandraDBFromContext(ctx)),
		repository.NewUserCacheRepository(middleware.GetRedisFromContext(ctx)),
	)
}

func RegisterUserUseCaseContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, utils.UserUseCaseContextKey, newUserUseCase(ctx))
}

func UnaryRegisterUserUseCaseContext() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(context.WithValue(ctx, utils.UserUseCaseContextKey, newUserUseCase(ctx)), req)
	}
}

func GetUserUseCaseFromContext(c context.Context) userUseCase.UserUseCaseStore {
	return c.Value(utils.UserUseCaseContextKey).(userUseCase.UserUseCaseStore)
}
