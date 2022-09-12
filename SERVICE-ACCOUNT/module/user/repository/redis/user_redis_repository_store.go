package redis

import (
	"context"

	"github.com/Ferza17/event-driven-account-service/model/pb"
)

type UserRedisRepositoryCommand interface {
	SetUserLastStateByTransactionId(ctx context.Context, transactionId string, payload *pb.User) (err error)
	DeleteUserLastStateByTransactionId(ctx context.Context, transactionId string) (err error)
}

type UserRedisRepositoryQuery interface {
	FindUserLastStateByTransactionId(ctx context.Context, transactionId string) (response *pb.User, err error)
}

type UserRedisRepositoryStore interface {
	UserRedisRepositoryCommand
	UserRedisRepositoryQuery
}
