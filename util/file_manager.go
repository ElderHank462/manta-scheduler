package util

import (
	"fmt"
	"os"

	"encoding/json"

	"manta/models"
)

func LoadTasks() []models.Task{
	tasksPath := GetTasksPath()

	// read tasks file
	tasksContent, err := os.ReadFile(tasksPath)
	if err != nil {
		fmt.Println("No tasks found: ", err)
		return nil
	}

	// parse tasks
	var tasks []models.Task
	if err := json.Unmarshal(tasksContent, &tasks); err != nil {
		fmt.Println("Error reading tasks: ", err)
		return nil
	}

	return tasks
}