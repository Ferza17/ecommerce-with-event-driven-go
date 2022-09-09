package repository

import (
	"context"

	"github.com/RoseRocket/xerrs"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Ferza17/event-driven-account-service/utils"
)

const (
	databaseUser    = "user"
	collectionUsers = "users"
)

type cartNOSQLRepository struct {
	db *mongo.Client
}

func NewCartNOSQLRepository(db *mongo.Client) UserNOSQLRepositoryStore {
	return &cartNOSQLRepository{
		db: db,
	}
}

func (q *cartNOSQLRepository) CreateNoSQLSession() (session mongo.Session, err error) {
	session, err = q.db.StartSession()
	if err != nil {
		err = xerrs.Mask(err, utils.ErrQueryTxBegin)
	}
	return
}

func (q *cartNOSQLRepository) StartTransaction(session mongo.Session) (err error) {
	if err = session.StartTransaction(); err != nil {
		err = xerrs.Mask(err, utils.ErrQueryTxBegin)
	}
	return
}

func (q *cartNOSQLRepository) AbortTransaction(session mongo.Session, ctx context.Context) (err error) {
	if err = session.AbortTransaction(ctx); err != nil {
		err = xerrs.Mask(err, utils.ErrQueryTxRollback)
	}
	return
}

func (q *cartNOSQLRepository) CommitTransaction(session mongo.Session, ctx context.Context) (err error) {
	if err = session.CommitTransaction(ctx); err != nil {
		err = xerrs.Mask(err, utils.ErrQueryTxCommit)
	}
	return
}
