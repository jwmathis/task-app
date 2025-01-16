package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"task-app/utils/fileio"
	"task-app/utils/help"
	"task-app/utils/task"
)

func main() {

	helpFlag := flag.Bool("help", false, "Show help menu")
	flag.Parse()

	if *helpFlag {
		help.PrintHelp()
		return
	}
	// Check if the user provided a command
	if len(os.Args) < 2 {
		help.PrintHelp() // Give instructions for use
		return
	}

	// Get the command
	command := os.Args[1] // Store first argument as command using inference

	filePath := filepath.Join("data", "user_tasks.json") // Set the file path to the JSON file; cross-platform compatible
	file, _ := fileio.CheckAndCreateFile(filePath)       // Check if file exists and create if not and open the file
	defer file.Close()
	existing_tasks, err := fileio.ReadFromJsonFile(filePath) // Read tasks from the JSON file
	help.CheckErr(err)

	// Use the command to execute the appropriate function
	switch command {

	case "add", "a":
		// Handle adding a task
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task description. Usage: task-app add [description]")
			return
		}

		// Get the task description from the second argument
		taskDescription := strings.Join(os.Args[2:], " ") // Join the arguments into a single string

		// Add the task
		existing_tasks, _ = task.AddTask(existing_tasks, taskDescription)

	case "list", "li", "view":
		// Handle listing tasks
		task.ListTasks(existing_tasks)

	case "mark":
		// Handle marking a task as done
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task ID. Usage: task-app done [taskID]")
			return
		}
		command := os.Args[2]
		// Get the task ID
		taskID := os.Args[3]
		// Convert the task ID to an integer
		taskID_int, err := strconv.Atoi(taskID)
		help.CheckErr(err)

		if command == "done" {
			// Mark the task as done
			existing_tasks, err = task.MarkTaskAsDone(existing_tasks, taskID_int)
			if err != nil {
				help.CheckErr(err)
			}
			// Print a confirmation message
			fmt.Printf("Task marked as done.\n")

		} else if command == "todo" {
			// Mark the task as todo
			existing_tasks, err = task.MarkTaskAsTodo(existing_tasks, taskID_int)
			if err != nil {
				help.CheckErr(err)
			}
			// Print a confirmation message
			fmt.Printf("Task marked as todo.\n")
		}

	case "delete", "del":
		// Handle deleting a task
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task ID. Usage: task-app delete [taskID]")
			return
		}
		// Get the task ID
		taskID := os.Args[2]
		if taskID == "all" {
			existing_tasks = []task.Task{} // Delete all tasks
			// Print a confirmation message
			fmt.Printf("All tasks deleted.\n")
		} else {
			// Convert the task ID to an integer
			taskID_int, err := strconv.Atoi(taskID)
			help.CheckErr(err)

			// Delete the task
			existing_tasks, err = task.DeleteTask(existing_tasks, taskID_int)
			help.CheckErr(err)
		}

	default:
		help.PrintUsage()
		return
	}

	// Update the JSON file with the changes
	err = fileio.WriteToJsonFile(filePath, existing_tasks)
	help.CheckErr(err)
}
