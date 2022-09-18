package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/Ferza17/event-driven-product-service/grpc"
)

var rpcCommand = &cobra.Command{
	Use: "grpc",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("========== Starting RPC Server ==========")
		grpc.NewServer(
			os.Getenv("RPC_HOST"),
			os.Getenv("RPC_PORT"),
			grpc.NewTracer(tracer),
			grpc.NewLogger(logger),
			grpc.NewRedisClient(redisClient),
			grpc.NewRabbitMQConnection(rabbitMQConnection),
			grpc.NewElasticsearchClient(esClient),
			grpc.NewPostgresClient(postgresSQlClient),
			grpc.NewCassandraSession(cassandraSession),
		).Serve()
	},
}
