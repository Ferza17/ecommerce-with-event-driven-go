package usecase

import (
	"context"

	"github.com/RoseRocket/xerrs"

	"github.com/Ferza17/event-driven-account-service/helper/hash"
	"github.com/Ferza17/event-driven-account-service/helper/tracing"
	"github.com/Ferza17/event-driven-account-service/model/pb"
	"github.com/Ferza17/event-driven-account-service/utils"
)

func (u *userUseCase) FindUserByEmailAndPassword(ctx context.Context, request *pb.LoginRequest) (response *pb.LoginResponse, err error) {
	response = &pb.LoginResponse{}
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCase-FindUserByEmailAndPassword")
	defer span.Finish()
	user, err := u.userMongoDBRepository.FindUserByEmail(ctx, request.GetEmail())
	if user.GetId() == "" {
		err = xerrs.Mask(utils.ErrNotFound, utils.ErrNotFound)
	}
	if isAuthenticated := hash.Compare(user.GetPassword(), request.Password); !isAuthenticated {
		err = xerrs.Mask(utils.ErrNotFound, utils.ErrNotFound)
		return
	}
	response.UserId = user.GetId()
	return
}

func (u *userUseCase) FindUserById(ctx context.Context, request *pb.FindUserByIdRequest) (response *pb.User, err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCase-FindUserById")
	defer span.Finish()
	response, err = u.userMongoDBRepository.FindUserById(ctx, request.GetId())
	return
}

func (u *userUseCase) FindUserByEmail(ctx context.Context, request *pb.FindUserByEmailRequest) (response *pb.User, err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCase-FindUserByEmail")
	defer span.Finish()
	response, err = u.userMongoDBRepository.FindUserByEmail(ctx, request.GetEmail())
	return
}
