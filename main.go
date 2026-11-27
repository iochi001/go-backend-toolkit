package main

import (
	"fmt"
	"os"

	"go-backend-toolkit/internal/notes"
	"go-backend-toolkit/internal/storage"
)

const notesFile = "notes.json"

func main() {
	if len(os.Args) < 2 {
		showHelp()
		return
	}

	command := os.Args[1]

	switch command {
	case "help":
		showHelp()

	case "create":
		handleCreate(os.Args[2:])

	default:
		fmt.Println("Unknown command:", command)
		showHelp()
	}
}

func showHelp() {
	fmt.Println("Go Backend Toolkit - Notes CLI")
	fmt.Println()
	fmt.Println("Available commands:")
	fmt.Println("  help                              Show available commands")
	fmt.Println("  create <title> <content>         Create a new note")
}

func handleCreate(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: go run . create <title> <content>")
		return
	}

	noteStore, err := loadNoteStore()
	if err != nil {
		fmt.Println("Loading error:", err)
		return
	}

	newID := nextNoteID(noteStore)

	note := notes.Note{
		ID:      newID,
		Title:   args[0],
		Content: args[1],
	}

	if err := notes.CreateNote(noteStore, note); err != nil {
		fmt.Println("Create error:", err)
		return
	}

	if err := storage.SaveNotes(notesFile, noteStore); err != nil {
		fmt.Println("Saving error:", err)
		return
	}

	fmt.Println("Note created successfully. ID:", newID)
}

func loadNoteStore() (map[int]notes.Note, error) {
	noteStore, err := storage.LoadNotes(notesFile)

	if err != nil {
		if os.IsNotExist(err) {
			return make(map[int]notes.Note), nil
		}
		return nil, err
	}

	if noteStore == nil {
		noteStore = make(map[int]notes.Note)
	}

	return noteStore, nil
}

func nextNoteID(noteStore map[int]notes.Note) int {
	newID := 1

	for id := range noteStore {
		if id >= newID {
			newID = id + 1
		}
	}

	return newID
}
