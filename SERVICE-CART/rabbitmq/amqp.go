package rabbitmq

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/gocql/gocql"
	"github.com/opentracing/opentracing-go"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"

	"github.com/Ferza17/event-driven-cart-service/middleware"
	"github.com/Ferza17/event-driven-cart-service/module/cart"
)

type (
	Server struct {
		amqpConn         *amqp.Connection
		redisClient      *redis.Client
		logger           *zap.Logger
		db               *mongo.Client
		tracer           opentracing.Tracer
		cassandraSession *gocql.Session
	}
	Option func(s *Server)
)

func NewServer(option ...Option) *Server {
	s := &Server{}
	for _, o := range option {
		o(s)
	}
	return s
}

func (srv *Server) Serve() {
	ctx := srv.setup()
	Subscriber(ctx, srv.amqpConn)
}

func (srv *Server) setup() context.Context {
	var ctx = context.Background()
	ctx = middleware.RegisterMongoDBContext(srv.db, ctx)
	ctx = middleware.RegisterRedisContext(srv.redisClient, ctx)
	ctx = middleware.RegisterTracerContext(srv.tracer, ctx)
	ctx = middleware.RegisterCassandraDBContext(srv.cassandraSession, ctx)
	ctx = middleware.RegisterRabbitMQAmqpContext(srv.amqpConn, ctx)

	ctx = cart.RegisterCartUseCaseContext(ctx)
	return ctx
}
