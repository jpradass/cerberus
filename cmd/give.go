package cmd

import (
	"github.com/spf13/cobra"

	"github.com/jpradass/cerberus/db"
	"github.com/jpradass/cerberus/fs"
	dbmodels "github.com/jpradass/cerberus/models/db"
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

		entry := new(dbmodels.Entry)
		entry.Key, entry.Value, entry.IsPath, entry.IsBinary = args[0], args[1], false, false
		// cmd.Printf("key: %s, value: %s\n", key, value)

		if fs.IsPath(entry.Value) {
			entry.IsPath = true
		}

		if err := db.SaveInDen(entry); err != nil {
			cmd.PrintErrf("cerberus is confused: %v\n", err)
			return
		}
		cmd.Printf("key given to cerberus!\n")
	},
}
