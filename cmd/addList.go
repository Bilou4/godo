package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// addListCmd represents the addList command
var addListCmd = &cobra.Command{
	Use:     "addList",
	Aliases: []string{"al"},
	Short:   "Create a new List.",
	Long:    `Create a new List.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return persistentPreRun()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		listName, err := cmd.Flags().GetString("list-name")
		if err != nil {
			return err
		}
		if listName == "" {
			return errors.New("the name of your list cannot be empty")
		}
		_, err = lr.CreateList(listName)
		if err != nil {
			return err
		}
		fmt.Printf("The new list '%s' has been added.\n", listName)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addListCmd)

	addListCmd.Flags().StringP("list-name", "n", "", "Name of the List to create.")

	addListCmd.MarkFlagRequired("list-name")
}
