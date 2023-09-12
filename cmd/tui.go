package cmd

import (
	"errors"

	"github.com/Bilou4/godo/tui"
	"github.com/spf13/cobra"

	_ "github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

// tuiCmd represents the tui command
var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "",
	Long:  ``,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return persistentPreRun()
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		listPerPage, err := cmd.Flags().GetInt("list-per-page")
		if err != nil {
			return err
		}
		if listPerPage < 1 {
			return errors.New("the list-per-page flag cannot be less than 1.")
		}
		mainModel := tui.NewModel(tr, lr, appConfig.Tui, listPerPage)
		_, err = tea.LogToFile("app.log", "")
		if err != nil {
			return err
		}
		p := tea.NewProgram(mainModel)
		if err := p.Start(); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
	tuiCmd.Flags().IntP("list-per-page", "l", 3, "Number of lists to display per page.")
}
