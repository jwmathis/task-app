package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if the user provided a command
	if len(os.Args) < 2 {
		fmt.Println("Try using the syntax: task-app [command] [arguments]")
		return
	}

	// Get the command
	command := os.Args[1]

	switch command {
	case "add", "a":
		// Handle adding a task

	case "list", "l":
		// Handle listing tasks

	case "delete", "del":
		// Handle deleting a task

	case "done", "d":
		// Handle marking a task as done
	default:
		fmt.Println("Invalid command. Valid commands include: add, list, delete, or done.")
	}

}
