package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Removes a Task.",
	Long:  `Removes a Task.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return persistentPreRun()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		taskId, err := cmd.Flags().GetUint("task-id")
		if err != nil {
			return err
		}
		exists, err := tr.TaskExists(taskId)
		if err != nil {
			return err
		}
		if !exists {
			return fmt.Errorf("'%d' does not exist", taskId)
		}

		err = tr.DeleteTaskById(taskId)
		if err != nil {
			return err
		}

		fmt.Println("Task removed")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
	rmCmd.Flags().UintP("task-id", "t", 0, "Id of the Task to remove.")
	rmCmd.MarkFlagRequired("task-id")
}
