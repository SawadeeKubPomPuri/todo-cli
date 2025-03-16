package cmd

import (
	"fmt"
	"todo-cli/todo"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [title]",
	Short: "Add a new todo",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		list := &todo.TodoList{}
		list.Load()
		list.Add(args[0])
		fmt.Println("✅ Todo added:", args[0])
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
