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
		cmd.Printf("key: %s\n", key)

		value, isPath, err := db.GetFromDen(key)
		if err != nil {
			cmd.PrintErrf("cerberus is confused: %v\n", err)
			return
		}

		if isPath == 1 {
			content, err := fs.ReadContent(value)
			if err != nil {
				cmd.PrintErrf("cerberus is confused: %v\n", err)
				return
			}

			cmd.Printf("%x\n", content)
			return
		}

		cmd.Printf("%s\n", value)
	},
}
