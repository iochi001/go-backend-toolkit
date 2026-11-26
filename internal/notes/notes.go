
package notes

import (
	"fmt"
	"strings"
)

type Note struct {
	ID      int
	Title   string
	Content string
}

func ValidateNote(note Note) error {
	if strings.TrimSpace(note.Title) == "" {
		return fmt.Errorf("note title cannot be empty")
	}

	if strings.TrimSpace(note.Content) == "" {
		return fmt.Errorf("note content cannot be empty")
	}

	return nil
}

func CreateNote(notes map[int]Note, note Note) error {
	if err := ValidateNote(note); err != nil {
		return err
	}

	if _, exists := notes[note.ID]; exists {
		return fmt.Errorf("note with ID %d already exists", note.ID)
	}

	notes[note.ID] = note
	return nil
}

func FindNote(notes map[int]Note, id int) (Note, error) {
	note, exists := notes[id]

	if !exists {
		return Note{}, fmt.Errorf("note with ID %d not found", id)
	}

	return note, nil
}
