package mongodb

import (
	"context"

	"github.com/RoseRocket/xerrs"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Ferza17/event-driven-account-service/helper/tracing"
	"github.com/Ferza17/event-driven-account-service/utils"
)

func (q *userMongoDBRepository) CreateNoSQLSession(ctx context.Context) (session mongo.Session, err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserMongoDBRepository-CreateNoSQLSession")
	defer span.Finish()
	session, err = q.db.StartSession()
	if err != nil {
		err = xerrs.Mask(err, utils.ErrQueryTxBegin)
	}
	return
}

func (q *userMongoDBRepository) StartTransaction(ctx context.Context, session mongo.Session) (err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserMongoDBRepository-StartTransaction")
	defer span.Finish()
	if err = session.StartTransaction(); err != nil {
		err = xerrs.Mask(err, utils.ErrQueryTxBegin)
	}
	return
}

func (q *userMongoDBRepository) AbortTransaction(ctx context.Context, session mongo.Session) (err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserMongoDBRepository-AbortTransaction")
	defer span.Finish()
	if err = session.AbortTransaction(ctx); err != nil {
		err = xerrs.Mask(err, utils.ErrQueryTxRollback)
	}
	return
}

func (q *userMongoDBRepository) CommitTransaction(ctx context.Context, session mongo.Session) (err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserMongoDBRepository-CommitTransaction")
	defer span.Finish()
	if err = session.CommitTransaction(ctx); err != nil {
		err = xerrs.Mask(err, utils.ErrQueryTxCommit)
	}
	return
}
