package cmd

import (
	"log"
	"os"
	"sync"

	"github.com/spf13/cobra"

	"github.com/Ferza17/event-driven-product-service/grpc"
)

var runCommand = &cobra.Command{
	Use: "run",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("========== Starting RPC Server ==========")
		wg := sync.WaitGroup{}
		wg.Add(2)
		go func() {
			defer wg.Done()
			grpc.NewServer(
				os.Getenv("RPC_HOST"),
				os.Getenv("RPC_PORT"),
				grpc.NewTracer(tracer),
				grpc.NewLogger(logger),
				grpc.NewRedisClient(redisClient),
				grpc.NewRabbitMQConnection(amqpConn),
				grpc.NewElasticsearchClient(esClient),
				grpc.NewPostgresClient(postgresSQlClient),
				grpc.NewCassandraSession(cassandraSession),
			).Serve()
		}()
		go func() {
			//TODO:
			// 1. Add Consumer Server
		}()
		wg.Wait()
	},
}
