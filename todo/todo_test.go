package todo

import (
	"os"
	"testing"
)

func setupTest(t *testing.T) {
	t.Helper()
	os.Remove(todoFile)
}

func teardownTest(t *testing.T) {
	t.Helper()
	os.Remove(todoFile)
}

func TestTodoList(t *testing.T) {
	t.Run("Add Todo", func(t *testing.T) {
		setupTest(t)
		defer teardownTest(t)

		todoList := &TodoList{}
		todoList.Add("Test Todo")

		if got := len(*todoList); got != 1 {
			t.Errorf("expected 1 todo, got %d", got)
		}

		if (*todoList)[0].Title != "Test Todo" {
			t.Errorf("expected 'Test Todo', got '%s'", (*todoList)[0].Title)
		}
	})

	t.Run("Delete Todo", func(t *testing.T) {
		setupTest(t)
		defer teardownTest(t)

		todoList := &TodoList{}
		todoList.Add("Test Todo")

		err := todoList.Delete(0)
		if err != nil {
			t.Errorf("unexpected error deleting valid index: %v", err)
		}

		if got := len(*todoList); got != 0 {
			t.Errorf("expected 0 todos, got %d", got)
		}

		err = todoList.Delete(0)
		if err == nil {
			t.Errorf("expected error when deleting from empty list, got nil")
		}
	})

	t.Run("Toggle Completion", func(t *testing.T) {
		setupTest(t)
		defer teardownTest(t)

		todoList := &TodoList{}
		todoList.Add("Test Todo")

		err := todoList.ToggleCompletion(0)
		if err != nil {
			t.Errorf("unexpected error toggling completion: %v", err)
		}

		todo := (*todoList)[0]
		if !todo.Completed {
			t.Errorf("expected todo to be marked as completed")
		}
		if todo.CompletedAt == nil {
			t.Errorf("expected CompletedAt to be set")
		}

		err = todoList.ToggleCompletion(0)
		if err != nil {
			t.Errorf("unexpected error toggling back: %v", err)
		}

		todo = (*todoList)[0]
		if todo.Completed {
			t.Errorf("expected todo to be incomplete after second toggle")
		}
		if todo.CompletedAt != nil {
			t.Errorf("expected CompletedAt to be nil after second toggle")
		}
	})

	t.Run("Save and Load", func(t *testing.T) {
		setupTest(t)
		defer teardownTest(t)

		todoList := &TodoList{}
		todoList.Add("Test Todo")

		err := todoList.Save()
		if err != nil {
			t.Fatalf("failed to save todos: %v", err)
		}

		loadedTodoList := &TodoList{}
		err = loadedTodoList.Load()
		if err != nil {
			t.Fatalf("failed to load todos: %v", err)
		}

		if got := len(*loadedTodoList); got != 1 {
			t.Errorf("expected 1 todo, got %d", got)
		}

		if (*loadedTodoList)[0].Title != "Test Todo" {
			t.Errorf("expected 'Test Todo', got '%s'", (*loadedTodoList)[0].Title)
		}
	})
}
