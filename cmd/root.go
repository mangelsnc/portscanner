package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "portscanner",
	Short: "PortScanner is, as it name suggest, a simple port scanner tool",
	Long: `PortScanner is, as it name suggest, a simple port scanner tool.

This tool comes with no warranties, beacuse was made just for fun and learning.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

// Execute runs the command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
