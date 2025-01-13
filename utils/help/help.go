package help

import "fmt"

// Function to check for errors
func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
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
