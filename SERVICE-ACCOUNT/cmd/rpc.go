package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rpcCommand = &cobra.Command{
	Use: "rpc",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Starting RPC Server")
	},
}
