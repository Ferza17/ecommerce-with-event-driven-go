package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/Ferza17/event-driven-api-gateway/rest"
)

var restCommand = &cobra.Command{
	Use: "rest",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Starting REST API server ...")
		rest.NewServer(
			os.Getenv("CODENAME"),
			os.Getenv("HTTP_HOST"),
			os.Getenv("HTTP_PORT"),
			rest.NewConsulClient(consulClient),
			rest.NewTracer(tracer),
			rest.NewLogger(logger),
			rest.NewAmqpConn(amqpConn),
		).Serve()
	},
}
