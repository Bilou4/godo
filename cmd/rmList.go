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
		force, err := cmd.Flags().GetBool("force")
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

		hasTasks, err := lr.ListHasTasks(listId)
		if err != nil {
			return err
		}

		if hasTasks {
			if force {
				err = tr.DeleteTasks(listId)
				if err != nil {
					return err
				}
			} else {
				fmt.Println("List is not empty.")
				return nil
			}
		}
		return lr.DeleteList(listId)
	},
}

func init() {
	rootCmd.AddCommand(rmListCmd)
	rmListCmd.Flags().Bool("force", false, "Force deletion of the List and all Tasks.")
	rmListCmd.Flags().UintP("list-id", "l", 0, "Id of the List to remove.")
	rmListCmd.MarkFlagRequired("list-id")

}
