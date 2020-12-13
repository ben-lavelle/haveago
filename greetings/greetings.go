package greetings

import (
	"errors"
	"fmt"
)

// Pants
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
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
