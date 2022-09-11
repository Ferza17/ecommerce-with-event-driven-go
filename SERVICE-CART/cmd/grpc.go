package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/Ferza17/event-driven-cart-service/grpc"
)

var grpcCommand = &cobra.Command{
	Use: "grpc",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("========== Starting gRPC Server ==========")
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
	},
}
