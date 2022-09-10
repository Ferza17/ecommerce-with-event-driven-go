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
	redis := middleware.GetRedisFromContext(ctx)
	db := middleware.GetMongoDBFromContext(ctx)
	newUserNOSQLRepository := repository.NewUserNOSQLRepository(db)
	newUserCacheRepository := repository.NewUserCacheRepository(redis)
	return userUseCase.NewUserUseCase(newUserNOSQLRepository, newUserCacheRepository)
}

func RegisterUserUseCaseContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, utils.UserUseCaseContextKey, newUserUseCase(ctx))
}

func UnaryRegisterUserUseCaseContext() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		ctx = context.WithValue(ctx, utils.UserUseCaseContextKey, newUserUseCase(ctx))
		resp, err = handler(ctx, req)
		return resp, err
	}
}

func GetUserUseCaseFromContext(c context.Context) userUseCase.UserUseCaseStore {
	return c.Value(utils.UserUseCaseContextKey).(userUseCase.UserUseCaseStore)
}
