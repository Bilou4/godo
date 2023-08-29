package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// TODO uncheck
// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:     "check",
	Aliases: []string{"c"},
	Short:   "Mark a Task as done.",
	Long:    `Mark a Task as done.`,
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
		err = tr.UpdateTaskIsDone(taskId, true)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	checkCmd.Flags().UintP("task-id", "t", 0, "Id of the Task to update.")
	checkCmd.MarkFlagRequired("task-id")
}
