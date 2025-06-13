package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("No command provided.")
		return
	}

	command := args[1]

	switch command {
	case "add":
		if len(args) < 3 {
			fmt.Println("Please provide a task description.")
			return
		}
		description := args[2]
		addTask(description)
	case "list":
		statusFilter := ""
		if len(args) > 2 {
			statusFilter = args[2]
		}
		listTasks(statusFilter)

	case "mark-in-progress":
		if len(args) < 3 {
			fmt.Println("Please provide task ID.")
			return
		}
		markTask(args[2], "in-progress")

	case "mark-done":
		if len(args) < 3 {
			fmt.Println("Please provide task ID.")
			return
		}
		markTask(args[2], "done")

	case "update":
		if len(args) < 4 {
			fmt.Println("Please provide task ID and new description.")
			return
		}
		// Combine all remining arguments into full description
		newDiscription := ""
		for i := 3; i < len(args); i++ {
			if i > 3 {
				newDiscription += " "
			}
			newDiscription += args[i]
			fmt.Println("New description:--->>>", newDiscription)
		}
		updateTask(args[2], newDiscription)

	case "delete":
		if len(args) < 3 {
			fmt.Println("Please provide task ID to delete.")
			return
		}
		deleteTask(args[2])

	case "help":
		showHelp()
	case "version":
		showVersion()
	case "contact":
		showContact()

	default:
		fmt.Println("Unknown command:", command)
	}
}

const dataFile = "tasks.json"

// LoadTasks
func LoadTasks() ([]Task, error) {
	var tasks []Task

	// If file does not exist, return empty slice
	if _, err := os.Stat(dataFile); os.IsNotExist(err) {
		return tasks, nil
	}

	data, err := os.ReadFile(dataFile)
	if err != nil {
		return nil, err
	}

	// Handle empty file content
	if len(data) == 0 {
		return tasks, nil
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(dataFile, data, 0644)
	return err
}

func addTask(description string) {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	// calculate new ID
	nextID := 1
	if len(tasks) > 0 {
		nextID = tasks[len(tasks)-1].ID + 1
	}

	now := time.Now().Format(time.RFC3339)

	newTask := Task{
		ID:          nextID,
		Description: description,
		Status:      "todo",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	tasks = append(tasks, newTask)

	err = SaveTasks(tasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Printf("Task added successfully (ID: %d)\n", newTask.ID)
}

func listTasks(statusFilter string) {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	filtered := []Task{}
	for _, task := range tasks {
		if statusFilter == "" || task.Status == statusFilter {
			filtered = append(filtered, task)
		}
	}

	if len(filtered) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	fmt.Println("Tasks:")
	for _, task := range tasks {
		if statusFilter == "" || task.Status == statusFilter {
			fmt.Printf("ID: %d | Description: %s | Status: %s | Created At: %s | Updated At: %s\n",
				task.ID, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
		}
	}
}

func markTask(idStr string, status string) {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid task ID:", idStr)
		return
	}

	updated := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now().Format(time.RFC3339)
			updated = true
			break
		}
	}

	if !updated {
		fmt.Printf("Task with ID %d not found.\n", id)
		return
	}

	err = SaveTasks(tasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Printf("Task ID %d marked as %s successfully.\n", id, status)
}

func updateTask(idStr string, newDescribtion string) {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid task ID:", idStr)
		return
	}

	updated := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = newDescribtion
			tasks[i].UpdatedAt = time.Now().Format(time.RFC3339)
			updated = true
			break
		}
	}

	if !updated {
		fmt.Printf("Task with ID %d not found.\n", id)
		return
	}

	err = SaveTasks(tasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Printf("Task ID %d updated seccessfully.\n", id)
}

func deleteTask(idStr string) {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid task ID:", idStr)
		return
	}

	index := -1
	for i, task := range tasks {
		if task.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Printf("Task with ID %d not found./n", id)
		return
	}

	// remove the task
	tasks = append(tasks[:index], tasks[index+1:]...)

	err = SaveTasks(tasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Printf("Task ID %d deleted successfully.\n", id)
}

func showHelp() {
	fmt.Println("Task CLI - Simple Task Manager")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("  add <description>               - Add new task")
	fmt.Println("  list [status]                   - List tasks (optional: todo, in-progress, done)")
	fmt.Println("  update <id> <new description>   - Update task description")
	fmt.Println("  mark-in-progress <id>           - Mark task as in-progress")
	fmt.Println("  mark-done <id>                  - Mark task as done")
	fmt.Println("  delete <id>                     - Delete task")
	fmt.Println("  version                          - Show version info")
	fmt.Println("  contact                          - Show contact info")
}

func showVersion() {
	fmt.Println("Task CLI Version 1.0.0")
	fmt.Println("Built with Go")
}

func showContact() {
	fmt.Println("Created by Reza Barzegar Gashti")
	fmt.Println("GitHub: https://github.com/rezaBG")
	fmt.Println("Email: rezabarzegargashti@example.com")
}
