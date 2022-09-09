package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var runCommand = &cobra.Command{
	Use: "amqp",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Starting AMQP CONSUMER Server")
	},
}
