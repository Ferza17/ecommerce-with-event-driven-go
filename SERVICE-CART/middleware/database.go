package middleware

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/gocql/gocql"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"

	"github.com/Ferza17/event-driven-cart-service/utils"
)

func UnaryRegisterRedisContext(client *redis.Client) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(context.WithValue(ctx, utils.RedisDBContextKey, client), req)
	}
}

func UnaryRegisterMongoDBContext(conn *mongo.Client) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(context.WithValue(ctx, utils.MongodbContextKey, conn), req)
	}
}

func UnaryRegisterCassandraDBContext(session *gocql.Session) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(context.WithValue(ctx, utils.CassandraDBContextKey, session), req)
	}
}

func RegisterRedisContext(client *redis.Client, ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, utils.RedisDBContextKey, client)
	return ctx
}

func RegisterMongoDBContext(conn *mongo.Client, ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, utils.MongodbContextKey, conn)
	return ctx
}

func RegisterCassandraDBContext(conn *gocql.Session, ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, utils.CassandraDBContextKey, conn)
	return ctx
}

func GetMongoDBFromContext(ctx context.Context) *mongo.Client {
	return ctx.Value(utils.MongodbContextKey).(*mongo.Client)
}

func GetRedisFromContext(ctx context.Context) *redis.Client {
	return ctx.Value(utils.RedisDBContextKey).(*redis.Client)
}

func GetCassandraDBFromContext(ctx context.Context) *gocql.Session {
	return ctx.Value(utils.CassandraDBContextKey).(*gocql.Session)
}
