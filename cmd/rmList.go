package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rmListCmd represents the rmList command
var rmListCmd = &cobra.Command{
	Use:   "rmList",
	Short: "Removes an empty List.",
	Long:  `Removes an empty List.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return persistentPreRun()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		listId, err := cmd.Flags().GetUint("list-id")
		if err != nil {
			return err
		}

		exists, err := lr.ListExists(listId)
		if err != nil {
			return err
		}
		if !exists {
			return fmt.Errorf("'%d' does not exist", listId)
		}

		err = tr.DeleteTasks(listId)
		if err != nil {
			return err
		}
		err = lr.DeleteList(listId)
		if err != nil {
			return err
		}
		return nil

	},
}

func init() {
	rootCmd.AddCommand(rmListCmd)
	// TODO add a 'force' flag to remove all tasks and the list
	rmListCmd.Flags().UintP("list-id", "l", 0, "Id of the List to remove.")
	rmListCmd.MarkFlagRequired("list-id")

}
