package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/Ferza17/event-driven-account-service/amqp"
)

var runCommand = &cobra.Command{
	Use: "amqp",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("========= Starting AMQP CONSUMER Server =========")
		amqp.NewServer(
			amqp.NewDB(db),
			amqp.NewLogger(logger),
			amqp.NewAMQP(amqpConn),
			amqp.NewTracer(tracer),
			amqp.NewRedisClient(redisClient),
		).Serve()
	},
}
