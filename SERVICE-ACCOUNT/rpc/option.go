package rpc

import (
	"github.com/go-redis/redis/v8"
	"github.com/gocql/gocql"
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

func NewMongoClient(db *mongo.Client) Option {
	return func(s *Server) {
		s.mongodbClient = db
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

func NewRabbitMQConnection(conn *amqp.Connection) Option {
	return func(s *Server) {
		s.amqpConn = conn
	}
}

func NewCassandraSession(session *gocql.Session) Option {
	return func(s *Server) {
		s.cassandraSession = session
	}
}
