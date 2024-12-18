package main

import "log"

// Concepts
// Loops and Ranges

func main() {
	// For Loop
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}

	// Slice of Strings
	animals := []string{"dog", "fish", "horse", "cow", "cat"}

	for i, animal := range animals {
		log.Println(i, animal)
	}

	// Advanced
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"John", "Smith", "john@smith.com", 30})
	users = append(users, User{"Mary", "Jones", "mary@jones.com", 20})
	users = append(users, User{"Sally", "Brown", "sally@brown.com", 25})
	users = append(users, User{"Alex", "Anderson", "alex@anderson", 35})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}
