package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Ferza17/event-driven-cart-service/model/pb"
)

type CartMongoDBTransactionCreator interface {
	CreateNoSQLSession(ctx context.Context) (session mongo.Session, err error)
	StartTransaction(ctx context.Context, session mongo.Session) (err error)
	AbortTransaction(ctx context.Context, session mongo.Session) (err error)
	CommitTransaction(ctx context.Context, session mongo.Session) (err error)
}

type CartMongoDBRepositoryCommand interface {
	CreateCart(ctx context.Context, session mongo.Session, request *pb.CreateCartRequest) (response *pb.Cart, err error)
	DeleteCartById(ctx context.Context, session mongo.Session, id string) (err error)
}

type CartMongoDBRepositoryQuery interface {
	FindCartByUserId(ctx context.Context, id string) (response *pb.Cart, err error)
}

type CartMongoDBRepositoryStore interface {
	CartMongoDBTransactionCreator
	CartMongoDBRepositoryCommand
	CartMongoDBRepositoryQuery
}
