package graphql

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chim "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/hashicorp/consul/api"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
	l "github.com/treastech/logger"
	"go.uber.org/zap"

	"github.com/Ferza17/event-driven-api-gateway/middleware"
	"github.com/Ferza17/event-driven-api-gateway/model/pb"
	"github.com/Ferza17/event-driven-api-gateway/model/schema/cart"
	"github.com/Ferza17/event-driven-api-gateway/model/schema/product"
	"github.com/Ferza17/event-driven-api-gateway/model/schema/user"
	"github.com/Ferza17/event-driven-api-gateway/utils"
)

type (
	Server struct {
		codename                 string
		host                     string
		port                     string
		rabbitMQConnection       *amqp.Connection
		logger                   *zap.Logger
		router                   *chi.Mux
		tracer                   opentracing.Tracer
		httpServer               *http.Server
		consulClient             *api.Client
		userServiceGrpcClient    pb.UserServiceClient
		productServiceGrpcClient pb.ProductServiceClient
		cartServiceGrpcClient    pb.CartServiceClient
	}
	Option func(s *Server)
)

func NewServer(codename, host, address string, option ...Option) *Server {
	s := &Server{
		codename: codename,
		host:     host,
		port:     address,
	}
	for _, o := range option {
		o(s)
	}
	s.setup()
	return s
}

func (srv *Server) Serve() {
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		srv.logger.Info(fmt.Sprintf("%s %s", method, route))
		return nil
	}
	if err := chi.Walk(srv.router, walkFunc); err != nil {
		log.Panicln(errors.Cause(err))
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", srv.host, srv.port), srv.router))
}

func (srv *Server) setup() {
	srv.router = srv.routes()
	srv.httpServer = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", srv.host, srv.port),
		Handler: srv.router,
	}
}

func (srv *Server) routes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		}),
		chim.RequestID,
		chim.RealIP,
		chim.Recoverer,
		chim.NoCache,
		render.SetContentType(render.ContentTypeJSON),
		l.Logger(srv.logger),
		middleware.Host(srv.codename),
		middleware.Header(),
		middleware.RegisterTracerHTTPContext(srv.tracer),
		middleware.RegisterRabbitMQAmqpHTTPContext(srv.rabbitMQConnection),
		middleware.RegisterUserServiceGrpcClientHttpContext(srv.userServiceGrpcClient),
		middleware.RegisterProductServiceGrpcClientHttpContext(srv.productServiceGrpcClient),
		middleware.RegisterCartServiceGrpcClientHttpContext(srv.cartServiceGrpcClient),
		middleware.RegisterSchemaConfigHTTPContext(user.UserSchemaConfig, utils.UserSchemaConfigContextKey),
		middleware.RegisterSchemaConfigHTTPContext(user.AuthSchemaConfig, utils.AuthSchemaConfigContextKey),
		middleware.RegisterSchemaConfigHTTPContext(cart.CartSchemaConfig, utils.CartSchemaConfigContextKey),
		middleware.RegisterSchemaConfigHTTPContext(product.ProductSchemaConfig, utils.ProductSchemaConfigContextKey),
		chim.Heartbeat("/ping"),
	)
	// GraphQL
	routes(r)
	return r
}
