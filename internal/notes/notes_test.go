package notes

import "testing"

func TestCreateNote(t *testing.T) {
	noteStore := make(map[int]Note)

	note := Note{
		ID:      1,
		Title:   "Learning Go",
		Content: "Learning automated testing",
	}

	err := CreateNote(noteStore, note)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(noteStore) != 1 {
		t.Fatalf("expected 1 note, got %d", len(noteStore))
	}
}

func TestCreateDuplicateNote(t *testing.T) {
	noteStore := make(map[int]Note)

	note := Note{
		ID:      1,
		Title:   "Learning Go",
		Content: "Testing duplicate IDs",
	}

	// First creation should succeed.
	err := CreateNote(noteStore, note)

	if err != nil {
		t.Fatalf("unexpected error on first creation: %v", err)
	}

	// Second creation should fail.
	err = CreateNote(noteStore, note)

	if err == nil {
		t.Fatal("expected an error for duplicate note ID")
	}
}

func TestUpdateNote(t *testing.T) {
	noteStore := map[int]Note{
		1: {
			ID:      1,
			Title:   "Old title",
			Content: "Old content",
		},
	}

	updatedNote := Note{
		Title:   "New title",
		Content: "New content",
	}

	err := UpdateNote(noteStore, 1, updatedNote)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	note := noteStore[1]

	if note.Title != "New title" {
		t.Errorf("expected title 'New title', got %q", note.Title)
	}

	if note.Content != "New content" {
		t.Errorf("expected content 'New content', got %q", note.Content)
	}

	if note.ID != 1 {
		t.Errorf("expected ID 1, got %d", note.ID)
	}
}

func TestUpdateNonexistentNote(t *testing.T) {
	noteStore := make(map[int]Note)

	updatedNote := Note{
		Title:   "New title",
		Content: "New content",
	}

	err := UpdateNote(noteStore, 99, updatedNote)

	if err == nil {
		t.Fatal("expected an error when updating a nonexistent note")
	}
}

func TestUpdateNoteWithInvalidContent(t *testing.T) {
	noteStore := map[int]Note{
		1: {
			ID:      1,
			Title:   "Original title",
			Content: "Original content",
		},
	}

	updatedNote := Note{
		Title:   "",
		Content: "New content",
	}

	err := UpdateNote(noteStore, 1, updatedNote)

	if err == nil {
		t.Fatal("expected an error for an empty title")
	}

	if noteStore[1].Title != "Original title" {
		t.Error("original note should remain unchanged after failed update")
	}
}
