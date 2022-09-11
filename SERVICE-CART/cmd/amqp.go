package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/Ferza17/event-driven-cart-service/rabbitmq"
)

var rabbitMQCommand = &cobra.Command{
	Use: "rabbitmq",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("========= Starting RabbitMQ CONSUMER Server =========")
		rabbitmq.NewServer(
			rabbitmq.NewMongoClient(mongoClient),
			rabbitmq.NewLogger(logger),
			rabbitmq.NewRabbitMQConnection(rabbitMQConnection),
			rabbitmq.NewTracer(tracer),
			rabbitmq.NewRedisClient(redisClient),
			rabbitmq.NewCassandraSession(cassandraSession),
		).Serve()
	},
}
