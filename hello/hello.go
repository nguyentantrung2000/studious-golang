package main

import (
    "fmt"

    "example.com/greetings"
)

func main() {
    // Get a greeting message and print it.s
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}