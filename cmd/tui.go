package cmd

import (
	"fmt"
	"os"

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

	Run: func(cmd *cobra.Command, args []string) {
		mainModel := tui.NewModel(tr, lr)
		_, err := tea.LogToFile("app.log", "")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		p := tea.NewProgram(mainModel)
		if err := p.Start(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}
