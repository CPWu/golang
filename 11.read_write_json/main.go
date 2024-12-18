package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Concepts
// Read/Write JSON

// Define type Person struct to receive JSON data
type Person struct {
	FirstName string `json:"first_name"` // Specifies to Golang what to use for names
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
	[
		{
			"first_name": "Clark",
			"last_name": "Kent",
			"hair_color": "black",
			"has_dog": true
		},
		{
			"first_name": "Bruce",
			"last_name": "Wayne",
			"hair_color": "black",
			"has_dog": false
		}
	]`

	// Read JSON into Struct
	// Person slice
	var unmarshalled []Person

	// Unmarshal takes two arguments, slice of bytes and an interface, requires in this case converting from string to slice of bytes
	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled)

	// Write JSON from Struct
	var mySlice []Person

	var m1 Person
	m1.FirstName = "Wally"
	m1.LastName = "West"
	m1.HairColor = "Red"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Diana"
	m2.LastName = "Prince"
	m2.HairColor = "Black"
	m2.HasDog = false

	// Slice with two entries
	mySlice = append(mySlice, m2)

	// Converts slice of structs to JSON
	// MarshallIndent makes it more readable otherwise use Marhshall for production
	newJson, err := json.MarshalIndent(mySlice, "", "     ")
	if err != nil {
		log.Println("error marshalling", err)
	}

	// Convert slice of bytes to String for output
	fmt.Println(string(newJson))
}
