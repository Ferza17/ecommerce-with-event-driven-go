package rest

import (
	"github.com/hashicorp/consul/api"
	"github.com/opentracing/opentracing-go"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

func NewLogger(logger *zap.Logger) Option {
	return func(s *Server) {
		s.logger = logger
	}
}

func NewConsulClient(client *api.Client) Option {
	return func(s *Server) {
		s.consulClient = client
	}
}

func NewTracer(tracer opentracing.Tracer) Option {
	return func(s *Server) {
		s.tracer = tracer
	}
}

func NewAmqpConn(connection *amqp.Connection) Option {
	return func(s *Server) {
		s.amqpConn = connection
	}
}
