package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// clearListCmd represents the clearList command
var clearListCmd = &cobra.Command{
	Use:     "clearList",
	Aliases: []string{"cl"},
	Short:   "Removes Tasks from a List.",
	Long:    `Removes Tasks or only done Tasks from a List.`,
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

		onlyDone, err := cmd.Flags().GetBool("done")
		if err != nil {
			return err
		}

		if onlyDone {
			err = tr.DeleteDoneTasks(listId)
		} else {
			err = tr.DeleteTasks(listId)
		}
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(clearListCmd)
	clearListCmd.Flags().UintP("list-id", "l", 0, "Id of the List to clear")
	clearListCmd.Flags().Bool("done", false, "Remove only done Tasks")

	clearListCmd.MarkFlagRequired("list-id")
}
