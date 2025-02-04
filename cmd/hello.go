package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Prints Hello World",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		fmt.Printf("Hello, %s!\n", name)
	},
}

func init() {
	RootCmd.AddCommand(helloCmd)
	helloCmd.Flags().StringP("name", "n", "World", "Name to greet")
}
