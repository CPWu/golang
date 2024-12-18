package main

import "log"

// Pointers points to a specific location in memory that may hold a value
func main() {
	// Locally scoped variable
	var myString string
	myString = "Green"

	log.Println("myString is set to", myString)
	// Passes a reference to myString to be modified
	changeUsingPointer(&myString)
	log.Println("after function call myString is set to:", myString)
}

func changeUsingPointer(s *string) {
	// S is now set to a memory address
	log.Println("s is set to:", s)
	newValue := "Red"
	*s = newValue
}
