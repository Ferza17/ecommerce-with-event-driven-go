package user

import (
	"context"
	"net/http"

	"github.com/Ferza17/event-driven-api-gateway/middleware"
	userPub "github.com/Ferza17/event-driven-api-gateway/module/user/publisher"
	userUseCase "github.com/Ferza17/event-driven-api-gateway/module/user/usecase"
	"github.com/Ferza17/event-driven-api-gateway/utils"
)

func newUserUseCase(ctx context.Context) userUseCase.UserUseCaseStore {
	return userUseCase.NewUserUseCase(
		middleware.GetUserServiceGrpcClientFromContext(ctx),
		userPub.NewUserPublisher(middleware.GetRabbitMQAmqpFromContext(ctx)),
	)
}

func RegisterUserUseCaseHTTPContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ctx = r.Context()
		next.ServeHTTP(w, r.WithContext(context.WithValue(ctx, utils.UserUseCaseContextKey, newUserUseCase(ctx))))
	})
}

func RegisterUserUseCaseContext(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, utils.UserUseCaseContextKey, newUserUseCase(ctx))
	return ctx
}

func GetUserUseCaseFromContext(ctx context.Context) userUseCase.UserUseCaseStore {
	return ctx.Value(utils.UserUseCaseContextKey).(userUseCase.UserUseCaseStore)
}
