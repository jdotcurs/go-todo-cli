# go-todo-cli

go-todo-cli is a simple command-line interface (CLI) todo application built with Go and SQLite. It allows users to manage their tasks efficiently from the terminal.

## Purpose

The purpose of creating this CLI tool is to:

1. Provide a lightweight, fast, and easy-to-use task management solution.
2. Demonstrate the use of Go for building CLI applications.
3. Showcase the integration of a SQLite database for persistent storage.
4. Offer a practical example of using the Cobra library for creating powerful CLI applications.

## Installation

To install go-todo-cli, make sure you have Go installed on your system, then run:

```
go install github.com/jdotcurs/go-todo-cli@latest
```

## Usage

The go-todo-cli supports the following commands:

### Create a task

```
todo create -d "Task description" -t "YYYY-MM-DD"
```

- `-d` or `--description`: Task description (required)
- `-t` or `--due`: Due date in YYYY-MM-DD format (required)

### List all tasks

```
todo list
```

### Edit a task

```
todo edit <task_id> [-d "New description"] [-t "YYYY-MM-DD"]
```

- `<task_id>`: The ID of the task to edit (required)
- `-d` or `--description`: New task description (optional)
- `-t` or `--due`: New due date in YYYY-MM-DD format (optional)

### Delete a task

```
todo delete <task_id>
```

- `<task_id>`: The ID of the task to delete (required)

## Examples

1. Create a new task:
   ```
   todo create -d "Buy groceries" -t "2024-03-15"
   ```

2. List all tasks:
   ```
   todo list
   ```

3. Edit a task:
   ```
   todo edit 1 -d "Buy groceries and cleaning supplies" -t "2024-03-16"
   ```

4. Delete a task:
   ```
   todo delete 1
   ```

# Development

To contribute to this project or run it locally, clone the repository and install the dependencies:

```
git clone https://github.com/jdotcurs/go-todo-cli.git
cd go-todo-cli
go mod tidy
```

To build the application, use the following command:

```
go build
```

To run the application, use the following command:

```
./go-todo-cli [command]
```

This will start the CLI application.

