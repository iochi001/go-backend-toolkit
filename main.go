package main

import (
	"fmt"

	"go-backend-toolkit/internal/notes"
)

func main() {
	noteStore := make(map[int]notes.Note)

	note := notes.Note{
		ID:      1,
		Title:   "Learning Go",
		Content: "Understanding packages",
	}

	err := notes.CreateNote(noteStore, note)

	if err != nil {
		fmt.Println("Create error:", err)
	} else {
		fmt.Println("Note created successfully")
	}

	foundNote, err := notes.FindNote(noteStore, 1)

	if err != nil {
		fmt.Println("Find error:", err)
	} else {
		fmt.Println("Found note:", foundNote.Title)
	}
}
