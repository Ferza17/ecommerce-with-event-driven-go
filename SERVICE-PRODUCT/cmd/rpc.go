package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/Ferza17/event-driven-product-service/rpc"
)

var rpcCommand = &cobra.Command{
	Use: "rpc",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("========== Starting RPC Server ==========")
		rpc.NewServer(
			os.Getenv("RPC_HOST"),
			os.Getenv("RPC_PORT"),
			rpc.NewTracer(tracer),
			rpc.NewLogger(logger),
			rpc.NewRedisClient(redisClient),
			rpc.NewRabbitMQConnection(amqpConn),
			rpc.NewElasticsearchClient(esClient),
			rpc.NewCassandraSession(cassandraSession),
		).Serve()
	},
}