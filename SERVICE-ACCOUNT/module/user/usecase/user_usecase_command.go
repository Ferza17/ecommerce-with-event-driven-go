package usecase

import (
	"context"
	"encoding/json"

	"github.com/RoseRocket/xerrs"

	"github.com/Ferza17/event-driven-account-service/helper/hash"
	"github.com/Ferza17/event-driven-account-service/helper/tracing"
	"github.com/Ferza17/event-driven-account-service/model/pb"
	"github.com/Ferza17/event-driven-account-service/utils"
)

func (u *userUseCase) CreateUser(ctx context.Context, request *pb.RegisterRequest) (response *pb.RegisterResponse, err error) {
	response = &pb.RegisterResponse{}
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCase-CreateUser")
	defer span.Finish()

	request.Password, err = hash.Hashed(request.GetPassword())
	if err != nil {
		return
	}

	session, err := u.userMongoDBRepository.CreateNoSQLSession(ctx)
	if err != nil {
		return
	}
	defer session.EndSession(ctx)

	if err = u.userMongoDBRepository.StartTransaction(ctx, session); err != nil {
		return
	}

	userLastState, err := u.userMongoDBRepository.CreateUser(ctx, session, request)
	if err != nil {
		err = u.userMongoDBRepository.AbortTransaction(ctx, session)
		return
	}
	response.UserId = userLastState.GetId()

	if err = u.userRedisRepository.SetUserLastStateByTransactionId(ctx, request.GetTransactionId(), userLastState); err != nil {
		return
	}

	cartRequest, err := json.Marshal(
		pb.CreateCartRequest{
			TransactionId: request.GetTransactionId(),
			UserId:        userLastState.GetId(),
		},
	)
	if err != nil {
		err = u.userMongoDBRepository.AbortTransaction(ctx, session)
		return
	}

	if err = u.userPub.PublishOrdinaryMessage(ctx, utils.CreateCartEvent, string(cartRequest)); err != nil {
		err = u.userMongoDBRepository.AbortTransaction(ctx, session)
		return
	}

	err = u.userMongoDBRepository.CommitTransaction(ctx, session)
	return
}

func (u *userUseCase) UpdateUserByUserId(ctx context.Context, request *pb.UpdateUserByUserIdRequest) (err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCase-UpdateUserByUserId")
	defer span.Finish()
	session, err := u.userMongoDBRepository.CreateNoSQLSession(ctx)
	if err != nil {
		return
	}
	defer session.EndSession(ctx)

	if err = u.userMongoDBRepository.StartTransaction(ctx, session); err != nil {
		return
	}

	if err = u.userMongoDBRepository.UpdateUserByUserId(ctx, session, request); err != nil {
		err = u.userMongoDBRepository.AbortTransaction(ctx, session)
		return
	}

	if err = u.userMongoDBRepository.CommitTransaction(ctx, session); err != nil {
		return
	}

	user, err := u.userMongoDBRepository.FindUserById(ctx, request.GetId())
	if err != nil {
		return
	}

	payload, err := u.userPub.ParsePayloadToString(ctx, user)
	if err != nil {
		return
	}

	if err = u.userPub.PublishOrdinaryMessage(ctx, utils.UserNewState, payload); err != nil {
		err = xerrs.Mask(err, utils.ErrInternalServerError)
		return
	}

	return
}
