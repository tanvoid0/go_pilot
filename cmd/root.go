package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go_pilot/cmd/config" // Import config package to register the command
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "pilot", // This will be your main command
	Short: "Pilot CLI tool",
	Run: func(cmd *cobra.Command, args []string) {
		// Default action when no subcommands are given
		fmt.Println("Welcome to Pilot CLI!")
	},
}

// CheckAndCreateConfig checks if the Viper config file exists,
// and if not, creates the necessary directories and the config file.
func CheckAndCreateConfig() error {
	// Set Viper config name and type
	configDir := ".pilot"
	configFile := "config"
	configFileType := "json"
	configFilePath := filepath.Join(configDir, configFile+"."+configFileType)

	viper.SetConfigName(configFile)
	viper.AddConfigPath(configDir)
	viper.SetConfigType(configFileType)

	// Attempt to read the config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found, create the directory and config file
			fmt.Println("Config file not found, creating...")

			// Ensure the directory exists
			if err := os.MkdirAll(configDir, os.ModePerm); err != nil {
				return fmt.Errorf("error creating config directory: %w", err)
			}

			// Create an empty config file by writing an empty structure (optional)
			if err := viper.WriteConfigAs(configFilePath); err != nil {
				return fmt.Errorf("error creating config file: %w", err)
			}

			fmt.Println("Config file created successfully.")
		} else {
			return fmt.Errorf("error reading config file: %w", err)
		}
	}
	return nil
}

// Initialize config file and add config command to the root
func init() {
	err := CheckAndCreateConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	// Add the config command to the root
	RootCmd.AddCommand(config.ConfigCmd)
}

// Execute runs the root command
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
