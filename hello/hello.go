package main

// 'Code executed as an application must go in a main package'

import (
	"fmt"
	"log"

	"bjml.uk/greetings"
)

func main() {
	// Set up a Logger
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names
	names := []string{"Benjamin", "Robert", "Sara"}
	// Request a greeting
	messages, err := greetings.Hellos(names)
	// If an error is returned, print it and quit.
	if err != nil {
		log.Fatal(err)
	}

	// If not, print the message
	fmt.Println(messages)
}
