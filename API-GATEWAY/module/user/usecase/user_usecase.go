package usecase

import (
	"context"
	"encoding/json"

	"github.com/RoseRocket/xerrs"

	errorHandler "github.com/Ferza17/event-driven-api-gateway/helper/error"
	"github.com/Ferza17/event-driven-api-gateway/helper/tracing"
	"github.com/Ferza17/event-driven-api-gateway/middleware"
	"github.com/Ferza17/event-driven-api-gateway/model/pb"
	"github.com/Ferza17/event-driven-api-gateway/model/schema"
	userPub "github.com/Ferza17/event-driven-api-gateway/module/user/publisher"
	"github.com/Ferza17/event-driven-api-gateway/utils"
)

type userUseCase struct {
	userService   pb.UserServiceClient
	userPublisher userPub.UserPublisherStore
}

func NewUserUseCase(
	userService pb.UserServiceClient,
	userPublisher userPub.UserPublisherStore,
) UserUseCaseStore {
	return &userUseCase{
		userService:   userService,
		userPublisher: userPublisher,
	}
}

func (u *userUseCase) FindUserByEmailAndPassword(ctx context.Context, request *pb.LoginRequest) (response *pb.LoginResponse, err error) {
	response = &pb.LoginResponse{}
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCase-FindUserByEmailAndPassword")
	defer span.Finish()
	response, err = u.userService.Login(ctx, request)
	if err != nil {
		err = errorHandler.HandlerGrpcError(err)
		return
	}
	token, err := middleware.CreateToken(response.GetUserId())
	if err != nil {
		return
	}
	response.Token = token
	return
}

func (u *userUseCase) FindUserById(ctx context.Context, request *pb.FindUserByIdRequest) (response *pb.User, err error) {
	response = &pb.User{}
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCase-FindUserById")
	defer span.Finish()
	response, err = u.userService.FindUserById(ctx, request)
	if err != nil {
		err = errorHandler.HandlerGrpcError(err)
	}
	return
}

func (u *userUseCase) CreateUser(ctx context.Context, request *pb.RegisterRequest) (response schema.CommandResponse, err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCase-CreateUser")
	defer span.Finish()
	payload, err := json.Marshal(request)
	if err != nil {
		err = xerrs.Mask(err, utils.ErrBadRequest)
		return
	}
	// validate email user if already exist
	previousUser, err := u.userService.FindUserByEmail(ctx, &pb.FindUserByEmailRequest{Email: request.GetEmail()})
	if err != nil {
		err = errorHandler.HandlerGrpcError(err)
		return
	}
	if previousUser != nil && previousUser.GetId() != "" {
		err = xerrs.Mask(utils.ErrItemAlreadyExist, utils.ErrItemAlreadyExist)
		return
	}

	if err = u.userPublisher.PublishOrdinaryMessage(ctx, utils.CreateUserEvent, string(payload)); err != nil {
		err = xerrs.Mask(utils.ErrInternalServerError, utils.ErrInternalServerError)
		return
	}
	response.Message = schema.CommandSuccess
	return
}
