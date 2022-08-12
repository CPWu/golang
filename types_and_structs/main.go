package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName: "Trevor",
		LastName:  "Sawyer",
    PhoneNumber: "613-414-1134",
	}

	log.Println(user.FirstName, user.LastName, "Phone Number:", user.PhoneNumber)
}
