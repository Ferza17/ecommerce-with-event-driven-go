package usecase

import (
	"context"

	"github.com/Ferza17/event-driven-cart-service/helper/tracing"
	"github.com/Ferza17/event-driven-cart-service/model/pb"
	"github.com/Ferza17/event-driven-cart-service/saga"
	"github.com/Ferza17/event-driven-cart-service/utils"
)

func (u *cartUseCase) CreateCart(ctx context.Context, request *pb.CreateCartRequest) (response *pb.CreateCartResponse, err error) {
	response = &pb.CreateCartResponse{}
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCase-CreateUser")
	defer span.Finish()

	sagaSuccess := &saga.Step{
		TransactionId: request.GetTransactionId(),
		Status:        utils.SagaStatusSuccess,
		Counter:       2,
	}

	sagaFailed := &saga.Step{
		TransactionId: request.GetTransactionId(),
		Status:        utils.SagaStatusFailed,
		Counter:       2,
	}

	session, err := u.cartMongoDBRepository.CreateNoSQLSession(ctx)
	if err != nil {
		if err = u.cartPublisher.PublishSagaMessage(ctx, utils.NewCartEventSaga, sagaFailed); err != nil {
			return
		}
		return
	}
	defer session.EndSession(ctx)

	if err = u.cartMongoDBRepository.StartTransaction(ctx, session); err != nil {
		if err = u.cartPublisher.PublishSagaMessage(ctx, utils.NewCartEventSaga, sagaSuccess); err != nil {
			return
		}
		return
	}

	cartLastState, err := u.cartMongoDBRepository.CreateCart(ctx, session, request)
	if err != nil {
		if err = u.cartPublisher.PublishSagaMessage(ctx, utils.NewCartEventSaga, sagaFailed); err != nil {
			return
		}
		err = u.cartMongoDBRepository.AbortTransaction(ctx, session)
		return
	}

	if err = u.cartRedisRepository.SetCartLastStateByTransactionId(ctx, request.GetTransactionId(), cartLastState); err != nil {
		if err = u.cartPublisher.PublishSagaMessage(ctx, utils.NewCartEventSaga, sagaFailed); err != nil {
			return
		}
		err = u.cartMongoDBRepository.AbortTransaction(ctx, session)
		return
	}

	if err = u.cartPublisher.PublishSagaMessage(ctx, utils.NewCartEventSaga, sagaSuccess); err != nil {
		if err = u.cartPublisher.PublishSagaMessage(ctx, utils.NewCartEventSaga, sagaFailed); err != nil {
			return
		}
		err = u.cartMongoDBRepository.AbortTransaction(ctx, session)
		return
	}

	if err = u.cartMongoDBRepository.CommitTransaction(ctx, session); err != nil {
		err = u.cartPublisher.PublishSagaMessage(ctx, utils.NewCartEventSaga, sagaFailed)
	}
	return
}
