package help

import "fmt"

// Function to check for errors
func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}

// Function to print usage instructions
func PrintUsage() {
	fmt.Println("Usage syntax: task [command] [arguments]")
	fmt.Println("Valid commands include: add, list, delete, or done.")
	fmt.Println("Example: task add \"Buy groceries\"")
	fmt.Println("Type task -help for more information.")
}

// Function to print help menu
func PrintHelp() {
	fmt.Println("Task App Help Menu:")
	fmt.Println("Usage: task <command> [options]")
	fmt.Println("Commands:")
	fmt.Println("  add <description>   - Add a new task with description")
	fmt.Println("  list   - List all tasks")
	fmt.Println("  delete <task-id>	- Delete a task by its ID")
	fmt.Println("  done   - Mark a task as done")
	fmt.Println("  -help   - Display this help menu")
}
