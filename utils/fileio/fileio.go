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
		if errors.Is(err, os.ErrNotExist) { // if error is ErrNotExist, then file does not exist
			file, createErr := os.Create(filename) // Create the file
			if createErr != nil {                  // If error creating file
				return nil, fmt.Errorf("error creating file: %w", createErr) // Return no file and error
			}
			return file, nil // Return the file and no error if successful
		} else {
			return nil, fmt.Errorf("error checking file: %w", err) // Return no file and error if any other problem
		}
	}

	// If File exists
	file, openErr := os.OpenFile(filename, os.O_RDWR, 0644) // Open the file
	if openErr != nil { // If error opening file
		return nil, fmt.Errorf("file exists but error opening file: %w", openErr) // Return no file and error if opening file fails
	}
	return file, nil // Return the file and no error if successful file opening
}

// Reads and unmarshals JSON data from a file into a slice of tasks
func ReadFromJsonFile(filename string) ([]task.Task, error) {
	var existing_tasks []task.Task // Create a slice to hold the tasks

	data, err := os.ReadFile(filename) // Read the file
	if err == nil {                    // If no error
		if len(data) > 0 { // If the file is not empty
			if err := json.Unmarshal(data, &existing_tasks); err != nil { // Unmarshal the data
				return nil, fmt.Errorf("error unmarshalling data: %w", err) // Return empty Task Struct Array and error if unmarshalling fails
			}
		}
	} else if !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("error finding/reading file: %w", err) // Return error if reading file fails
	}
	return existing_tasks, nil // Return the an array of tasks and no error if successful
}

// Marshals tasks and writes them to the provided file
func WriteToJsonFile(filename string, tasks []task.Task) error {
	data, err := json.Marshal(tasks) // Marshal the tasks
	if err != nil {
		return fmt.Errorf("error marshalling data: %w", err) // Return error if marshalling fails
	}
	return os.WriteFile(filename, data, 0644) // Write the marshalled data to the file and return error code
}
