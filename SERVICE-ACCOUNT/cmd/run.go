package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/Ferza17/event-driven-account-service/grpc"
	"github.com/Ferza17/event-driven-account-service/rabbitmq"
)

var runCommand = &cobra.Command{
	Use: "run",
	Run: func(cmd *cobra.Command, args []string) {
		chClose := make(chan os.Signal, 2)
		signal.Notify(chClose, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			log.Println("========= Starting RabbitMQ CONSUMER Server =========")
			rabbitmq.NewServer(
				rabbitmq.NewMongoClient(mongoClient),
				rabbitmq.NewLogger(logger),
				rabbitmq.NewRabbitMQConnection(rabbitMQConnection),
				rabbitmq.NewTracer(tracer),
				rabbitmq.NewRedisClient(redisClient),
				rabbitmq.NewCassandraSession(cassandraSession),
			).Serve()
		}()
		go func() {
			log.Println("========= Starting gRPC Server =========")
			grpc.NewServer(
				os.Getenv("RPC_HOST"),
				os.Getenv("RPC_PORT"),
				grpc.NewTracer(tracer),
				grpc.NewMongoClient(mongoClient),
				grpc.NewLogger(logger),
				grpc.NewRedisClient(redisClient),
				grpc.NewRabbitMQConnection(rabbitMQConnection),
				grpc.NewCassandraSession(cassandraSession),
			).Serve()
		}()

		<-chClose
		if err := Shutdown(context.Background()); err != nil {
			log.Fatalln(err)
			return
		}
		log.Println("Exit...")
	},
}
