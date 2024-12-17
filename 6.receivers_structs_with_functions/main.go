package main

import (
	"log"
)

type myStruct struct {
	FirstName string
}

// Receiver, ties this function to the type of myStruct
func (m *myStruct) printFirstName() string {
	return m.FirstName

}
func main() {
	var myVar myStruct
	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "Mary",
	}

	log.Println("myVar is set to", myVar.FirstName)
	log.Println("myVar2 is set to", myVar2.FirstName)

	// Using receivers
	log.Println("Using receivers myVar is set to", myVar.printFirstName())
	log.Println("Using receivers myVar2 is set to", myVar2.printFirstName())
}
