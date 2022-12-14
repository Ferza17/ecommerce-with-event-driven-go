package grpc

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-redis/redis/v8"
	"github.com/gocql/gocql"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

func NewLogger(logger *zap.Logger) Option {
	return func(s *Server) {
		s.logger = logger
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
		s.rabbitMQConnection = conn
	}
}

func NewCassandraSession(session *gocql.Session) Option {
	return func(s *Server) {
		s.cassandraSession = session
	}
}

func NewElasticsearchClient(client *elasticsearch.Client) Option {
	return func(s *Server) {
		s.elasticsearchClient = client
	}
}

func NewPostgresClient(client *sqlx.DB) Option {
	return func(s *Server) {
		s.postgresClient = client
	}
}
