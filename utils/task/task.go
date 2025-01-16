package task

import (
	"fmt"
	"task-app/utils/help"
)

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
	help.ClearScreen()

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
	help.ClearScreen()
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
	help.ClearScreen()

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
