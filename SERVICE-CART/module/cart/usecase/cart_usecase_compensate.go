package usecase

import (
	"context"

	"github.com/Ferza17/event-driven-cart-service/helper/tracing"
)

func (u *cartUseCase) RollbackNewUserSAGA(ctx context.Context, transactionId string) (err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "CartUseCase-RollbackNewUserSAGA")
	defer span.Finish()
	session, err := u.cartMongoDBRepository.CreateNoSQLSession(ctx)
	if err != nil {
		return
	}
	defer session.EndSession(ctx)

	if err = u.cartMongoDBRepository.StartTransaction(ctx, session); err != nil {
		return
	}

	cartLastState, err := u.cartRedisRepository.FindCartLastStateByTransactionId(ctx, transactionId)
	if err != nil {
		return
	}

	if err = u.cartMongoDBRepository.DeleteCartById(ctx, session, cartLastState.GetId()); err != nil {
		return
	}

	if err = u.cartRedisRepository.DeleteCartLastStateByTransactionId(ctx, transactionId); err != nil {
		return
	}

	err = u.cartMongoDBRepository.CommitTransaction(ctx, session)
	return
}
