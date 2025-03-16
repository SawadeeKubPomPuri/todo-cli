package cmd

import (
	"fmt"
	"strconv"
	"todo-cli/todo"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "del [index]",
	Short: "Delete a todo by index",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("❌ Invalid index")
			return
		}

		list := &todo.TodoList{}
		list.Load()

		if err := list.Delete(index); err != nil {
			fmt.Println("❌", err)
			return
		}

		fmt.Println("✅ Todo deleted at index:", index)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
