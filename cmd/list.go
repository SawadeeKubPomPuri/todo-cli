package cmd

import (
	"todo-cli/todo"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	Run: func(cmd *cobra.Command, args []string) {
		list := &todo.TodoList{}
		list.Load()
		list.PrintList()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
