package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/Ferza17/event-driven-api-gateway/graphql"
)

var graphqlCommand = &cobra.Command{
	Use: "graphql",
	Run: func(cmd *cobra.Command, args []string) {
		chClose := make(chan os.Signal, 2)
		signal.Notify(chClose, syscall.SIGINT, syscall.SIGTERM)
		go func() {
			log.Println("Starting GraphQL API server ...")
			graphql.NewServer(
				os.Getenv("CODENAME"),
				os.Getenv("HTTP_HOST"),
				os.Getenv("HTTP_PORT"),
				graphql.NewConsulClient(consulClient),
				graphql.NewTracer(tracer),
				graphql.NewLogger(logger),
				graphql.NewRabbitMQConnection(rabbitMQConnection),
				graphql.NewUserServiceGrpcClientConnection(userServiceGrpcClient),
				graphql.NewProductServiceGrpcClientConnection(productServiceGrpcClient),
				graphql.NewCartServiceGrpcClientConnection(cartServiceGrpcClient),
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
