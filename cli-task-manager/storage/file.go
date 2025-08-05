package storage

import (
	"encoding/json"
	"errors"
	"os"

	"example.com/task-manager/models"
)

//  this is what are we expected tos do on this file
// 	1. loadtasks from Tasks.Json
//  2. Save tasks to tasks.json


const FileName = "tasks.json"

func LoadTasks() ([]models.Task, error) {
	data, err := os.ReadFile(FileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []models.Task{}, nil
		}
		return nil, err
	}

	var tasks []models.Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func SaveTasks (tasks []models.Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(FileName, data, 0644)
} 