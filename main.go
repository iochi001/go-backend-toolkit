package main

import "fmt"

type Note struct {
    ID      int
    Title   string
    Content string
}

func (n Note) Display() {
    fmt.Println("Note ID:", n.ID)
    fmt.Println("Title:", n.Title)
    fmt.Println("Content:", n.Content)
}

func (n *Note) UpdateTitle(newTitle string) {
    n.Title = newTitle
}

func main() {
    note := Note{
        ID:      1,
        Title:   "My first note",
        Content: "Learning backend development",
    }

    fmt.Println("Before update:")
    note.Display()

    note.UpdateTitle("Learning Go properly")

    fmt.Println("\nAfter update:")
    note.Display()
}
