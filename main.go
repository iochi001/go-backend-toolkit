package main

import (
	"fmt"
	"os"

	"go-backend-toolkit/internal/notes"
	"go-backend-toolkit/internal/storage"
)

func main() {
	filename := "notes.json"

	// Load existing notes.
	noteStore, err := storage.LoadNotes(filename)

	if err != nil {
		if os.IsNotExist(err) {
			noteStore = make(map[int]notes.Note)
		} else {
			fmt.Println("Loading error:", err)
			return
		}
	}

	note := notes.Note{
		ID:      1,
		Title:   "Learning Go",
		Content: "Understanding persistent storage",
	}

	// Create the note.
	err = notes.CreateNote(noteStore, note)

	if err != nil {
		fmt.Println("Create error:", err)
	} else {
		fmt.Println("Note created successfully")
	}

	// Save notes.
	err = storage.SaveNotes(filename, noteStore)

	if err != nil {
		fmt.Println("Saving error:", err)
		return
	}

	fmt.Println("Notes saved successfully")

	// Find the note.
	foundNote, err := notes.FindNote(noteStore, 1)

	if err != nil {
		fmt.Println("Find error:", err)
	} else {
		fmt.Println("Found note:", foundNote.Title)
	}
}
