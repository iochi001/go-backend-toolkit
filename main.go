package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"

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

	case "list":
		handleList()

	case "view":
		handleView(os.Args[2:])

	case "update":
		handleUpdate(os.Args[2:])

	case "delete":
		handleDelete(os.Args[2:])

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
	fmt.Println("  list                              List all notes")
	fmt.Println("  view <id>                         View a note by ID")
	fmt.Println("  update <id> <title> <content>    Update a note by ID")
	fmt.Println("  delete <id>                       Delete a note by ID")
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

func handleList() {
	noteStore, err := loadNoteStore()
	if err != nil {
		fmt.Println("Loading error:", err)
		return
	}

	if len(noteStore) == 0 {
		fmt.Println("No notes found.")
		return
	}

	// Extract note IDs into a slice.
	var ids []int

	for id := range noteStore {
		ids = append(ids, id)
	}

	// Sort IDs in ascending order.
	sort.Ints(ids)

	fmt.Println("Your Notes")
	fmt.Println("------------------------------")

	// Display notes using sorted IDs.
	for _, id := range ids {
		note := noteStore[id]

		fmt.Printf("ID:      %d\n", id)
		fmt.Printf("Title:   %s\n", note.Title)
		fmt.Printf("Content: %s\n", note.Content)
		fmt.Println()
	}

	fmt.Println("------------------------------")
	fmt.Printf("Total notes: %d\n", len(noteStore))
}

func handleView(args []string) {
	// Check that the user supplied an ID.
	if len(args) < 1 {
		fmt.Println("Usage: go run . view <id>")
		return
	}

	// Convert the string ID into an integer.
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid note ID. Please provide a number.")
		return
	}

	// Load existing notes.
	noteStore, err := loadNoteStore()
	if err != nil {
		fmt.Println("Loading error:", err)
		return
	}

	// Find the requested note.
	note, err := notes.FindNote(noteStore, id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Display the note.
	fmt.Println("Note Details")
	fmt.Println("------------------------------")
	fmt.Printf("ID:      %d\n", note.ID)
	fmt.Printf("Title:   %s\n", note.Title)
	fmt.Printf("Content: %s\n", note.Content)
	fmt.Println("------------------------------")
}

func handleUpdate(args []string) {
	if len(args) < 3 {
		fmt.Println("Usage: go run . update <id> <title> <content>")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid note ID. Please provide a number.")
		return
	}

	noteStore, err := loadNoteStore()
	if err != nil {
		fmt.Println("Loading error:", err)
		return
	}

	updatedNote := notes.Note{
		Title:   args[1],
		Content: args[2],
	}

	if err := notes.UpdateNote(noteStore, id, updatedNote); err != nil {
		fmt.Println("Update error:", err)
		return
	}

	if err := storage.SaveNotes(notesFile, noteStore); err != nil {
		fmt.Println("Saving error:", err)
		return
	}

	fmt.Println("Note updated successfully. ID:", id)
}

func handleDelete(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: go run . delete <id>")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid note ID. Please provide a number.")
		return
	}

	noteStore, err := loadNoteStore()
	if err != nil {
		fmt.Println("Loading error:", err)
		return
	}

	if err := notes.DeleteNote(noteStore, id); err != nil {
		fmt.Println("Delete error:", err)
		return
	}

	if err := storage.SaveNotes(notesFile, noteStore); err != nil {
		fmt.Println("Saving error:", err)
		return
	}

	fmt.Println("Note deleted successfully. ID:", id)
}
