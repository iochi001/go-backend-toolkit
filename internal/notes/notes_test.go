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
