package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new key-value pair",
	Long:  "Adds a new key-value pair into the database",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("args: %v", args)
	},
}
