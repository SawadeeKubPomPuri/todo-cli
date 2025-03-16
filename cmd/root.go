package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo-cli",
	Short: "A simple CLI-based todo list manager",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("❌ Error:", err)
		os.Exit(1)
	}
}
