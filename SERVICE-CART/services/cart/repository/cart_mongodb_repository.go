package repository

import (
	"context"

	"github.com/RoseRocket/xerrs"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Ferza17/event-driven-cart-service/utils"
)

const (
	databaseCart    = "cart"
	collectionCarts = "carts"
)

type cartMongoDBRepository struct {
	db *mongo.Client
}

func NewCartMongoDBRepository(db *mongo.Client) CartMongoDBRepositoryStore {
	return &cartMongoDBRepository{
		db: db,
	}
}

func (q *cartMongoDBRepository) CreateNoSQLSession() (session mongo.Session, err error) {
	session, err = q.db.StartSession()
	if err != nil {
		err = xerrs.Mask(err, utils.ErrQueryTxBegin)
	}
	return
}

func (q *cartMongoDBRepository) StartTransaction(session mongo.Session) (err error) {
	if err = session.StartTransaction(); err != nil {
		err = xerrs.Mask(err, utils.ErrQueryTxBegin)
	}
	return
}

func (q *cartMongoDBRepository) AbortTransaction(session mongo.Session, ctx context.Context) (err error) {
	if err = session.AbortTransaction(ctx); err != nil {
		err = xerrs.Mask(err, utils.ErrQueryTxRollback)
	}
	return
}

func (q *cartMongoDBRepository) CommitTransaction(session mongo.Session, ctx context.Context) (err error) {
	if err = session.CommitTransaction(ctx); err != nil {
		err = xerrs.Mask(err, utils.ErrQueryTxCommit)
	}
	return
}
