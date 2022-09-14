package usecase

import (
	"context"

	"github.com/Ferza17/event-driven-api-gateway/helper/tracing"
	"github.com/Ferza17/event-driven-api-gateway/model/pb"
)

type userUseCase struct {
	userService pb.UserServiceClient
}

func NewUserUseCase(userService pb.UserServiceClient) UserUseCaseStore {
	return &userUseCase{
		userService: userService,
	}
}

func (u userUseCase) Login(ctx context.Context, request *pb.LoginRequest) (response *pb.LoginResponse, err error) {
	response = &pb.LoginResponse{}
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCase-FindUserById")
	defer span.Finish()
	response, err = u.userService.Login(ctx, request)
	return
}

func (u userUseCase) FindUserById(ctx context.Context, request *pb.FindUserByIdRequest) (response *pb.User, err error) {
	response = &pb.User{}
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCase-FindUserById")
	defer span.Finish()
	response, err = u.userService.FindUserById(ctx, request)
	return
}
