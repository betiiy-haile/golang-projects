package commands

import (
	"fmt"
	"time"

	"example.com/task-manager/storage"
)

func TaskStats() {
	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	total := len(tasks)
	completed := 0
	pending := 0
	overdue := 0

	now := time.Now()
	for _, task := range tasks {
		if task.Completed {
			completed++
		} else {
			pending++
			if task.DueDate.Before(now) {
				overdue++
			}
		}
	}

	fmt.Println("📊 Task Statistics:")
	fmt.Println("Total:", total)
	fmt.Println("✅ Completed:", completed)
	fmt.Println("🕒 Pending:", pending)
	fmt.Println("⚠️ Overdue:", overdue)
}