package usecase

import (
	"context"

	"github.com/Ferza17/event-driven-account-service/helper/tracing"
)

func (u *userUseCase) RollbackNewUserSAGA(ctx context.Context, transactionId string) (err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCase-RollbackNewUserSAGA")
	defer span.Finish()
	session, err := u.userMongoDBRepository.CreateNoSQLSession(ctx)
	if err != nil {
		return
	}
	defer session.EndSession(ctx)

	if err = u.userMongoDBRepository.StartTransaction(ctx, session); err != nil {
		return
	}

	userLastState, err := u.userRedisRepository.FindUserLastStateByTransactionId(ctx, transactionId)
	if err != nil {
		return
	}

	if err = u.userMongoDBRepository.DeleteUserByUserId(ctx, session, userLastState.GetId()); err != nil {
		return
	}

	if err = u.userRedisRepository.DeleteUserLastStateByTransactionId(ctx, transactionId); err != nil {
		return
	}

	err = u.userMongoDBRepository.CommitTransaction(ctx, session)
	return
}
