package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/go-redis/redis/v8"
	"github.com/gocql/gocql"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	otgrpc "github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/Ferza17/event-driven-account-service/middleware"
	userService "github.com/Ferza17/event-driven-account-service/module/user"
)

type (
	Server struct {
		address            string
		port               string
		listener           *net.Listener
		redisClient        *redis.Client
		grpcServer         *grpc.Server
		logger             *zap.Logger
		mongodbClient      *mongo.Client
		cassandraSession   *gocql.Session
		rabbitMQConnection *amqp.Connection
		tracer             opentracing.Tracer
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
				middleware.UnaryRegisterMongoDBContext(srv.mongodbClient),
				middleware.UnaryRegisterRedisContext(srv.redisClient),
				middleware.UnaryRegisterRabbitMQAmqpContext(srv.rabbitMQConnection),
				middleware.UnaryRegisterTracerContext(srv.tracer),
				middleware.UnaryRegisterCassandraDBContext(srv.cassandraSession),
				userService.UnaryRegisterUserUseCaseContext(),
			),
		),
	}
	srv.grpcServer = grpc.NewServer(opts...)
	srv.listener = &listen
	srv.RegisterService()
}
