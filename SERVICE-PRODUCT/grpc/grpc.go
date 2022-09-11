package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-redis/redis/v8"
	"github.com/gocql/gocql"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	otgrpc "github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/Ferza17/event-driven-product-service/middleware"
	product "github.com/Ferza17/event-driven-product-service/services/product"
)

type (
	Server struct {
		address             string
		port                string
		listener            *net.Listener
		redisClient         *redis.Client
		elasticsearchClient *elasticsearch.Client
		grpcServer          *grpc.Server
		logger              *zap.Logger
		cassandraSession    *gocql.Session
		rabbitMQConnection  *amqp.Connection
		tracer              opentracing.Tracer
	}
	Option func(s *Server)
)

func NewServer(address, port string, option ...Option) *Server {
	s := &Server{
		address: address,
		port:    port,
	}
	for _, o := range option {
		o(s)
	}
	s.setup()
	return s
}

func (srv *Server) Serve() {
	// Enable Reflection to Evans grpc client
	reflection.Register(srv.grpcServer)
	if err := srv.grpcServer.Serve(*srv.listener); err != nil {
		log.Fatalln(err)
	}
}

func (srv *Server) setup() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%s", srv.address, srv.port))
	if err != nil {
		log.Fatalln(err)
	}
	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				otgrpc.OpenTracingServerInterceptor(srv.tracer),
				middleware.UnaryRegisterTracerContext(srv.tracer),
				middleware.UnaryRegisterRedisContext(srv.redisClient),
				middleware.UnaryRegisterRabbitMQAmqpContext(srv.rabbitMQConnection),
				middleware.UnaryRegisterElasticsearchContext(srv.elasticsearchClient),
				middleware.UnaryRegisterCassandraDBContext(srv.cassandraSession),
				product.UnaryRegisterProductUseCaseContext(),
			),
		),
	}
	srv.grpcServer = grpc.NewServer(opts...)
	srv.listener = &listen
	srv.RegisterService()
}
