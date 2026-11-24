package main

import "fmt"

type Note struct {
    ID      int
    Title   string
    Content string
}

func CreateNote(notes map[int]Note, note Note) {
    notes[note.ID] = note
}

func ListNotes(notes map[int]Note) {
    for _, note := range notes {
        fmt.Println("ID:", note.ID)
        fmt.Println("Title:", note.Title)
        fmt.Println("Content:", note.Content)
        fmt.Println("----------------")
    }
}

func FindNote(notes map[int]Note, id int) {
    note, exists := notes[id]

    if !exists {
        fmt.Println("Note not found")
        return
    }

    fmt.Println("Found:", note.Title)
}

func UpdateNote(notes map[int]Note, id int, newTitle string) {
    note, exists := notes[id]

    if !exists {
        fmt.Println("Note not found")
        return
    }

    note.Title = newTitle
    notes[id] = note
}

func DeleteNote(notes map[int]Note, id int) {
    if _, exists := notes[id]; !exists {
        fmt.Println("Note not found")
        return
    }

    delete(notes, id)
    fmt.Println("Note deleted")
}

func main() {
    notes := make(map[int]Note)

    CreateNote(notes, Note{
        ID:      1,
        Title:   "Learning Go",
        Content: "Understanding Go fundamentals",
    })

    CreateNote(notes, Note{
        ID:      2,
        Title:   "Learning SQL",
        Content: "Understanding databases",
    })

    fmt.Println("ALL NOTES")
    ListNotes(notes)

    fmt.Println("\nFIND NOTE")
    FindNote(notes, 1)

    fmt.Println("\nUPDATE NOTE")
    UpdateNote(notes, 1, "Learning Go Backend")

    fmt.Println("\nDELETE NOTE")
    DeleteNote(notes, 2)

    fmt.Println("\nFINAL NOTES")
    ListNotes(notes)
}
