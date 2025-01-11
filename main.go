package main

import (
	"fmt"
	"os"
	"strconv"
	"task-app/utils/fileio"
	"task-app/utils/task"
)

// Function to check for errors
func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Check if the user provided a command
	// Give instructions for use
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	// Get the command
	command := os.Args[1] // Store firts argument as command using inference
	filePath := "data/user_tasks.json"

	file, _ := fileio.CheckAndCreateFile(filePath) // Check if file exists and create if not and open the file
	defer file.Close()

	existing_tasks, err := fileio.ReadFromJsonFile(filePath)
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
		existing_tasks, err = task.AddTask(existing_tasks, taskDescription)
		checkErr(err)

		// Print a confirmation message
		fmt.Printf("Task added: %s\n", taskDescription)

	case "list", "li", "view":
		// Handle listing tasks
		fmt.Println("Listing tasks:")

		task.ListTasks(existing_tasks)

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
		} else {
			taskID_int, err := strconv.Atoi(taskID)
			checkErr(err)

			existing_tasks, err = task.MarkTaskAsDone(existing_tasks, taskID_int)
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
			existing_tasks = []task.Task{}
		} else {
			taskID_int, err := strconv.Atoi(taskID)
			checkErr(err)
			existing_tasks, err = task.DeleteTask(existing_tasks, taskID_int)
			checkErr(err)

			// Print a confirmation message
			fmt.Printf("Task deleted.\n")
		}

	default:
		printUsage()
		return
	}

	// Write the updated tasks to the file
	err = fileio.WriteToJsonFile(filePath, existing_tasks)
	checkErr(err)
}

func printUsage() {
	fmt.Println("Usage syntax: task-app [command] [arguments]")
	fmt.Println("Valid commands include: add, list, delete, or done.")
	fmt.Println("Example: task-app add \"Buy groceries\"")
}
