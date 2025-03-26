package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type TodoList []Todo

const todoFile = "todos.json"

func newTodo(title string) Todo {
	return Todo{
		Title:       title,
		CreatedAt:   time.Now(),
		Completed:   false,
		CompletedAt: nil,
	}
}

func (t *TodoList) Load() error {
	data, err := os.ReadFile(todoFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			*t = TodoList{}
			return nil
		}
		return err
	}
	return json.Unmarshal(data, t)
}

func (t *TodoList) Save() error {
	data, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(todoFile, data, 0644)
}

func (t *TodoList) validateIndex(index int) error {
	if index < 0 || index >= len(*t) {
		return errors.New("invalid index")
	}
	return nil
}

func (t *TodoList) Add(title string) {
	*t = append(*t, newTodo(title))
	t.Save()
	fmt.Println("✅ Added:", title)
}

func (t *TodoList) Delete(index int) error {
	if err := t.validateIndex(index); err != nil {
		return err
	}
	*t = append((*t)[:index], (*t)[index+1:]...)
	t.Save()
	fmt.Println("❌ Deleted Todo at index", index)
	return nil
}

func (t *TodoList) ToggleCompletion(index int) error {
	if err := t.validateIndex(index); err != nil {
		return err
	}
	todo := &(*t)[index]
	if todo.Completed {
		todo.CompletedAt = nil
	} else {
		completionTime := time.Now()
		todo.CompletedAt = &completionTime
	}
	todo.Completed = !todo.Completed
	t.Save()
	fmt.Println("🔄 Toggled Todo at index", index)
	return nil
}

func formatCompletedAt(t *Todo) string {
	if t.CompletedAt == nil {
		return ""
	}
	return t.CompletedAt.Format(time.RFC3339)
}

func (t *TodoList) PrintList() {
	tbl := table.New(os.Stdout)
	tbl.SetHeaders("#", "Title", "Completion", "Created At", "Completed At")

	for index, todo := range *t {
		status := "❌"
		if todo.Completed {
			status = "✅"
		}

		tbl.AddRow(
			strconv.Itoa(index),
			todo.Title,
			status,
			todo.CreatedAt.Format(time.RFC3339),
			formatCompletedAt(&todo),
		)
	}

	tbl.Render()
}

func (t *TodoList) Edit(index int, title string) error {
	if err := t.validateIndex(index); err != nil {
		println("Invalid index")
		return err
	}
	(*t)[index].Title = title
	t.Save()
	fmt.Println("✏️ Edited Todo at index", index)
	return nil
}
