package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// removeDatabaseCmd represents the removeDatabase command
var removeDatabaseCmd = &cobra.Command{
	Use:   "cleanApp",
	Short: "Clean the Application.",
	Long:  `Clean the Application by removing the database and the configuration file.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		err := initAppConfig()
		if err != nil {
			return fmt.Errorf("nothing to do, the application was not initialized")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		// remove the database and the configuration file
		if err := os.Remove(appConfig.DbPath); err != nil {
			return err
		}
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		configDir := filepath.Join(homeDir, ".config", "godo")
		err = os.MkdirAll(configDir, os.ModePerm)
		if err != nil {
			return err
		}
		configPath := filepath.Join(configDir, "config.json")

		if err := os.Remove(configPath); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(removeDatabaseCmd)
}
