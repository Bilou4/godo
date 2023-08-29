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

		lists, err := lr.GetAllLists()
		if err != nil {
			return err
		}
		if len(lists) > 0 {
			for _, l := range lists {
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
					prettyTable := table.NewWriter()
					prettyTable.Style().Options.DrawBorder = true
					prettyTable.Style().Options.SeparateRows = true
					prettyTable.SetStyle(table.StyleLight)
					prettyTable.SetTitle(fmt.Sprintf("ID:%d  - Name: %s", l.ID, l.Name))
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
					fmt.Println("No tasks to show.")
				}
			}
		} else {
			fmt.Println("Nothing to show.")
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
	case diffHours > 24:
		return "yesterday"
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

func init() {
	rootCmd.AddCommand(lsCmd)

	// TODO add an optional list-id flag to show only tasks from a list
	lsCmd.Flags().Bool("done", false, "Show only done Tasks.")
}
