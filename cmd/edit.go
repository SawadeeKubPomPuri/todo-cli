package cmd

import (
	"fmt"
	"strconv"
	"todo-cli/todo"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [index] [new title]",
	Short: "Edit an existing todo item",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid index. Please provide a valid number.")
			return
		}

		newTitle := args[1]

		todoList := &todo.TodoList{}
		err = todoList.Load()
		if err != nil {
			fmt.Println("Failed to load todos:", err)
			return
		}

		err = todoList.Edit(index, newTitle)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		err = todoList.Save()
		if err != nil {
			fmt.Println("Failed to save todos:", err)
			return
		}

		fmt.Println("Todo updated successfully.")
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
