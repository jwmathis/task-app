package help

import "fmt"

func PrintHelp() {
	fmt.Println("Task App Help Menu:")
	fmt.Println("Usage: task-app <command> [options]")
	fmt.Println("Commands:")
	fmt.Println("  add <description>   - Add a new task with description")
	fmt.Println("  list   - List all tasks")
	fmt.Println("  delete <task-id>	- Delete a task by its ID")
	fmt.Println("  done   - Mark a task as done")
	fmt.Println("  -help   - Display this help menu")
}
