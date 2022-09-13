package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

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
	logger             *zap.Logger
	tracer             opentracing.Tracer
	consulClient       *api.Client
	rabbitMQConnection *amqp.Connection
)

func init() {
	consulClient = NewConsulClient()

	logger = NewLogger()
	tracer = NewTracer()
	rabbitMQConnection = NewAmqp()
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
		panic(err)
	}
	return
}

func NewTracer() opentracing.Tracer {
	cfg, err := config.FromEnv()
	if err != nil {
		panic(fmt.Sprintf("ERROR: failed to read config from env vars: %v\n", err))
	}
	tc, _, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tc
}

func NewConsulClient() *api.Client {
	port, _ := strconv.Atoi(os.Getenv("HTTP_PORT"))
	address := fmt.Sprintf("http://%s:%s/check", os.Getenv("HTTP_HOST"), os.Getenv("HTTP_PORT"))
	defaultConfig := api.DefaultConfig()
	client, err := api.NewClient(defaultConfig)
	if err != nil {
		log.Fatalln(err)
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
		log.Fatalln("error when register service")
	}

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
		log.Fatal(err)
	}
	return conn
}
