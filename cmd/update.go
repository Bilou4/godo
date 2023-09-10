package cmd

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Bilou4/godo/model"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"u"},
	Short:   "Updates a Task.",
	Long:    `Updates a Task with fields to change.`,
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

		var changes bool
		newTaskName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		taskNameIsSet := cmd.Flags().Lookup("name").Changed
		if taskNameIsSet {
			if newTaskName == "" {
				return errors.New("the name of your task cannot be empty")
			} else {
				err = tr.UpdateTaskName(taskId, newTaskName)
				if err != nil {
					return err
				}
				changes = true
			}
		}

		priorityString, err := cmd.Flags().GetString("priority")
		if err != nil {
			return err
		}
		priorityString = strings.ToUpper(priorityString)
		priorityIsSet := cmd.Flags().Lookup("priority").Changed
		if priorityIsSet {
			if priorityString == "" {
				return errors.New("the priority of your task cannot be empty")
			} else {
				priority := model.Priority(priorityString)
				if err := priority.IsValid(); err != nil {
					return err
				}
				err = tr.UpdateTaskPriority(taskId, priority)
				if err != nil {
					return err
				}
				changes = true
			}
		}

		dueDateString, err := cmd.Flags().GetString("due-date")
		if err != nil {
			return err
		}
		dueDateIsSet := cmd.Flags().Lookup("due-date").Changed
		if dueDateIsSet {
			if dueDateString == "" {
				return errors.New("the due date of your task cannot be empty")
			} else {
				dueDate, err := time.Parse(dueDateFormat, dueDateString)
				if err != nil {
					return err
				}

				if dueDate.Before(time.Now()) {
					fmt.Println("WARNING: your due date is in the past")
				}

				err = tr.UpdateTaskDueDate(taskId, dueDate)
				if err != nil {
					return err
				}
				changes = true
			}
		}

		if changes {
			fmt.Println("Task updated")
			t, err := tr.GetTaskById(taskId)
			if err != nil {
				return err
			}
			fmt.Println(t)
		} else {
			fmt.Println("Nothing to do.")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().UintP("task-id", "t", 0, "Id of the Task to update.")
	updateCmd.Flags().StringP("name", "n", "", "New name of the Task.")
	updateCmd.Flags().StringP("priority", "p", "", "Priority of your Task [HIGH, MEDIUM, LOW].")
	updateCmd.Flags().StringP("due-date", "d", "", fmt.Sprintf("Due date for the Task (format: %s).", dueDateFormat))

	updateCmd.MarkFlagRequired("task-id")
}
