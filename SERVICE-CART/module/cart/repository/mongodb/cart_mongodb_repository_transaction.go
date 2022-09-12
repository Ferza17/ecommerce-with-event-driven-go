package mongodb

import (
	"context"

	"github.com/RoseRocket/xerrs"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Ferza17/event-driven-cart-service/helper/tracing"
	"github.com/Ferza17/event-driven-cart-service/utils"
)

func (q *cartMongoDBRepository) CreateNoSQLSession(ctx context.Context) (session mongo.Session, err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "CartMongoDBRepository-CreateNoSQLSession")
	defer span.Finish()
	session, err = q.db.StartSession()
	if err != nil {
		err = xerrs.Mask(err, utils.ErrQueryTxBegin)
	}
	return
}

func (q *cartMongoDBRepository) StartTransaction(ctx context.Context, session mongo.Session) (err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "CartMongoDBRepository-StartTransaction")
	defer span.Finish()
	if err = session.StartTransaction(); err != nil {
		err = xerrs.Mask(err, utils.ErrQueryTxBegin)
	}
	return
}

func (q *cartMongoDBRepository) AbortTransaction(ctx context.Context, session mongo.Session) (err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "CartMongoDBRepository-AbortTransaction")
	defer span.Finish()
	if err = session.AbortTransaction(ctx); err != nil {
		err = xerrs.Mask(err, utils.ErrQueryTxRollback)
	}
	return
}

func (q *cartMongoDBRepository) CommitTransaction(ctx context.Context, session mongo.Session) (err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "CartMongoDBRepository-CommitTransaction")
	defer span.Finish()
	if err = session.CommitTransaction(ctx); err != nil {
		err = xerrs.Mask(err, utils.ErrQueryTxCommit)
	}
	return
}
