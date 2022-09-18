package cmd

import (
	"log"
	"os"
	"sync"

	"github.com/spf13/cobra"

	"github.com/Ferza17/event-driven-account-service/grpc"
	"github.com/Ferza17/event-driven-account-service/rabbitmq"
)

var runCommand = &cobra.Command{
	Use: "run",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("========= Starting RabbitMQ CONSUMER Server & gRPC Server =========")
		wg := sync.WaitGroup{}
		wg.Add(2)
		go func() {
			defer wg.Done()
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
			defer wg.Done()
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
		wg.Wait()
	},
}
