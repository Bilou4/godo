package cmd

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

// lsListCmd represents the lsList command
var lsListCmd = &cobra.Command{
	Use:     "lsList",
	Aliases: []string{"ll"},
	Short:   "Prints existing Lists.",
	Long:    `Prints existing Lists.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return persistentPreRun()
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		lists, err := lr.GetAllLists()
		if err != nil {
			return err
		}

		if len(lists) > 0 {
			prettyTable := table.NewWriter()
			prettyTable.Style().Options.DrawBorder = true
			prettyTable.Style().Options.SeparateRows = true
			prettyTable.SetStyle(table.StyleLight)
			prettyTable.SetTitle("Lists")
			prettyTable.Style().Title.Align = text.AlignCenter
			prettyTable.AppendRow(table.Row{"ID", "Name"})
			prettyTable.SetColumnConfigs(
				[]table.ColumnConfig{
					{Number: 1, Align: text.AlignCenter},
					{Number: 2, Align: text.AlignCenter},
				},
			)
			for _, l := range lists {
				prettyTable.AppendRow(table.Row{l.ID, l.Name})
			}
			fmt.Println(prettyTable.Render())

		} else {
			fmt.Println("Nothing to show.")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(lsListCmd)
}
