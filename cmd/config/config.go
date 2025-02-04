package config

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ConfigCmd is the parent command for config operations.
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage global config settings",
}

// Add command
var addCmd = &cobra.Command{
	Use:   "add <key> <value>",
	Short: "Set a global configuration value",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key, value := args[0], args[1]
		viper.Set(key, value)
		if err := viper.WriteConfig(); err != nil {
			fmt.Println("Error writing config:", err)
			os.Exit(1)
		}
		fmt.Printf("Config added: %s = %s\n", key, value)
	},
}

// Remove command
var removeCmd = &cobra.Command{
	Use:   "remove <key>",
	Short: "Remove a global configuration value",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		viper.Set(key, nil)
		if err := viper.WriteConfig(); err != nil {
			fmt.Println("Error updating config:", err)
			os.Exit(1)
		}
		fmt.Printf("Config removed: %s\n", key)
	},
}

// View command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View all global configuration values",
	Run: func(cmd *cobra.Command, args []string) {
		settings := viper.AllSettings()
		if len(settings) == 0 {
			fmt.Println("No config values set.")
			return
		}
		for key, value := range settings {
			fmt.Printf("%s = %v\n", key, value)
		}
	},
}

func init() {
	// Add the commands under configCmd
	ConfigCmd.AddCommand(addCmd)
	ConfigCmd.AddCommand(removeCmd)
	ConfigCmd.AddCommand(viewCmd)
}
