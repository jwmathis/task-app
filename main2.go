package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Function to check for errors
func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

// Checks if the file exists and creates it if necessary
func CheckAndCreateFile(filename string) (*os.File, error) {
	// Check if file exists
	_, err := os.Stat(filename) // Check if file exists
	if err != nil {             // If file does not exist
		if errors.Is(err, os.ErrNotExist) { // Check if error is ErrNotExist
			file, createErr := os.Create(filename) // Create the file
			if createErr != nil {                  // If error creating file
				return nil, fmt.Errorf("error creating file: %w", createErr) // Return error
			}
			// Debugging messages
			//fmt.Printf("File '%s' created successfully\n", filename)
			return file, nil // Return the file if successful
		} else {
			return nil, fmt.Errorf("error checking file: %w", err) // Return error
		}
	}

	// If File already exists
	file, openErr := os.OpenFile(filename, os.O_RDWR, 0644) // Open the file
	if openErr != nil {                                     // If error opening file
		return nil, fmt.Errorf("file exists but error opening file: %w", openErr) // Return error if opening file fails
	}
	// Debugging messages
	//fmt.Printf("File '%s' already exists\n", filename)
	return file, nil // Return the file
}

// Reads and unmarshals JSON data from a file into a slice of tasks
func ReadFromJsonFile(filename string) ([]Task, error) {
	var existing_tasks []Task // Create a slice to hold the tasks

	data, err := os.ReadFile(filename) // Read the file
	if err == nil {                    // If no error
		if len(data) > 0 { // If the file is not empty
			if err := json.Unmarshal(data, &existing_tasks); err != nil { // Unmarshal the data
				return nil, fmt.Errorf("error unmarshalling data: %w", err) // Return error if unmarshalling fails
			}
		}
	} else if !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("error reading file: %w", err) // Return error if reading file fails
	}
	return existing_tasks, nil // Return the tasks if successful
}

// Marshals tasks and writes them to the provided file
func WriteToJsonFile(filename string, tasks []Task) error {
	data, err := json.Marshal(tasks) // Marshal the tasks
	if err != nil {
		return fmt.Errorf("error marshalling data: %w", err) // Return error if marshalling fails
	}
	return os.WriteFile(filename, data, 0644) // Write the data to the file
}

// Function to print usage instructions
func PrintUsage() {
	ClearScreen()
	fmt.Println("Usage syntax: task [command] [arguments]")
	fmt.Println("Valid commands include: add, list, delete, or done.")
	fmt.Println("Example: task add \"Buy groceries\"")
	fmt.Println("Type task -help for more information.")
}

// Function to print help menu
func PrintHelp() {
	ClearScreen()
	fmt.Println("\nTask App Help Menu:")
	fmt.Println("\nUsage: ./task <command> [options]\n")
	fmt.Println("Commands:             	Example: 			Meaning:")
	fmt.Println("----------------	--------------------		-----------------------")
	fmt.Println("add, a    		./task add <description>	- Add a new task with description")
	fmt.Println("view, list, li    	./task view			- List all tasks")
	fmt.Println("delete, del <task-id>	./task delete <task-id>		- Delete a task by its ID")
	fmt.Println("done, complete    	./task done <task-id>		- Mark a task as done")
	fmt.Println("-help   		./task -help			- Display this help menu\n")
}

// Task represents a single task in the task app
type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

//40.98

// Define colors
const (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Green = "\033[32m"
)

// AddTask adds a new task to the task list
func AddTask(existing_tasks []Task, description string) ([]Task, error) {
	ClearScreen()

	// Create a new task
	new_task := Task{
		ID:          len(existing_tasks) + 1,
		Description: description,
		Status:      "todo",
	}

	// Print a confirmation message
	fmt.Printf("Task added: %s\n", description)

	// Append the new task to the existing tasks and return
	return append(existing_tasks, new_task), nil
}

// ListTasks lists all tasks in the task list
func ListTasks(existing_tasks []Task) {
	ClearScreen()
	// Print the list of tasks
	if len(existing_tasks) == 0 {

		fmt.Println("No tasks found. Add a task using the 'add' command.") // Print a message if no tasks are found
		fmt.Println("Type task -help for more information.")
	} else {

		color := Reset // Default color
		fmt.Println("ID    Status	Description")
		fmt.Println("----  ------	-------------")
		for i, task := range existing_tasks { // Iterate over the tasks

			if task.Status == "done" { // If the task is done
				color = Green // Set the color to green
			} else if task.Status == "todo" { // If the task is todo
				color = Red // Set the color to red
			}

			fmt.Printf("%s%d.    %s%s   	%s%s\n", color, i+1, Reset, task.Status, Reset, task.Description) // Print the task
		}

	}

}

// MarkTaskAsDone marks a task as done
func MarkTaskAsDone(tasks []Task, taskID int) ([]Task, error) {
	// Check if the task ID is valid
	if taskID < 1 || taskID > len(tasks) {

		return tasks, fmt.Errorf("invalid task ID: %d", taskID) // Return an error

	}

	// Mark the task as done
	tasks[taskID-1].Status = "done"

	return tasks, nil // Return the updated tasks
}

func MarkTaskAsTodo(tasks []Task, taskID int) ([]Task, error) {
	// Check if the task ID is valid
	if taskID < 1 || taskID > len(tasks) {
		return tasks, fmt.Errorf("invalid task ID: %d", taskID) // Return an error
	}

	// Mark the task as todo
	tasks[taskID-1].Status = "todo"

	return tasks, nil
}

// DeleteTask deletes a task
func DeleteTask(tasks []Task, taskID int) ([]Task, error) {
	ClearScreen()

	// Check if the task ID is valid
	if taskID < 1 || taskID > len(tasks) {
		return tasks, fmt.Errorf("invalid task ID: %d", taskID) // Return original tasks and an error
	}

	var updatedTasks []Task // Create a new slice to hold the updated tasks
	var new_updated_tasks []Task
	for _, task := range tasks { // Iterate over the tasks
		if task.ID != taskID { // If the task ID is not the one to delete
			updatedTasks = append(updatedTasks, task) // Add the task to the updated slice
		}
	}

	for i, task := range updatedTasks { // Iterate over the updated tasks
		task.ID = i + 1
		new_updated_tasks = append(new_updated_tasks, task) // Add the task to the new updated slice
	}

	// Print a confirmation message
	fmt.Printf("Task deleted.\n")

	return new_updated_tasks, nil
}

func main() {

	helpFlag := flag.Bool("help", false, "Show help menu")
	flag.Parse()

	if *helpFlag {
		PrintHelp()
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter command:")
	command, _ := reader.ReadString('\n')
	command = strings.TrimSpace(command) // Remove leading and trailing whitespace

	filePath := filepath.Join("data", "user_tasks.json") // Set the file path to the JSON file; cross-platform compatible
	file, _ := CheckAndCreateFile(filePath)              // Check if file exists and create if not and open the file
	defer file.Close()
	existing_tasks, err := ReadFromJsonFile(filePath) // Read tasks from the JSON file
	CheckErr(err)

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
		existing_tasks, _ = AddTask(existing_tasks, taskDescription)

	case "list", "li", "view":
		// Handle listing tasks
		ListTasks(existing_tasks)

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
		CheckErr(err)

		if command == "done" {
			// Mark the task as done
			existing_tasks, err = MarkTaskAsDone(existing_tasks, taskID_int)
			if err != nil {
				CheckErr(err)
			}
			// Print a confirmation message
			fmt.Printf("Task marked as done.\n")

		} else if command == "todo" {
			// Mark the task as todo
			existing_tasks, err = MarkTaskAsTodo(existing_tasks, taskID_int)
			if err != nil {
				CheckErr(err)
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
			existing_tasks = []Task{} // Delete all tasks
			// Print a confirmation message
			fmt.Printf("All tasks deleted.\n")
		} else {
			// Convert the task ID to an integer
			taskID_int, err := strconv.Atoi(taskID)
			CheckErr(err)

			// Delete the task
			existing_tasks, err = DeleteTask(existing_tasks, taskID_int)
			CheckErr(err)
		}

	default:
		PrintUsage()
		return
	}

	// Update the JSON file with the changes
	err = WriteToJsonFile(filePath, existing_tasks)
	CheckErr(err)
}
