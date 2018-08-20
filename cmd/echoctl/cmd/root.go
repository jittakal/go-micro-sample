package cmd

import (
	"fmt"
	"os"

	"github.com/jittakal/go-micro-sample/cmd/echoctl/client"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "echoctl",
		Short: "echoctl cli application for echo service",
		Long:  `echoctl cli application for echo service`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("echoctl cli application!")
		},
	}

	echoCmd = &cobra.Command{
		Use:   "echo",
		Short: "invoke Echo service",
		Long:  "It invokes Echo service using gRPC client library",
		Run: func(cmd *cobra.Command, args []string) {
			messageFlag := cmd.Flag("message")
			err := client.DoEcho(messageFlag.Value.String())

			if err != nil {
				fmt.Println(err)
			}
		},
	}
)

func init() {
	var message = ""
	echoCmd.Flags().StringVarP(&message, "message", "m", "", "Message to be echoed")
	echoCmd.MarkFlagRequired("message")

	rootCmd.AddCommand(echoCmd)
}

// Execute main executor function of echoctl cli application
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
