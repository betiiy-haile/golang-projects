package commands

import (
	"fmt"
	"strings"
	"time"

	"example.com/task-manager/models"
	"example.com/task-manager/storage"
)

func AddTask(args []string) {
	if len(args) < 2{
		fmt.Println("Usage: add --title=\"Task title\" [--description=\"desc\"] [--due=\"YYYY-MM-DD\"] [--priority=low|medium|high]")
		return
	}

	task := models.Task{
		Title: "",
		Priority: "low",
		Completed: false,
	}

	for _, arg := range args {
		if strings.HasPrefix(arg, "--title=") {
			task.Title =  strings.TrimPrefix(arg, "--title=")
		} else if strings.HasPrefix(arg, "--description=") {
			task.Description = strings.TrimPrefix(arg, "--description=")
		} else if strings.HasPrefix(arg, "--due=") {
			dueStr := strings.TrimPrefix(arg, "--due=")
			parsedDue, err := time.Parse("2006-01-02", dueStr)
			if err != nil {
				fmt.Println("Error: Invalid due date format. Use YYYY-MM-DD.")
				return
			}
			task.DueDate = parsedDue
		} else if strings.HasPrefix(arg, "--priority=") {
			task.Priority = strings.TrimPrefix(arg, "--priority=")
		}
	}

	if task.Title == "" {
		fmt.Println("Err: Title is required.")
		return
	}


	tasks, _ := storage.LoadTasks()
	task.ID = getNextID(tasks)
	tasks = append(tasks, task)

	if err := storage.SaveTasks(tasks); err != nil {
		fmt.Println("Error saving task:", err)
		return
	}

	fmt.Println("✅ Task added successfully.")
}

func getNextID(tasks []models.Task) int {
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}	