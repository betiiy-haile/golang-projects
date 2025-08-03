package notes

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"example.com/cli-notes/models"
	"example.com/cli-notes/storage"
)


func New(title, content string) (models.Note, error) {
    if title == "" || content == "" {
        return models.Note{}, errors.New("title and content are required")
    }

    return models.Note{
        Title:     title,
        Content:   content,
        CreatedAt: time.Now(),
    }, nil
}


func AddNote(title, content string) {
	notes, err := storage.LoadNotes()
	if err != nil {
		fmt.Println("❌ Failed to load notes:", err)
		return
	}

	for _, note := range notes {
		if strings.EqualFold(note.Title, title) {
			fmt.Println("⚠️ A note with this title already exists.")
			return
		}
	}

	newNote := models.Note{
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}

	notes = append(notes, newNote)

	err = storage.SaveNotes(notes)
	if err != nil {
		fmt.Println("❌ Failed to save note:", err)
		return 
	}

	fmt.Println("✅ Note added successfully.")
}


func ListNotes()  {
	notes, err := storage.LoadNotes()
	if err != nil {
		fmt.Println("❌ Failed to load notes:", err)
		return 
	}

	if len(notes) == 0{
		fmt.Println("📭 No notes found.")
		return
	}

	for i, note := range notes {
		fmt.Printf("%d. %s (created at: %s) \n", i+1, note.Title, note.CreatedAt.Format("2006-01-02 15:04"))
	}
}


func ViewNote(title string) {
	notes, err := storage.LoadNotes()
	if err != nil {
		fmt.Println("❌ Failed to load notes:", err)
		return
	}

	for _, note := range notes {
		if strings.EqualFold(note.Title, title) {
			fmt.Printf("📝 %s\n%s\n", note.Title, note.Content)
			return
		}
	}

	fmt.Println("⚠️ Note not found.")
}

func DeleteNote(title string) {
	notes, err := storage.LoadNotes()
	if err != nil {
		fmt.Println("❌ Failed to load notes:", err)
		return
	}

	index := -1
	for i, note := range notes {
		if strings.EqualFold(note.Title, title) {
			index = i
			break
		}
	}

	notes = append(notes[:index], notes[index+1:]...)

	err = storage.SaveNotes(notes)
	if err != nil {
		fmt.Println("❌ Failed to delete note:", err)
		return
	}

	fmt.Println("🗑️ Note deleted successfully.")
}