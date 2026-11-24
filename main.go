package main

import "fmt"

type Note struct {
    ID      int
    Title   string
    Content string
}

func ListNotes(notes []Note) {
    for _, note := range notes {
        fmt.Println("ID:", note.ID)
        fmt.Println("Title:", note.Title)
        fmt.Println("Content:", note.Content)
        fmt.Println("----------------")
    }
}

func main() {
    notes := []Note{
        {
            ID:      1,
            Title:   "Go fundamentals",
            Content: "Learning structs and methods",
        },
        {
            ID:      2,
            Title:   "Backend development",
            Content: "Learning application architecture",
        },
    }

    notes = append(notes, Note{
        ID:      3,
        Title:   "Database fundamentals",
        Content: "Learning PostgreSQL",
    })

    ListNotes(notes)
}
