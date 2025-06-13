# Task Tracker CLI (Golang)

Simple task management application built with Go.  
It allows you to track, update, and manage tasks directly from the command line using a local JSON file as storage.

---

## Features

- Add new tasks
- List all tasks
- List tasks filtered by status (`todo`, `in-progress`, `done`)
- Update task description
- Mark tasks as `done` or `in-progress`
- Delete tasks
- Help, version, and contact information

---

## Usage

```bash
# Add a new task
go run . add "Buy groceries"

# List all tasks
go run . list

# List tasks by status
go run . list todo
go run . list in-progress
go run . list done

# Update task description
go run . update <id> "New description"

# Mark task as done / in-progress
go run . mark-done <id>
go run . mark-in-progress <id>

# Delete a task
go run . delete <id>

# Show help
go run . help

# Show version
go run . version

# Show contact
go run . contact
```
