package cmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/Bilou4/godo/model"
	"github.com/spf13/cobra"
)

const dueDateFormat string = "2006-01-02 15:04"

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Add a new Task.",
	Long:    `Add a new Task in a List.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return persistentPreRun()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		taskName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		if taskName == "" {
			return errors.New("the name of your task cannot be empty")
		}
		listId, err := cmd.Flags().GetUint("list-id")
		if err != nil {
			return err
		}
		exists, err := lr.ListExists(listId)
		if err != nil {
			return err
		}
		if !exists {
			return fmt.Errorf("list id '%d' does not exist", listId)
		}

		priority, err := cmd.Flags().GetString("priority")
		if err != nil {
			return err
		}

		dueDateString, err := cmd.Flags().GetString("due-date")
		if err != nil {
			return err
		}
		var dueDate time.Time
		if dueDateString != "" {
			dueDate, err = time.Parse(dueDateFormat, dueDateString)
			if err != nil {
				return err
			}
		}

		t, err := tr.CreateTask(listId, taskName, dueDate, model.Priority(priority))
		if err != nil {
			return err
		}
		fmt.Println("New Task added.")
		fmt.Println(t)
		return nil

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().UintP("list-id", "l", 0, "Id of the List where to add your Task.")
	addCmd.Flags().StringP("name", "n", "", "Name of the Task.")
	addCmd.Flags().StringP("priority", "p", "LOW", "Priority of your Task [HIGH, MEDIUM, LOW].")
	addCmd.Flags().StringP("due-date", "d", "", fmt.Sprintf("Due date for the Task (format: %s).", dueDateFormat))

	addCmd.MarkFlagRequired("list-id")
	addCmd.MarkFlagRequired("name")
}
