package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type CartMongoDBTransactionCreator interface {
	CreateNoSQLSession() (session mongo.Session, err error)
	StartTransaction(session mongo.Session) (err error)
	AbortTransaction(session mongo.Session, ctx context.Context) (err error)
	CommitTransaction(session mongo.Session, ctx context.Context) (err error)
}

type CartMongoDBRepositoryCommand interface {
}

type CartMongoDBRepositoryQuery interface {
}

type CartMongoDBRepositoryStore interface {
	CartMongoDBTransactionCreator
	CartMongoDBRepositoryCommand
	CartMongoDBRepositoryQuery
}
