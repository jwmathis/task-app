package task

import (
	"fmt"
	"task-app/utils/help"
)

// Define Task struct to represent a single task in the task app
type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// Define colors
const (
	Reset = "\033[0m" // Set color back to default color scheme of user terminal
	Red   = "\033[31m" // Set color scheme to red
	Green = "\033[32m" // Set color scheme to green
)

// AddTask adds a new task to the task list
func AddTask(existing_tasks []Task, description string) ([]Task) {
	help.ClearScreen()

	// Create a new task
	new_task := Task{
		ID:          len(existing_tasks) + 1, // ID is based on length of Task array
		Description: description,
		Status:      "todo",
	}

	// Print a confirmation message
	fmt.Printf("Task added: %s\n", description)

	// Append the new task to the existing tasks array and return
	return append(existing_tasks, new_task)
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

			if task.Status == "done" { // If the task status is done
				color = Green // Set the color to green
			} else if task.Status == "todo" { // If the task is todo
				color = Red // Set the color to red
			}

			fmt.Printf("%s%d.    %s%s   	%s%s\n", color, i+1, Reset, task.Status, Reset, task.Description) // Print the formatted task
		}

	}

}

// MarkTaskAsDone changes a task status to "Done"
func MarkTaskAsDone(tasks []Task, taskID int) ([]Task, error) {
	// Check if the task ID is valid
	if taskID < 1 || taskID > len(tasks) {

		return tasks, fmt.Errorf("invalid task ID: %d", taskID) // Return provided tasks and an error

	}

	// Chagnge the task status to done
	tasks[taskID-1].Status = "Done"

	return tasks, nil // Return the updated tasks and no error
}


// MarkTaskAsTodo changes a task status to "Pending"
func MarkTaskAsTodo(tasks []Task, taskID int) ([]Task, error) {
	// Check if the task ID is valid
	if taskID < 1 || taskID > len(tasks) {
		return tasks, fmt.Errorf("invalid task ID: %d", taskID) // Return provided tasks and an error
	}

	// Mark the task as todo
	tasks[taskID-1].Status = "Pending"

	return tasks, nil // Return updated tasks and no error
}

// DeleteTask removes a task
func DeleteTask(tasks []Task, taskID int) ([]Task, error) {
	help.ClearScreen()

	// Check if the task ID is valid
	if taskID < 1 || taskID > len(tasks) {
		return tasks, fmt.Errorf("invalid task ID: %d", taskID) // Return provided tasks and an error
	}

	var updatedTasks []Task // Create a new slice to hold the updated tasks
	var new_updated_tasks []Task
	var j int = 0
	for _, task := range tasks { // Iterate over the tasks
		if task.ID != taskID { // If the task ID is not the one to delete
			updatedTasks = append(updatedTasks, task) // Append the task to the updated slice
			updatedTasks[j].ID = j + 1	// Change the task ID
			j++ // Increment j
		}
	}

	// Print a confirmation message
	fmt.Printf("Task deleted.\n")

	// return new_updated_tasks, nil
	return updatedTasks, nil // Return updatedTasks array and no error
}
