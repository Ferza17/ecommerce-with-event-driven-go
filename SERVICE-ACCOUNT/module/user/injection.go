package user

import (
	"context"

	"google.golang.org/grpc"

	"github.com/Ferza17/event-driven-account-service/middleware"
	"github.com/Ferza17/event-driven-account-service/module/user/publisher"
	"github.com/Ferza17/event-driven-account-service/module/user/repository/cassandradb"
	"github.com/Ferza17/event-driven-account-service/module/user/repository/mongodb"
	"github.com/Ferza17/event-driven-account-service/module/user/repository/redis"
	userUseCase "github.com/Ferza17/event-driven-account-service/module/user/usecase"
	"github.com/Ferza17/event-driven-account-service/saga"
	"github.com/Ferza17/event-driven-account-service/utils"
)

func newUserUseCase(ctx context.Context) userUseCase.UserUseCaseStore {
	return userUseCase.NewUserUseCase(
		mongodb.NewUserMongoDBRepository(middleware.GetMongoDBFromContext(ctx)),
		cassandradb.NewUserCassandraDBRepository(middleware.GetCassandraDBFromContext(ctx)),
		redis.NewUserRedisRepository(middleware.GetRedisFromContext(ctx)),
		publisher.NewUserPublisher(middleware.GetRabbitMQAmqpFromContext(ctx)),
		saga.NewSaga(),
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
