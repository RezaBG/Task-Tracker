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
		listTasks()

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

func listTasks() {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	fmt.Println("Tasks:")
	for _, task := range tasks {
		fmt.Printf("ID: %d | Description: %s | Status: %s | Created At: %s | Updated At: %s\n",
			task.ID, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
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
