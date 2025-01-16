package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/jpradass/cerberus/db"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Creates Cerberus den",
	Long:  "Creates the db where all stuff will be stored",
	Run: func(cmd *cobra.Command, args []string) {
		if !db.CheckDenExistence() {
			if err := db.CreateDen(); err != nil {
				fmt.Printf("don't know where cerberus lives: %v\n", err)
			}
		}
		fmt.Printf("Cerberus den's at: %s\n", db.GetCerberusDenPath())
	},
}
