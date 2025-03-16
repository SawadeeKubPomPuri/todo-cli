package cmd

import (
	"fmt"
	"strconv"
	"todo-cli/todo"

	"github.com/spf13/cobra"
)

var toggleCmd = &cobra.Command{
	Use:   "tog [index]",
	Short: "Toggle a todo's completion status",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("❌ Invalid index")
			return
		}

		list := &todo.TodoList{}
		list.Load()

		if err := list.ToggleCompletion(index); err != nil {
			fmt.Println("❌", err)
			return
		}

		fmt.Println("✅ Todo toggled at index:", index)
	},
}

func init() {
	rootCmd.AddCommand(toggleCmd)
}
