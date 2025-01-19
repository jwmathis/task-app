package help

import (
	"fmt"
	"os"
)

// Function to check for errors
func CheckErr(e error) {
	if e != nil {
		fmt.Println("Error: ", e) // Print error
		os.Exit(1)                // Exit program with error
	}
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

// Function to print usage instructions
func PrintUsage() {
	ClearScreen()
	fmt.Println("Usage syntax: task [command] [arguments]")
	fmt.Println("Valid commands include: add, list, delete, or done. Type --help for more options.")
	fmt.Println("Example: task add Buy groceries")
	fmt.Println("Type task --help for more information.")
}

// Function to print help menu
func PrintHelp() {
	ClearScreen()
	fmt.Println("\nTask App Help Menu:")
	fmt.Println("\nUsage: task <command> [options]\n")
	fmt.Println("Commands:             	  Example: 			Meaning:")
	fmt.Println("----------------	  --------------------		-----------------------")
	fmt.Println("add, a    		  task add <description>	- Add a new task with description")
	fmt.Println("view, list, li    	  task view			- List all tasks")
	fmt.Println("delete, del 		  task delete <task-id>		- Delete a task by its ID")
	fmt.Println("mark todo, mark pending   task mark todo <task-id>	- Mark a task as todo")
	fmt.Println("mark done, mark complete  task mark done <task-id>	- Mark a task as done")
	fmt.Println("--help   		  task --help			- Display this help menu\n")
}
