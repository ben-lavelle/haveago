package main

// 'Code executed as an application must go in a main package'

import (
	"fmt"

	"bjml.uk/greetings"
)

func main() {
	message := greetings.Hello("Benjamin")
	fmt.Println(message)
}
