package main

import (
	"encoding/json"
	"fmt"
	"os"
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
