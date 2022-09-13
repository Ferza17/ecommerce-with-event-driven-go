package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/Ferza17/event-driven-api-gateway/graphql"
)

var graphqlCommand = &cobra.Command{
	Use: "schema",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Starting GraphQL API server ...")
		graphql.NewServer(
			os.Getenv("CODENAME"),
			os.Getenv("HTTP_HOST"),
			os.Getenv("HTTP_PORT"),
			graphql.NewConsulClient(consulClient),
			graphql.NewTracer(tracer),
			graphql.NewLogger(logger),
			graphql.NewRabbitMQConnection(rabbitMQConnection),
		).Serve()
	},
}
