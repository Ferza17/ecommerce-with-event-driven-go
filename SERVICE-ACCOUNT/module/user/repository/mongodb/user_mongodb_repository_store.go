package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Ferza17/event-driven-account-service/model/pb"
)

type UserMongoDBTransactionCreator interface {
	CreateNoSQLSession(ctx context.Context) (session mongo.Session, err error)
	StartTransaction(ctx context.Context, session mongo.Session) (err error)
	AbortTransaction(ctx context.Context, session mongo.Session) (err error)
	CommitTransaction(ctx context.Context, session mongo.Session) (err error)
}

type UserMongoDBRepositoryCommand interface {
	CreateUser(ctx context.Context, session mongo.Session, request *pb.RegisterRequest) (response *pb.User, err error)
	DeleteUserByUserId(ctx context.Context, session mongo.Session, id string) (err error)
}

type UserMongoDBRepositoryQuery interface {
	FindUserByEmail(ctx context.Context, email string) (response *pb.User, err error)
	FindUserById(ctx context.Context, id string) (response *pb.User, err error)
}

type UserMongoDBRepositoryStore interface {
	UserMongoDBTransactionCreator
	UserMongoDBRepositoryCommand
	UserMongoDBRepositoryQuery
}
