package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Cerberus",
	Long:  `All software has versions. This is Cerberus'`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Cerberus Key Value Keeper v0.1 -- HEAD")
	},
}
