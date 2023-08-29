package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Bilou4/godo/configuration"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"i"},
	Short:   "Init the Application.",
	Long:    `Initialize the Application by creating a configuration file in '~/.config/godo'.`,
	RunE: func(cmd *cobra.Command, args []string) error {
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

		// get command line flags
		dbPath, err := cmd.Flags().GetString("db-path")
		if err != nil {
			return err
		}

		force, err := cmd.Flags().GetBool("force")
		if err != nil {
			return err
		}

		// if dbPath is the default value
		if strings.HasPrefix(dbPath, "~/") {
			// change '~' by the user home directory
			dbPath = filepath.Join(homeDir, dbPath[2:])
		}

		dbPath, err = filepath.Abs(dbPath)
		if err != nil {
			return err
		}

		// if the database file already exists
		if fi, err := os.Stat(dbPath); err == nil {
			if fi.IsDir() {
				return fmt.Errorf("provided path is a directory: '%s'", dbPath)
			}
			if !force {
				return fmt.Errorf("path to the database already exists: '%s'", dbPath)
			} else {
				// removing previous database
				err = os.Remove(dbPath)
				if err != nil {
					return fmt.Errorf("cannot remove previous database: %w", err)
				}
			}
		}

		config := configuration.Config{DbPath: dbPath}

		content, err := json.Marshal(config)
		if err != nil {
			return err
		}

		var perms int = os.O_CREATE | os.O_WRONLY
		if _, err := os.Stat(configPath); err == nil {
			// configPath already exists
			if force {
				perms = os.O_TRUNC | os.O_CREATE | os.O_WRONLY
			} else {
				return fmt.Errorf("'%s' file already exists", configPath)
			}
		}

		f, err := os.OpenFile(configPath, perms, 0644)
		if err != nil {
			return err
		}

		_, err = f.Write(content)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringP("db-path", "p", filepath.Join("~", ".config", "godo", "godo.db"), "Path where to store the database.")
	initCmd.Flags().Bool("force", false, "Overwrites the configuration and/or the database if they already exist.")
}
