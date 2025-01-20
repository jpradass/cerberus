package cmd

import (
	"github.com/spf13/cobra"

	"github.com/jpradass/cerberus/db"
	"github.com/jpradass/cerberus/fs"
)

func init() {
	rootCmd.AddCommand(takeCmd)
}

var takeCmd = &cobra.Command{
	Use:   "take",
	Short: "Takes object from Cerberus den",
	Long:  "Gets an object from the database using its key to find it",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Println("it is needed a key to search into cerberus den")
			return
		}

		key := args[0]
		// cmd.Printf("key: %s\n", key)

		entry, err := db.GetFromDen(key)
		if err != nil {
			cmd.PrintErrf("cerberus is confused: %v\n", err)
			return
		}

		if entry.IsPath {
			content, err := fs.ReadContent(entry.Value)
			if err != nil {
				cmd.PrintErrf("cerberus is confused: %v\n", err)
				return
			}

			cmd.Printf("%x", content)
			return
		}

		cmd.Printf("%s", entry.Value)
	},
}
