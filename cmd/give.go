package cmd

import (
	"github.com/spf13/cobra"

	"github.com/jpradass/cerberus/db"
	"github.com/jpradass/cerberus/fs"
)

func init() {
	rootCmd.AddCommand(giveCmd)
}

var giveCmd = &cobra.Command{
	Use:   "give",
	Short: "Gives object to Cerberus",
	Long:  "Adds a new object into the database",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			cmd.Println("it is needed two arguments as key - value pair to store in cerberus den")
			return
		}

		key, value, isPath := args[0], args[1], 0
		cmd.Printf("key: %s, value: %s\n", key, value)

		if fs.IsPath(value) {
			isPath = 1
		}

		if err := db.SaveInDen(key, value, isPath); err != nil {
			cmd.PrintErrf("cerberus is confused: %v\n", err)
			return
		}
	},
}
