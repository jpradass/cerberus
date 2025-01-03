package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cerberus",
	Short: "Cerberus is a gatekeeper for everything you want to keep",
	Long:  "Cerberus is a Key Value application that keeps everything inside a local database encrypted.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from cerberus")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
