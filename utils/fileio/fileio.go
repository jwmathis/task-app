package fileio

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"task-app/utils/task"
)

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
func ReadFromJsonFile(filename string) ([]task.Task, error) {
	var existing_tasks []task.Task // Create a slice to hold the tasks

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
func WriteToJsonFile(filename string, tasks []task.Task) error {
	data, err := json.Marshal(tasks) // Marshal the tasks
	if err != nil {
		return fmt.Errorf("error marshalling data: %w", err) // Return error if marshalling fails
	}
	return os.WriteFile(filename, data, 0644) // Write the data to the file
}
