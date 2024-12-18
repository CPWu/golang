package main

import (
	"log"

	"github.com/cpwu/channels/helpers"
)

// Concepts
// Channels are unique to Golang.
// They are a means of sending information from one part of your program to another part of your program

// Fixed size argument
const numPool = 10

func main() {
	// 1. Passing argument to a function
	PrintText("Hello")

	// 2. Using Pointers
	message := "Hello, World!"
	ptr := &message
	PrintPointerText(ptr)

	// 3. Channels, Sends information from one part of the program to another part of the program
	// INT Channel
	intChannel := make(chan int)
	// Closes the INT channel, defer ensures that `close()` is ran as soon as intChannel is done.
	// Good practice to ensure for future uses cases like connections are closed where handles may be limited.
	defer close(intChannel)

	// Concurrent operation using go routines.
	go CalculateValue((intChannel))

	num := <-intChannel
	log.Println(num)
}

func PrintText(s string) {
	log.Println(s)
}

func PrintPointerText(s *string) {
	log.Println(*s)
}

func CalculateValue(intChannel chan int) {
	// Generate random number from helper function
	randomNumber := helpers.RandomNumber(numPool)

	// Pass generated number into INT channel
	intChannel <- randomNumber
}
