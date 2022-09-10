package middleware

import (
	"context"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"

	"github.com/Ferza17/event-driven-account-service/utils"
)

func UnaryRegisterRedisContext(client *redis.Client) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		newCtx := context.WithValue(ctx, utils.RedisDBContextKey, client)
		resp, err = handler(newCtx, req)
		return resp, err
	}
}

func UnaryRegisterMongoDBContext(conn *mongo.Client) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		newCtx := context.WithValue(ctx, utils.MongodbContextKey, conn)
		resp, err = handler(newCtx, req)
		return resp, err
	}
}

func RegisterRedisContext(client *redis.Client, parentContext context.Context) (ctx context.Context) {
	ctx = context.WithValue(parentContext, utils.RedisDBContextKey, client)
	return
}

func RegisterMongoDBContext(conn *mongo.Client, parentContext context.Context) (ctx context.Context) {
	ctx = context.WithValue(parentContext, utils.MongodbContextKey, conn)
	return
}

func GetMongoDBFromContext(ctx context.Context) *mongo.Client {
	return ctx.Value(utils.MongodbContextKey).(*mongo.Client)
}

func GetRedisFromContext(ctx context.Context) *redis.Client {
	return ctx.Value(utils.RedisDBContextKey).(*redis.Client)
}
