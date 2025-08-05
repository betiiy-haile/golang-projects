package main

import (
	"fmt"
	"os"
	"strings"

	"example.com/task-manager/commands"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  add --title=\"Title\" [--description=\"Desc\"] [--due=\"YYYY-MM-DD\"] [--priority=low|medium|high]")
		fmt.Println("  list [--completed] [--today] [--priority=high|medium|low]")
		fmt.Println("  complete <task_id>")
		return
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "add":
		commands.AddTask(args)

	case "list":
		commands.ListTasks(args)

	case "complete":
		if len(args) != 1 {
			fmt.Println("Usage: complete <task_id>")
			return
		}
		idStr := strings.TrimPrefix(args[0], "--id=")
		commands.CompleteTask(idStr)
	
	case "delete":
		if len(args) != 1 {
			fmt.Println(("Usage: delete <task_id>"))
			return
		}
		idStr := strings.TrimPrefix(args[0], "--id=")
		commands.DeleteTask(idStr)

	case "stats":
		commands.TaskStats()

	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Available commands: add, list, complete")
	}
}