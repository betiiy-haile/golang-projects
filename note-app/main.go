package main

import (
	"fmt"
	"os"

	"example.com/cli-notes/notes"
)

func main() {
	args := os.Args

	if len(args) < 2{
		fmt.Println("Please provide a command: add, list, search or delete")
		return
	}

	command := args[1]
	switch command {
	case "add": 
		if len(args) < 4 {
			fmt.Println("usage add \"Title\" \"Content\"")
			return
		}
		title := args[2]
		content := args[3]
		notes.AddNote(title, content)
		
	case "list":
		notes.ListNotes()

	case "view": 
		if (len(args) < 3) {
			fmt.Println("Usage: cli-notes view <title>")
			return
		}

		title := args[2]
		notes.ViewNote(title)

	case "delete": 
		if len(os.Args) < 3 {
			fmt.Println("Usage: cli-notes delete <title>")
			return
		}
		title := os.Args[2]
		notes.DeleteNote(title)

	default:
		fmt.Println("❓ Unknown command. Available: add, list, view, delete")
	}
	

}