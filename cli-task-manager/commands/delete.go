package commands

import (
	"fmt"
	"strconv"

	"example.com/task-manager/models"
	"example.com/task-manager/storage"
)

func DeleteTask(idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid task ID")
		return
	}

	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	newTasks := []models.Task{}
	found := false
	for _, task := range tasks {
		if task.ID != id {
			newTasks = append(newTasks, task)
		} else {
			found = true
		}
	}

	if !found {
		fmt.Println("Task not found")
		return
	}

	err = storage.SaveTasks(newTasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	fmt.Println("Task deleted.")
}