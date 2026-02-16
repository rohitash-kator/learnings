package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
	// close(doneChan)
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
	doneChan <- true
	close(doneChan)
}

func main() {
	// Channel:
	// dones := make([]chan bool, 4)
	done := make(chan bool)

	// dones[0] = make(chan bool)
	go greet("Nice to meet you!", done)

	// dones[1] = make(chan bool)
	go greet("How are you?", done)

	// dones[2] = make(chan bool)
	go slowGreet("How...are...you?", done)

	// dones[3] = make(chan bool)
	go greet("I hope I'll be advanced in GoLang soon!!", done)

	// for _, done := range dones {
	// 	<-done
	// }

	for range done {
		// fmt.Println(doneChan)
	}
}
