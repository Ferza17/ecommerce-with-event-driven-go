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
	userMongoDBRepository     repository.UserMongoDBRepositoryStore
	userCassandraDBRepository repository.UserCassandraDBRepositoryStore
	userRedisRepository       repository.UserRedisRepositoryStore
}

func NewUserUseCase(
	userNOSQLRepository repository.UserMongoDBRepositoryStore,
	userCassandraDBRepository repository.UserCassandraDBRepositoryStore,
	userCacheRepository repository.UserRedisRepositoryStore,
) UserUseCaseStore {
	return &userUseCase{
		userMongoDBRepository:     userNOSQLRepository,
		userCassandraDBRepository: userCassandraDBRepository,
		userRedisRepository:       userCacheRepository,
	}
}

func (u *userUseCase) CreateUser(ctx context.Context, request *pb.RegisterRequest) (response *pb.RegisterResponse, err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCase-CreateUser")
	defer span.Finish()

	request.Password, err = hash.Hashed(request.GetPassword())
	if err != nil {
		return
	}

	session, err := u.userMongoDBRepository.CreateNoSQLSession()
	if err != nil {
		return
	}
	defer session.EndSession(ctx)

	if err = u.userMongoDBRepository.StartTransaction(session); err != nil {
		return
	}

	if response, err = u.userMongoDBRepository.CreateUser(ctx, session, request); err != nil {
		err = u.userMongoDBRepository.AbortTransaction(session, ctx)
		return
	}

	err = u.userMongoDBRepository.CommitTransaction(session, ctx)
	return
}

func (u *userUseCase) FindUserByEmailAndPassword(ctx context.Context, request *pb.LoginRequest) (response *pb.LoginResponse, err error) {
	response = &pb.LoginResponse{}
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCase-FindUserByEmail")
	defer span.Finish()
	user, err := u.userMongoDBRepository.FindUserByEmail(ctx, request.GetEmail())
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
