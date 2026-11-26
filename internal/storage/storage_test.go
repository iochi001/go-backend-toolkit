package storage

import (
	"testing"

	"go-backend-toolkit/internal/notes"
)

func TestSaveAndLoadNotes(t *testing.T) {
    // Create temporary storage location.
    filename := t.TempDir() + "/notes.json"

    // Create a note.
    originalNotes := map[int]notes.Note{
        1: {
            ID:      1,
            Title:   "Learning Go",
            Content: "Testing file storage",
        },
    }

    // Save notes.
    err := SaveNotes(filename, originalNotes)
    if err != nil {
        t.Fatalf("failed to save notes: %v", err)
    }

    // Load notes.
    loadedNotes, err := LoadNotes(filename)
    if err != nil {
        t.Fatalf("failed to load notes: %v", err)
    }

    // Verify the loaded note.
    loadedNote, exists := loadedNotes[1]
    if !exists {
        t.Fatal("expected note with ID 1")
    }

    if loadedNote.Title != "Learning Go" {
        t.Errorf(
            "expected title %q, got %q",
            "Learning Go",
            loadedNote.Title,
        )
    }
}
