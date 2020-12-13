package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Go exports names that start with capitals for use
// outside the package.

// Hello returns a greeting for a named person.
func Hello(name string) (string, error) {
	// If no name is given, return an error
	if name == "" {
		return "", errors.New("empty name")
	}
	// := declares and initialises; could've used:
	// var message string; message = ...
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// If a pre-run init with Seed isn't called, rand
// assumes Seed(1) -- always the same output.
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	// A slice is like a dynamically-sized array
	// Initially omitting size in [] allows this
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v.",
		"Wassup %v.",
	}

	// Select a random message to return
	return formats[rand.Intn(len(formats))]
}
