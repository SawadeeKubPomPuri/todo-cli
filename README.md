# Todo CLI

A simple command-line todo list manager built with Go and Cobra.

## ✨ Features

- ➕ **Add tasks** - Quickly capture new todo items
- ❌ **Delete tasks** - Remove completed or unwanted todos by index
- 📋 **List tasks** - View all your todos in a clean, formatted table
- ✅ **Toggle completion** - Mark todos as complete or incomplete
- 💾 **Persistent storage** - Your todos are saved locally in JSON format

## 📋 Dependencies

- [Cobra](https://github.com/spf13/cobra) v1.9.1 - Powerful command-line interface library
- [Aquasecurity Table](https://github.com/aquasecurity/table) v1.8.0 - Beautiful table formatting

## 🚀 Installation

### Option 1: From Source

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/todo-cli.git
   cd todo-cli
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Build the project:
   ```bash
   go build -o todo-cli
   ```

## 📖 Usage

Run the CLI with:

```bash
todo-cli <command> [arguments]
```

### Commands

| Command | Description | Example |
|---------|-------------|---------|
| `add <title>` | Add a new todo | `todo-cli add "Homework 1"` |
| `del <index>` | Delete a todo by index | `todo-cli del 2` |
| `list` | List all todos | `todo-cli list` |
| `tog <index>` | Toggle a todo's completion status | `todo-cli tog 1` |

## 📁 Data Storage

Your todos are stored in `~/.todo-cli/todos.json` and automatically loaded when you run any command.

## 🛠️ Development

### Project Structure
```
todo-cli/
├── cmd/
│   └── root.go
├── internal/
│   └── todo/
│       ├── model.go
│       └── storage.go
├── main.go
├── go.mod
└── go.sum
```

### Running Tests
```bash
go test ./...
```
