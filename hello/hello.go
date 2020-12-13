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

	// Request a greeting
	message, err := greetings.Hello("Benjamin")
	// If an error is returned, print it and quit.
	if err != nil {
		log.Fatal(err)
	}

	// If not, print the message
	fmt.Println(message)
}
