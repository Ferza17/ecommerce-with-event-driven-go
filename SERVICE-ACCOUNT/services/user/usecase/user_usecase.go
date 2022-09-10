package usecase

import (
	"context"

	"github.com/RoseRocket/xerrs"

	"github.com/Ferza17/event-driven-account-service/helper/hash"
	"github.com/Ferza17/event-driven-account-service/helper/tracing"
	"github.com/Ferza17/event-driven-account-service/model/pb"
	"github.com/Ferza17/event-driven-account-service/services/user/repository"
	"github.com/Ferza17/event-driven-account-service/utils"
)

type userUseCase struct {
	userNOSQLRepository repository.UserNOSQLRepositoryStore
	userCacheRepository repository.UserCacheRepositoryStore
}

func NewUserUseCase(
	userNOSQLRepository repository.UserNOSQLRepositoryStore,
	userCacheRepository repository.UserCacheRepositoryStore,
) UserUseCaseStore {
	return &userUseCase{
		userNOSQLRepository: userNOSQLRepository,
		userCacheRepository: userCacheRepository,
	}
}

func (u *userUseCase) CreateUser(ctx context.Context, request *pb.RegisterRequest) (response *pb.RegisterResponse, err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCase-CreateUser")
	defer span.Finish()

	request.Password, err = hash.Hashed(request.GetPassword())
	if err != nil {
		return
	}

	session, err := u.userNOSQLRepository.CreateNoSQLSession()
	if err != nil {
		return
	}
	defer session.EndSession(ctx)

	if err = u.userNOSQLRepository.StartTransaction(session); err != nil {
		return
	}

	if response, err = u.userNOSQLRepository.CreateUser(ctx, session, request); err != nil {
		err = u.userNOSQLRepository.AbortTransaction(session, ctx)
		return
	}

	err = u.userNOSQLRepository.CommitTransaction(session, ctx)
	return
}

func (u *userUseCase) FindUserByEmailAndPassword(ctx context.Context, request *pb.LoginRequest) (response *pb.LoginResponse, err error) {
	response = &pb.LoginResponse{}
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCase-FindUserByEmail")
	defer span.Finish()
	user, err := u.userNOSQLRepository.FindUserByEmail(ctx, request.GetEmail())
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
	response, err = u.userNOSQLRepository.FindUserById(ctx, request.GetId())
	return
}
