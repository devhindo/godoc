package main

import (
    "fmt"

    "github.com/devhindo/CreateModule/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}