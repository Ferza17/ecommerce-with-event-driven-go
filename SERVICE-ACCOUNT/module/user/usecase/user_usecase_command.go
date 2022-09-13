package usecase

import (
	"context"
	"encoding/json"

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

	if err = u.userPublisher.PublishOrdinaryMessage(ctx, utils.CreateCartEvent, string(cartRequest)); err != nil {
		err = u.userMongoDBRepository.AbortTransaction(ctx, session)
		return
	}

	err = u.userMongoDBRepository.CommitTransaction(ctx, session)
	return
}
