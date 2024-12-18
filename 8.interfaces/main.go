package main

import "fmt"

// Concepts
// Interfaces

// Animal Interface Type
type Animal interface {
	Says() string
	NumberOfLegs() int
}

// Dog Interface Type
type Dog struct {
	Name  string
	Breed string
}

// Gorilla Interface Type
type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Samson",
		Breed: "German Sheppard",
	}

	// In golang most receivers are pointer types, receivers being pointer types makes things faster
	// and recommended as a golang best practice in documentation. Therefore we changed the receiver to be a pointer to the interface
	// and then Printo using the reference to pointer
	//PrintInfo(dog)
	PrintInfo(&dog)

	gorilla := Gorilla{
		Name:          "Jock",
		Color:         "Grey",
		NumberOfTeeth: 42,
	}

	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}

// Receiver function for Dog Interface Type
func (d *Dog) Says() string {
	return "Woof"
}

// Receiver function for Dog Interface Type
func (d *Dog) NumberOfLegs() int {
	return 4
}

// Receiver function for Gorilla Interface Type
func (d *Gorilla) Says() string {
	return "Ugh"
}

// Receiver function for Gorilla Interface Type
func (d *Gorilla) NumberOfLegs() int {
	return 2
}
