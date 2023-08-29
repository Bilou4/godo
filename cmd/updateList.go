package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// updateListCmd represents the updateList command
var updateListCmd = &cobra.Command{
	Use:     "updateList",
	Aliases: []string{"ul"},
	Short:   "Updates a List.",
	Long:    `Updates the name of a List.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return persistentPreRun()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		newName, err := cmd.Flags().GetString("new-name")
		if err != nil {
			return err
		}
		if newName == "" {
			return errors.New("the new name of your list cannot be empty")
		}
		listId, err := cmd.Flags().GetUint("list-id")
		if err != nil {
			return err
		}

		err = lr.UpdateListName(listId, newName)
		if err != nil {
			return err
		}
		fmt.Printf("List name updated: '%s'\n", newName)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateListCmd)

	updateListCmd.Flags().StringP("new-name", "n", "", "New name of the List.")
	updateListCmd.Flags().UintP("list-id", "l", 0, "Id of the List to update.")

	updateListCmd.MarkFlagRequired("new-name")
	updateListCmd.MarkFlagRequired("list-id")
}
