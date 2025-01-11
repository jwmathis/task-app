package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Defining a task
type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// Function to check for errors
func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func checkAndCreateFile(filename string) (*os.File, error) {
	// Check if file exists
	_, err := os.Stat(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// File does not exist, create it
			file, createErr := os.Create(filename)
			if createErr != nil {
				return nil, fmt.Errorf("error creating file: %w", createErr)
			}
			// Debugging messages
			//fmt.Printf("File '%s' created successfully\n", filename)
			return file, nil
		} else {
			return nil, fmt.Errorf("error checking file: %w", err)
		}
	}

	// File already exists
	file, openErr := os.OpenFile(filename, os.O_RDWR, 0644)
	if openErr != nil {
		return nil, fmt.Errorf("file exists but error opening file: %w", openErr)
	}
	// Debugging messages
	//fmt.Printf("File '%s' already exists\n", filename)
	return file, nil
}

func readFromJsonFile(filename string) ([]Task, error) {
	var existing_tasks []Task

	data, err := os.ReadFile(filename)
	if err == nil {
		if len(data) > 0 {
			if err := json.Unmarshal(data, &existing_tasks); err != nil {
				return nil, fmt.Errorf("error unmarshalling data: %w", err)
			}
		}
	} else if !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	return existing_tasks, nil
}

func writeToJsonFile(filename string, tasks []Task) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("error marshalling data: %w", err)
	}
	return os.WriteFile(filename, data, 0644)
}

func main() {

	// Check if the user provided a command
	// Give instructions for use
	if len(os.Args) < 2 {
		fmt.Println("Usage syntax: task-app [command] [arguments]")
		fmt.Println("Valid commands include: add, list, delete, or done.")
		fmt.Println("Example: task-app add \"Buy groceries\"")
		return
	}

	// Get the command
	command := os.Args[1] // Store firts argument as command using inference
	filePath := "temp.json"

	file, _ := checkAndCreateFile(filePath) // Check if file exists and create if not and open the file
	defer file.Close()

	existing_tasks, err := readFromJsonFile(filePath)
	checkErr(err)

	// Use the command to execute the appropriate function
	switch command {

	case "add", "a":
		// Handle adding a task
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task description. Usage: task-app add [description]")
			return
		}

		// Get the task description
		taskDescription := os.Args[2]

		// Create a new task
		new_task := Task{
			ID:          len(existing_tasks) + 1,
			Description: taskDescription,
			Status:      "todo",
		}

		// Append the new task to the existing tasks
		existing_tasks = append(existing_tasks, new_task)

		// Write the updated tasks to the file
		err := writeToJsonFile(filePath, existing_tasks)
		checkErr(err)

		// Print a confirmation message
		fmt.Printf("Task added: %s\n", taskDescription)

	case "list", "li":
		// Handle listing tasks
		fmt.Println("Listing tasks:")

		// Loop through the existing tasks and print them
		if len(existing_tasks) == 0 {
			fmt.Println("No tasks found.")
		} else {
			for i, task := range existing_tasks {
				fmt.Printf("%d. %s (%s)\n", i+1, task.Description, task.Status)
			}
		}

	case "done", "complete", "c", "mark":
		// Handle marking a task as done
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task ID. Usage: task-app done [taskID]")
			return
		}

		taskID := os.Args[2]
		if taskID == "all" {
			for i := range existing_tasks {
				existing_tasks[i].Status = "done"
			}
			err := writeToJsonFile(filePath, existing_tasks)
			checkErr(err)
		} else {
			taskID_int, err := strconv.Atoi(taskID)
			checkErr(err)
			if taskID_int > len(existing_tasks) || taskID_int < 1 {
				fmt.Println("Invalid task ID.")
				return
			}
			existing_tasks[taskID_int-1].Status = "done"
			err = writeToJsonFile(filePath, existing_tasks)
			checkErr(err)
		}

	case "delete", "del":
		// Handle deleting a task
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task ID. Usage: task-app delete [taskID]")
			return
		}

		taskID := os.Args[2]
		if taskID == "all" {
			existing_tasks = []Task{}
			err := writeToJsonFile(filePath, existing_tasks)
			checkErr(err)
		} else {
			taskID_int, err := strconv.Atoi(taskID)
			checkErr(err)
			updateTasks := []Task{}
			var j int = 0
			for _, task := range existing_tasks {
				if task.ID != taskID_int {
					task.ID = j
					updateTasks = append(updateTasks, task)
					j++
				}
			}

			err = writeToJsonFile(filePath, updateTasks)
			checkErr(err)

			// Print a confirmation message
			fmt.Printf("Task deleted.\n")
		}

	default:
		fmt.Println("Invalid command. Valid commands include: add, list, delete, or done.")
		return
	}

}
