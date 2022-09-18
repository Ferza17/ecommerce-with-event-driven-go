package graphql

import (
	"context"

	"github.com/google/uuid"

	"github.com/Ferza17/event-driven-api-gateway/helper/tracing"
	"github.com/Ferza17/event-driven-api-gateway/middleware"
	"github.com/Ferza17/event-driven-api-gateway/model/graph/model"
	"github.com/Ferza17/event-driven-api-gateway/model/pb"
	"github.com/Ferza17/event-driven-api-gateway/module/user"
	"github.com/Ferza17/event-driven-api-gateway/module/user/subscriber"
)

func HandleUserLogin(ctx context.Context, input *model.LoginRequest) (response *model.LoginResponse, err error) {
	var (
		userUseCase = user.GetUserUseCaseFromContext(ctx)
	)
	span, ctx := tracing.StartSpanFromContext(ctx, "UserGRPCPresenter-HandleUserLogin")
	defer span.Finish()
	user, err := userUseCase.FindUserByEmailAndPassword(
		ctx,
		&pb.LoginRequest{
			Email:    input.Email,
			Password: input.Password,
		},
	)
	if err != nil {
		return
	}
	response = &model.LoginResponse{
		UserID: user.GetUserId(),
		Token:  user.GetToken(),
	}
	return
}

func HandleFindUserByIdFindUserByID(ctx context.Context) (response *model.User, err error) {
	var (
		userUseCase = user.GetUserUseCaseFromContext(ctx)
		identity    = middleware.GetTokenIdentityFromContext(ctx)
	)
	span, ctx := tracing.StartSpanFromContext(ctx, "UserGRPCPresenter-HandleFindUserById")
	defer span.Finish()
	user, err := userUseCase.FindUserById(
		ctx,
		&pb.FindUserByIdRequest{
			Id: identity.UserId,
		},
	)
	if err != nil {
		return
	}
	response = &model.User{
		ID:          user.GetId(),
		Username:    user.GetUsername(),
		Email:       user.GetEmail(),
		Password:    user.GetPassword(),
		CreatedAt:   int(user.GetCreatedAt()),
		UpdatedAt:   int(user.GetUpdatedAt()),
		DiscardedAt: int(user.GetDiscardedAt()),
	}
	return
}

func HandleUserRegister(ctx context.Context, input *model.RegisterRequest) (response *model.CommandResponse, err error) {
	var (
		userUseCase = user.GetUserUseCaseFromContext(ctx)
	)
	span, ctx := tracing.StartSpanFromContext(ctx, "UserGRPCPresenter-HandleUserRegister")
	defer span.Finish()
	response, err = userUseCase.CreateUser(
		ctx,
		&pb.RegisterRequest{
			Username:      input.Username,
			Email:         input.Email,
			Password:      input.Password,
			TransactionId: uuid.NewString(),
		},
	)
	return
}

func HandleSubscribeChangeUserState(ctx context.Context, input *model.SubscribeChangeUserState) (<-chan *model.User, error) {
	var (
		userSub = subscriber.NewUserSubscriber(middleware.GetRabbitMQAmqpFromContext(ctx))
		err     error
	)
	userCh, err := userSub.SubscribeNewUserState(ctx, input.ID)
	if err != nil {
		return nil, err
	}
	return userCh, nil
}
