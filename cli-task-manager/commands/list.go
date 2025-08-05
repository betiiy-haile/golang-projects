package commands

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"example.com/task-manager/models"
	"example.com/task-manager/storage"
)

func ListTasks(args []string) {

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	fmt.Println("LIST COMMANDS", listCmd)
	today := listCmd.Bool("today", false, "List only today's tasks")
	completed := listCmd.Bool("completed", false, "List completed tasks")
	priority := listCmd.String("priority", "", "Filter by priority: low, medium, high")

	err := listCmd.Parse(args)
	if err != nil {
		fmt.Println("Error parsing flags:", err)
		return
	}

	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}



	var filtered []models.Task
	for _, task := range tasks {
		if *today && !sameDay(task.DueDate, time.Now()) {
			continue
		}
		if *completed && !task.Completed {
			continue
		}
		if *priority != "" && !strings.EqualFold(task.Priority, *priority) {
			continue
		}
		filtered = append(filtered, task)
	}

	if len(filtered) == 0 {
		fmt.Println("No tasks found with the specified filters.")
		return
	}

	for _, task := range filtered {
		status := "❌"
		if task.Completed {
			status = "✅"
		}
		fmt.Printf("[%s] ID: %d | %s (Due: %s | Priority: %s)\n",
			status, task.ID, task.Title, task.DueDate.Format("2006-01-02"), task.Priority)
	}
}





func sameDay(t1, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}