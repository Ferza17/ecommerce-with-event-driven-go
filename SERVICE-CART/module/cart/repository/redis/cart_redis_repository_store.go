package redis

import (
	"context"

	"github.com/Ferza17/event-driven-cart-service/model/pb"
)

type CartRedisRepositoryCommand interface {
	SetCartLastStateByTransactionId(ctx context.Context, transactionId string, payload *pb.Cart) (err error)
	DeleteCartLastStateByTransactionId(ctx context.Context, transactionId string) (err error)
}

type CartRedisRepositoryQuery interface {
	FindCartLastStateByTransactionId(ctx context.Context, transactionId string) (response *pb.Cart, err error)
}

type CartRedisRepositoryStore interface {
	CartRedisRepositoryCommand
	CartRedisRepositoryQuery
}
