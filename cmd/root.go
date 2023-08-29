package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Bilou4/godo/configuration"
	"github.com/Bilou4/godo/model"
	"github.com/glebarez/sqlite"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:          "godo",
	Short:        "A simple Todo list Application.",
	Long:         `Simple Todo list Application to save and manage to-dos by saving them to your local storage.`,
	SilenceUsage: true,
}

var tr model.TaskRepository
var lr model.ListRepository
var appConfig configuration.Config

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}

func initAppConfig() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	configDir := filepath.Join(homeDir, ".config", "godo")
	configPath := filepath.Join(configDir, "config.json")

	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		// configPath does not exist on disk
		return fmt.Errorf("cannot find %s on disk. Did you run 'godo init'?", configPath)
	}
	content, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, &appConfig)
	if err != nil {
		return err
	}
	return nil
}

func openDatabase(path string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{
		FullSaveAssociations: true,
		Logger:               logger.Default.LogMode(logger.Silent),
		// Logger:               logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	err = db.AutoMigrate(&model.Task{})
	if err != nil {
		return nil, fmt.Errorf("AutoMigrate model.Task error: %w", err)
	}

	db.AutoMigrate(&model.List{})
	if err != nil {
		return nil, fmt.Errorf("AutoMigrate model.List error: %w", err)
	}
	return db, nil
}

func persistentPreRun() error {
	err := initAppConfig()
	if err != nil {
		return err
	}

	db, err := openDatabase(appConfig.DbPath)
	if err != nil {
		return err
	}
	tr = model.TaskRepository{DB: db}
	lr = model.ListRepository{DB: db}
	return nil
}
