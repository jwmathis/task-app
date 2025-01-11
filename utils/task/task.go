package task

import "fmt"

// Task represents a single task in the task app
type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func AddTask(existing_tasks []Task, description string) ([]Task, error) {
	// Create a new task
	new_task := Task{
		ID:          len(existing_tasks) + 1,
		Description: description,
		Status:      "todo",
	}

	// Append the new task to the existing tasks
	return append(existing_tasks, new_task), nil
}

func ListTasks(existing_tasks []Task) {
	if len(existing_tasks) == 0 {
		fmt.Println("No tasks found.")
	} else {
		for i, task := range existing_tasks {
			fmt.Printf("%d. %s (%s)\n", i+1, task.Description, task.Status)
		}
	}
}

func MarkTaskAsDone(tasks []Task, taskID int) ([]Task, error) {
	if taskID < 1 || taskID > len(tasks) {
		return nil, fmt.Errorf("invalid task ID: %d", taskID)
	}

	tasks[taskID-1].Status = "done"

	return tasks, nil
}

func DeleteTask(tasks []Task, taskID int) ([]Task, error) {
	if taskID < 1 || taskID > len(tasks) {
		return nil, fmt.Errorf("invalid task ID: %d", taskID)
	}
	var updatedTasks []Task
	for i, task := range tasks {
		if task.ID != taskID {
			task.ID = i + 1
			updatedTasks = append(updatedTasks, task)
		}
	}
	return updatedTasks, nil
}
