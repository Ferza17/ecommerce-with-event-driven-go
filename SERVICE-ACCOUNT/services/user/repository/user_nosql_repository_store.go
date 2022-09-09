package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserNOSQLTransactionCreator interface {
	CreateNoSQLSession() (session mongo.Session, err error)
	StartTransaction(session mongo.Session) (err error)
	AbortTransaction(session mongo.Session, ctx context.Context) (err error)
	CommitTransaction(session mongo.Session, ctx context.Context) (err error)
}

type UserNOSQLRepositoryWriter interface {
}

type UserNOSQLRepositoryReader interface {
}

type UserNOSQLRepositoryStore interface {
	UserNOSQLTransactionCreator
	UserNOSQLRepositoryWriter
	UserNOSQLRepositoryReader
}
