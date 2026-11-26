package storage

import (
	"encoding/json"
	"os"

	"go-backend-toolkit/internal/notes"
)

// SaveNotes saves notes to a JSON file.
func SaveNotes(filename string, noteStore map[int]notes.Note) error {
	data, err := json.MarshalIndent(noteStore, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// LoadNotes reads notes from a JSON file.
func LoadNotes(filename string) (map[int]notes.Note, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var noteStore map[int]notes.Note

	err = json.Unmarshal(data, &noteStore)
	if err != nil {
		return nil, err
	}

	return noteStore, nil
}
