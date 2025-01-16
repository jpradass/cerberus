package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version    string = "v0.1"
	lastCommit string = "9a9995adac"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Cerberus",
	Long:  `All software has versions. This is Cerberus'`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Cerberus Key Value Keeper %s -- %s\n", version, lastCommit)
	},
}
