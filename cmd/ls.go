package cmd

import (
	"fmt"
	"time"

	"github.com/Bilou4/godo/model"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"l"},
	Short:   "Prints Tasks.",
	Long:    `Prints Tasks or only done Tasks.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return persistentPreRun()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		onlyDone, err := cmd.Flags().GetBool("done")
		if err != nil {
			return err
		}

		listIdIsSet := cmd.Flags().Lookup("list-id").Changed
		if listIdIsSet {
			listId, err := cmd.Flags().GetUint("list-id")
			if err != nil {
				return err
			}
			list, err := lr.GetListById(listId)
			if err != nil {
				return err
			}
			showTasksForList(list, onlyDone)

		} else {
			// show all lists and tasks
			lists, err := lr.GetAllLists()
			if err != nil {
				return err
			}
			if len(lists) > 0 {
				for _, l := range lists {
					showTasksForList(l, onlyDone)
				}
			} else {
				fmt.Println("Nothing to show.")
			}
		}
		return nil
	},
}

func getCreatedMessage(createdAt time.Time) string {
	diff := time.Since(createdAt)
	switch diffHours := int(diff.Hours()); {
	case diffHours > 48:
		diffHours = diffHours / 24
		return fmt.Sprintf("%d days ago", diffHours)
	case diffHours == 1:
		return fmt.Sprintf("%d hour ago", diffHours)
	case diffHours == 0:
		switch diffMinutes := int(diff.Minutes()); {
		case diffMinutes > 1:
			return fmt.Sprintf("%d minutes ago", diffMinutes)
		default:
			return fmt.Sprintf("%d minute ago", diffMinutes)
		}
	default:
		return fmt.Sprintf("%d hours ago", diffHours)
	}
}

func getPrettyTable() table.Writer {
	prettyTable := table.NewWriter()
	prettyTable.Style().Options.DrawBorder = true
	prettyTable.Style().Options.SeparateRows = true
	prettyTable.SetStyle(table.StyleLight)
	prettyTable.Style().Title.Align = text.AlignCenter
	prettyTable.AppendRow(table.Row{"ID", "Name", "Due Date", "Priority", "Done", "Created"})
	prettyTable.SetColumnConfigs(
		[]table.ColumnConfig{
			{Number: 1, Align: text.AlignCenter},
			{Number: 2, Align: text.AlignCenter},
			{Number: 3, Align: text.AlignCenter},
			{Number: 4, Align: text.AlignCenter},
			{Number: 5, Align: text.AlignCenter},
			{Number: 6, Align: text.AlignCenter},
		},
	)
	return prettyTable
}

func showTasksForList(l *model.List, onlyDone bool) error {
	var err error
	var tasks []model.Task
	if onlyDone {
		tasks, err = tr.GetDoneTasksByListId(l.ID)
	} else {
		tasks, err = tr.GetTasksByListId(l.ID)
	}
	if err != nil {
		return err
	}

	if len(tasks) > 0 {
		prettyTable := getPrettyTable()
		prettyTable.SetTitle(fmt.Sprintf("ID:%d  - Name: %s", l.ID, l.Name))

		for _, t := range tasks {
			var dueDateStr string = "None"
			if !t.DueDate.IsZero() {
				dueDateStr = t.DueDate.Format("2006-01-02 15:04")
			}
			var done string = "🗹"
			if !t.Done {
				done = "□"
			}
			created := getCreatedMessage(t.CreatedAt)
			prettyTable.AppendRow(table.Row{t.ID, t.Name, dueDateStr, t.Priority, done, created})
		}

		fmt.Println(prettyTable.Render())
		fmt.Println()
	} else {
		fmt.Printf("%s: no tasks to show.\n", l.Name)
	}
	return nil
}

func init() {
	rootCmd.AddCommand(lsCmd)

	lsCmd.Flags().Bool("done", false, "Show only done Tasks.")
	lsCmd.Flags().UintP("list-id", "l", 0, "Id of the List you want to see the tasks.")
}
