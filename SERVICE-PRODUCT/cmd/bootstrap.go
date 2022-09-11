package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-redis/redis/v8"
	"github.com/gocql/gocql"
	"github.com/hashicorp/consul/api"
	_ "github.com/joho/godotenv/autoload"
	"github.com/opentracing/opentracing-go"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	consulClient     *api.Client
	esClient         *elasticsearch.Client
	cassandraSession *gocql.Session
	logger           *zap.Logger
	tracer           opentracing.Tracer
	amqpConn         *amqp.Connection
	redisClient      *redis.Client
)

func init() {
	consulClient = NewConsulClient()
	esClient = NewElasticsearchClient()
	logger = NewLogger()
	tracer = NewTracer()
	amqpConn = NewAmqp()
	redisClient = NewRedisClient()
	cassandraSession = NewCassandraSession()
}

func NewLogger() (logger *zap.Logger) {
	var err error
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	if logger, err = logConfig.Build(); err != nil {
		log.Fatalf("error when register logger: %v\n", err)
	}
	log.Println("logger registered")
	return
}

func NewTracer() opentracing.Tracer {
	cfg, err := config.FromEnv()
	if err != nil {
		log.Fatalf("ERROR: failed to read config from env vars: %v\n", err)
	}
	tc, _, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		log.Fatalf("ERROR: cannot init Jaeger: %v\n", err)
	}
	log.Println("tracer connected")
	return tc
}

func NewConsulClient() *api.Client {
	port, _ := strconv.Atoi(os.Getenv("HTTP_PORT"))
	address := fmt.Sprintf("http://%s:%s/check", os.Getenv("HTTP_HOST"), os.Getenv("HTTP_PORT"))
	defaultConfig := api.DefaultConfig()
	client, err := api.NewClient(defaultConfig)
	if err != nil {
		log.Fatalf("error when connect to consul: %v\n", err)
	}
	serviceRegistration := &api.AgentServiceRegistration{
		Name:    os.Getenv("CODENAME"),
		Address: address,
		Port:    port,
		Check: &api.AgentServiceCheck{
			HTTP:     address,
			Interval: "10s",
			Timeout:  "30s",
		},
	}
	if err = client.Agent().ServiceRegister(serviceRegistration); err != nil {
		log.Fatalf("error when register service: %v\n", err)
	}
	log.Println("consul connected")
	return client
}

func NewAmqp() *amqp.Connection {
	conn, err := amqp.Dial(
		fmt.Sprintf("amqp://%s:%s@%s:%s/",
			os.Getenv("RABBITMQ_USERNAME"),
			os.Getenv("RABBITMQ_PASSWORD"),
			os.Getenv("RABBITMQ_HOST"),
			os.Getenv("RABBITMQ_PORT"),
		))
	if err != nil {
		log.Fatalf("error while connecting to RabbitMQ: %v\n", err)
	}
	log.Println("RabbitMQ connected")
	return conn
}

func NewRedisClient() *redis.Client {
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Fatalln("mongoClient: env REDIS_DB value should be an integer greater than 0")
	}
	client := redis.NewClient(
		&redis.Options{
			Addr:     os.Getenv("REDIS_ADDRESS"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       db,
		},
	)
	// Make sure that connection insurable
	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("error while connecting to redis: %v\n", err)
	}
	log.Println("redis connected")
	return client
}

func NewCassandraSession() *gocql.Session {
	cluster := gocql.NewCluster(os.Getenv("CASSANDRA_HOST"))
	cluster.Keyspace = os.Getenv("CASSANDRA_KEYSPACE")
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Could not connect to CassandraDB: %v\n", err)
	}
	log.Println("CassandraDB Connected")
	return session
}

func NewElasticsearchClient() *elasticsearch.Client {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			os.Getenv("ELASTICSEARCH_ADDRESS"),
		},
	})

	if err != nil {
		log.Fatalf("Could not connect to Elasticsearch: %v\n", err)
	}

	// Make sure that connection insurable
	_, err = client.Ping()
	if err != nil {
		log.Fatalf("error while connecting to elasticsearch: %v\n", err)
	}

	log.Println("Elasticsearch Connected")
	return client
}
