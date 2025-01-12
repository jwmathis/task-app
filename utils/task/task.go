package task

import "fmt"

// Task represents a single task in the task app
type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// AddTask adds a new task to the task list
func AddTask(existing_tasks []Task, description string) ([]Task, error) {

	// Create a new task
	new_task := Task{
		ID:          len(existing_tasks) + 1,
		Description: description,
		Status:      "todo",
	}

	// Append the new task to the existing tasks and return
	return append(existing_tasks, new_task), nil
}

// ListTasks lists all tasks in the task list
func ListTasks(existing_tasks []Task) {
	// Print the list of tasks
	if len(existing_tasks) == 0 {

		fmt.Println("No tasks found.") // Print a message if no tasks are found

	} else {

		fmt.Println("ID    Status	Description")
		fmt.Println("----  ------	-------------")
		for i, task := range existing_tasks { // Iterate over the tasks
			fmt.Printf("%d.    %s   	%s\n", i+1, task.Status, task.Description) // Print the task
		}

	}
}

// MarkTaskAsDone marks a task as done
func MarkTaskAsDone(tasks []Task, taskID int) ([]Task, error) {
	// Check if the task ID is valid
	if taskID < 1 || taskID > len(tasks) {

		return nil, fmt.Errorf("invalid task ID: %d", taskID) // Return an error

	}

	// Mark the task as done
	tasks[taskID-1].Status = "done"

	return tasks, nil // Return the updated tasks
}

// DeleteTask deletes a task
func DeleteTask(tasks []Task, taskID int) ([]Task, error) {
	// Check if the task ID is valid
	if taskID < 1 || taskID > len(tasks) {
		return nil, fmt.Errorf("invalid task ID: %d", taskID) // Return an error
	}

	var updatedTasks []Task      // Create a new slice to hold the updated tasks
	for i, task := range tasks { // Iterate over the tasks
		if task.ID != taskID { // If the task ID is not the one to delete
			task.ID = i + 1                           // Update the task ID
			updatedTasks = append(updatedTasks, task) // Add the task to the updated slice
		}
	}
	return updatedTasks, nil
}
