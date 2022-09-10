package rpc

import (
	"github.com/go-redis/redis/v8"
	"github.com/opentracing/opentracing-go"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func NewLogger(logger *zap.Logger) Option {
	return func(s *Server) {
		s.logger = logger
	}
}

func NewDB(db *mongo.Client) Option {
	return func(s *Server) {
		s.db = db
	}
}

func NewRedisClient(client *redis.Client) Option {
	return func(s *Server) {
		s.redisClient = client
	}
}

func NewTracer(tracer opentracing.Tracer) Option {
	return func(s *Server) {
		s.tracer = tracer
	}
}

func NewAMQP(conn *amqp.Connection) Option {
	return func(s *Server) {
		s.amqpConn = conn
	}
}
