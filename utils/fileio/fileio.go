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
	_, err := os.Stat(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// File does not exist, create it
			file, createErr := os.Create(filename)
			if createErr != nil {
				return nil, fmt.Errorf("error creating file: %w", createErr)
			}
			// Debugging messages
			//fmt.Printf("File '%s' created successfully\n", filename)
			return file, nil
		} else {
			return nil, fmt.Errorf("error checking file: %w", err)
		}
	}

	// File already exists
	file, openErr := os.OpenFile(filename, os.O_RDWR, 0644)
	if openErr != nil {
		return nil, fmt.Errorf("file exists but error opening file: %w", openErr)
	}
	// Debugging messages
	//fmt.Printf("File '%s' already exists\n", filename)
	return file, nil
}

// Reads and unmarshals JSON data from a file into a slice of tasks
func ReadFromJsonFile(filename string) ([]task.Task, error) {
	var existing_tasks []task.Task

	data, err := os.ReadFile(filename)
	if err == nil {
		if len(data) > 0 {
			if err := json.Unmarshal(data, &existing_tasks); err != nil {
				return nil, fmt.Errorf("error unmarshalling data: %w", err)
			}
		}
	} else if !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	return existing_tasks, nil
}

// Marshals tasks and writes them to the provided file
func WriteToJsonFile(filename string, tasks []task.Task) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("error marshalling data: %w", err)
	}
	return os.WriteFile(filename, data, 0644)
}
