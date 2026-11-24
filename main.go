package main

import "fmt"

func main() {
    appName := "Go Backend Toolkit"
    version := "1.0.0"

    fmt.Println("Application:", appName)
    fmt.Println("Version:", version)

    greet()
}

func greet() {
    fmt.Println("Welcome! Let's build something.")
}
