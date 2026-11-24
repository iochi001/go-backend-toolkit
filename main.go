package main

import "fmt"

type Note struct {
    ID      int
    Title   string
    Content string
}

func main() {
    note := Note{
        ID:      1,
        Title:   "My first note",
        Content: "I am learning Go backend development",
    }

    fmt.Println("Note ID:", note.ID)
    fmt.Println("Title:", note.Title)
    fmt.Println("Content:", note.Content)
}