package main

// Concepts
// Variable Scope and Variable Shadowing

import (
	"log"
	"time"
)

// Variable type for `s` is inferred and has package scope
var s = "seven"



type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	// Variable type for `s2` is inferred and has function level scope.
	var s2 = "six"

	log.Println("s is", s)
	log.Println("s2 is", s2)

	user := User{
		FirstName: "Trevor",
		LastName:  "Sawyer",
    PhoneNumber: "613-414-1134",
	}

	log.Println(user.FirstName, user.LastName, "Phone Number:", user.PhoneNumber)
}

func saySomething(s string) (string, string) {
	return s, "World"
}