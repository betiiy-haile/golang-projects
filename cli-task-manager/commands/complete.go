package commands

import (
	"fmt"
	"strconv"

	"example.com/task-manager/storage"
)


func CompleteTask(idStr string) {
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

	found := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Completed = true
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Task not found")
		return
	}

	err = storage.SaveTasks(tasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	fmt.Println("Task marked as completed.")
}