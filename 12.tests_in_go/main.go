package main

import (
	"errors"
	"log"
)

// Concepts
// Testing with Golang
// In Golang, tests live right besides the code that is being tested with the naming convention *_test.go

func main() {
	result, err := divide(100.0, 10.0)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("The result of division is:", result)
}

func divide(x, y float32) (float32, error) {
	var result float32

	// Division by zero is no good, so we test for this
	if y == 0 {
		return result, errors.New("cannot divide by zero")
	}

	result = x / y
	return result, nil
}
