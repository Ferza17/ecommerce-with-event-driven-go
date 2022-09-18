package graphql

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/go-chi/chi/v5"
	chim "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/gorilla/websocket"
	"github.com/hashicorp/consul/api"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
	l "github.com/treastech/logger"
	"go.uber.org/zap"

	"github.com/Ferza17/event-driven-api-gateway/middleware"
	"github.com/Ferza17/event-driven-api-gateway/model/graph/generated"
	"github.com/Ferza17/event-driven-api-gateway/model/graph/resolver"
	"github.com/Ferza17/event-driven-api-gateway/model/pb"
	"github.com/Ferza17/event-driven-api-gateway/module/cart"
	"github.com/Ferza17/event-driven-api-gateway/module/user"
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
		graphQLServer            *handler.Server
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
	// GraphQL Server
	c := generated.Config{
		Resolvers: &resolver.Resolver{},
		Directives: generated.DirectiveRoot{
			Jwt: middleware.DirectiveJwtRequired,
		},
	}

	gqlServer := handler.NewDefaultServer(generated.NewExecutableSchema(c))
	gqlServer.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		KeepAlivePingInterval: 10 * time.Second,
	})
	gqlServer.AddTransport(transport.Options{})
	gqlServer.AddTransport(transport.GET{})
	gqlServer.AddTransport(transport.MultipartForm{})
	gqlServer.SetQueryCache(lru.New(1000))
	gqlServer.Use(extension.Introspection{})
	gqlServer.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})
	srv.graphQLServer = gqlServer

	// HTTP Server
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
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "CONNECT", "TRACE", "HEAD", "PATCH"},
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
		user.RegisterUserUseCaseHTTPContext(),
		cart.RegisterCartUseCaseHTTPContext(),
		chim.Heartbeat("/ping"),
	)
	routes(r, srv.graphQLServer)
	return r
}
