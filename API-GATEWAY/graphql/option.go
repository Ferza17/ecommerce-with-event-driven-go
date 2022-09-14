package graphql

import (
	"github.com/hashicorp/consul/api"
	"github.com/opentracing/opentracing-go"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"

	"github.com/Ferza17/event-driven-api-gateway/model/pb"
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

func NewRabbitMQConnection(connection *amqp.Connection) Option {
	return func(s *Server) {
		s.rabbitMQConnection = connection
	}
}

func NewUserServiceGrpcClientConnection(c pb.UserServiceClient) Option {
	return func(s *Server) {
		s.userServiceGrpcClient = c
	}
}

func NewProductServiceGrpcClientConnection(c pb.ProductServiceClient) Option {
	return func(s *Server) {
		s.productServiceGrpcClient = c
	}
}

func NewCartServiceGrpcClientConnection(c pb.CartServiceClient) Option {
	return func(s *Server) {
		s.cartServiceGrpcClient = c
	}
}
