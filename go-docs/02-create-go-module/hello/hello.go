package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Greet("Gladys")

	fmt.Println(message)
}
