package middleware

import (
	"context"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-redis/redis/v8"
	"github.com/gocql/gocql"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"

	"github.com/Ferza17/event-driven-product-service/utils"
)

func UnaryRegisterRedisContext(client *redis.Client) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(context.WithValue(ctx, utils.RedisDBContextKey, client), req)
	}
}

func UnaryRegisterCassandraDBContext(session *gocql.Session) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(context.WithValue(ctx, utils.CassandraDBContextKey, session), req)
	}
}

func UnaryRegisterPostgresSQLContext(client *sqlx.DB) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(context.WithValue(ctx, utils.PostgresSQLContextKey, client), req)
	}
}

func UnaryRegisterElasticsearchContext(client *elasticsearch.Client) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(context.WithValue(ctx, utils.ElasticsearchContextKey, client), req)
	}
}

func RegisterRedisContext(client *redis.Client, ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, utils.RedisDBContextKey, client)
	return ctx
}

func RegisterCassandraDBContext(conn *gocql.Session, ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, utils.CassandraDBContextKey, conn)
	return ctx
}

func RegisterPostgresSQLContext(client *sqlx.DB, ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, utils.PostgresSQLContextKey, client)
	return ctx
}

func RegisterElasticsearchContext(client *elasticsearch.Client, ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, utils.ElasticsearchContextKey, client)
	return ctx
}

func GetRedisFromContext(ctx context.Context) *redis.Client {
	return ctx.Value(utils.RedisDBContextKey).(*redis.Client)
}

func GetCassandraDBFromContext(ctx context.Context) *gocql.Session {
	return ctx.Value(utils.CassandraDBContextKey).(*gocql.Session)
}

func GetPostgresSQLFromContext(ctx context.Context) *sqlx.DB {
	return ctx.Value(utils.PostgresSQLContextKey).(*sqlx.DB)
}

func GetElasticsearchFromContext(ctx context.Context) *elasticsearch.Client {
	return ctx.Value(utils.ElasticsearchContextKey).(*elasticsearch.Client)
}
